package contracts

import (
	"context"
	"crypto/ecdsa"
	cfg "file-registory/config"
	freg "file-registory/contracts/file_registry"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	Instance    *freg.FileRegistry
	Clinet      *ethclient.Client
	PrivateKey  *ecdsa.PrivateKey
	FromAddress common.Address
}

func NewContract(config *cfg.Config) (*Contract, error) {
	client, err := ethclient.Dial(config.SepoliaUrl)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.ContractAddress)
	instance, err := freg.NewFileRegistry(address, client)
	if err != nil {
		return nil, err
	}

	return &Contract{
		Instance:    instance,
		Clinet:      client,
		PrivateKey:  privateKey,
		FromAddress: fromAddress,
	}, nil
}

func (contract *Contract) Save(filePath string, cid string) error {
	nonce, err := contract.Clinet.PendingNonceAt(context.Background(), contract.FromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := contract.Clinet.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(contract.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	_, err = contract.Instance.Save(auth, filePath, cid)
	return err
}

func (contract *Contract) Get(filePath string) (string, error) {
	return contract.Instance.Get(&bind.CallOpts{}, filePath)
}
