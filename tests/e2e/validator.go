package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	sdkmath "cosmossdk.io/math"
	cometcfg "github.com/cometbft/cometbft/config"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/privval"
	sdkcrypto "github.com/cosmos/cosmos-sdk/crypto"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilstypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/skip-mev/slinky/tests/simapp"
)

type validator struct {
	chain        *chain
	index        int
	moniker      string
	mnemonic     string
	keyInfo      keyring.Record
	privateKey   cryptotypes.PrivKey
	consensusKey privval.FilePVKey
	nodeKey      p2p.NodeKey
}

func (v *validator) instanceName() string {
	return fmt.Sprintf("%s%d", v.moniker, v.index)
}

func (v *validator) configDir() string {
	return fmt.Sprintf("%s/%s", v.chain.configDir(), v.instanceName())
}

func (v *validator) createConfig() error {
	p := path.Join(v.configDir(), "config")
	return os.MkdirAll(p, 0o755)
}

func (v *validator) init() error {
	if err := v.createConfig(); err != nil {
		return err
	}

	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	genDoc, err := getGenDoc(v.configDir())
	if err != nil {
		return err
	}

	appState, err := json.MarshalIndent(simapp.ModuleBasics.DefaultGenesis(cdc), "", " ")
	if err != nil {
		return fmt.Errorf("failed to JSON encode app genesis state: %w", err)
	}

	genDoc.ChainID = v.chain.id
	genDoc.Validators = nil
	genDoc.AppState = appState

	if err := genDoc.SaveAs(config.GenesisFile()); err != nil {
		return err
	}

	genAppState, err := genutilstypes.AppGenesisFromFile(config.GenesisFile())
	if err != nil {
		return fmt.Errorf("failed to unmarshal genesis state: %w", err)
	}

	if err = genutil.ExportGenesisFile(genAppState, config.GenesisFile()); err != nil {
		return fmt.Errorf("failed to export app genesis state: %w", err)
	}

	cometcfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)
	return nil
}

func (v *validator) createNodeKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return err
	}

	v.nodeKey = *nodeKey
	return nil
}

func (v *validator) createConsensusKey() error {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(v.configDir())
	config.Moniker = v.moniker

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := os.MkdirAll(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return fmt.Errorf("could not create directory %q: %w", filepath.Dir(pvKeyFile), err)
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := os.MkdirAll(filepath.Dir(pvStateFile), 0o777); err != nil {
		return fmt.Errorf("could not create directory %q: %w", filepath.Dir(pvStateFile), err)
	}

	filePV := privval.LoadOrGenFilePV(pvKeyFile, pvStateFile)
	v.consensusKey = filePV.Key

	return nil
}

func (v *validator) createKeyFromMnemonic(name, mnemonic string) error {
	kb, err := keyring.New(keyringAppName, keyring.BackendTest, v.configDir(), nil, cdc)
	if err != nil {
		return err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), keyringAlgos)
	if err != nil {
		return err
	}

	info, err := kb.NewAccount(name, mnemonic, "", sdk.FullFundraiserPath, algo)
	if err != nil {
		return err
	}

	privKeyArmor, err := kb.ExportPrivKeyArmor(name, keyringPassphrase)
	if err != nil {
		return err
	}

	privKey, _, err := sdkcrypto.UnarmorDecryptPrivKey(privKeyArmor, keyringPassphrase)
	if err != nil {
		return err
	}

	v.keyInfo = *info
	v.mnemonic = mnemonic
	v.privateKey = privKey

	return nil
}

func (v *validator) createKey(name string) error {
	mnemonic, err := createMnemonic()
	if err != nil {
		return err
	}

	return v.createKeyFromMnemonic(name, mnemonic)
}

func (v *validator) buildCreateValidatorMsg(amount sdk.Coin) (sdk.Msg, error) {
	description := stakingtypes.NewDescription(v.moniker, "", "", "", "")
	commissionRates := stakingtypes.CommissionRates{
		Rate:          sdkmath.LegacyMustNewDecFromStr("0.1"),
		MaxRate:       sdkmath.LegacyMustNewDecFromStr("0.2"),
		MaxChangeRate: sdkmath.LegacyMustNewDecFromStr("0.01"),
	}

	// get the initial validator min self delegation
	minSelfDelegation := sdkmath.NewInt(1)

	valPubKey, err := cryptocodec.FromCmtPubKeyInterface(v.consensusKey.PubKey)
	if err != nil {
		return nil, err
	}
	valAddr, err := v.keyInfo.GetAddress()
	if err != nil {
		return nil, err
	}

	return stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(valAddr).String(), // todo: switch to address codec.
		valPubKey,
		amount,
		description,
		commissionRates,
		minSelfDelegation,
	)
}

func (v *validator) signMsg(msgs ...sdk.Msg) (*sdktx.Tx, error) {
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	if err := txBuilder.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	txBuilder.SetMemo(fmt.Sprintf("%s@%s:26656", v.nodeKey.ID(), v.instanceName()))
	txBuilder.SetFeeAmount(sdk.NewCoins())
	txBuilder.SetGasLimit(200_000)

	pubKey, err := v.keyInfo.GetPubKey()
	if err != nil {
		return nil, err
	}

	signerData := authsigning.SignerData{
		ChainID:       v.chain.id,
		AccountNumber: 0,
		Sequence:      0,
		PubKey:        pubKey,
	}

	// For SIGN_MODE_DIRECT, calling SetSignatures calls setSignerInfos on
	// TxBuilder under the hood, and SignerInfos is needed to generate the sign
	// bytes. This is the reason for setting SetSignatures here, with a nil
	// signature.
	//
	// Note: This line is not needed for SIGN_MODE_LEGACY_AMINO, but putting it
	// also doesn't affect its generated sign bytes, so for code's simplicity
	// sake, we put it here.
	if err != nil {
		return nil, err
	}

	sig := signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: nil,
		},
		Sequence: 0,
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	bytesToSign, err := authsigning.GetSignBytesAdapter(
		context.Background(),
		encodingConfig.TxConfig.SignModeHandler(),
		signing.SignMode_SIGN_MODE_DIRECT,
		signerData,
		txBuilder.GetTx(),
	)
	if err != nil {
		return nil, err
	}

	sigBytes, err := v.privateKey.Sign(bytesToSign)
	if err != nil {
		return nil, err
	}

	sig = signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
			Signature: sigBytes,
		},
		Sequence: 0,
	}
	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	signedTx := txBuilder.GetTx()
	bz, err := encodingConfig.TxConfig.TxEncoder()(signedTx)
	if err != nil {
		return nil, err
	}

	return decodeTx(bz)
}
