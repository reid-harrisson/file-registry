package contracts

import (
	"context"
	"crypto/ecdsa"
	cfg "file-registory/config"
	"file-registory/contracts/store"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	Instance    *store.Store
	Client      *ethclient.Client
	PrivateKey  *ecdsa.PrivateKey
	FromAddress common.Address
}

func NewContract(config *cfg.Config) (*Contract, error) {
	client, err := ethclient.Dial(config.RpcSrvUrl)
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

	return &Contract{
		Instance:    nil,
		Client:      client,
		PrivateKey:  privateKey,
		FromAddress: fromAddress,
	}, nil
}

func (contract *Contract) Deploy() error {
	nonce, err := contract.Client.PendingNonceAt(context.Background(), contract.FromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := contract.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(contract.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := store.DeployStore(auth, contract.Client)
	if err != nil {
		return err
	}
	fmt.Printf("Contract deployed at address: %s\n", address.Hex())
	fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
	contract.Instance = instance
	return nil
}

func (contract *Contract) Save(filePath string, cid string) error {
	// Check if the contract instance is properly initialized
	if contract.Instance == nil {
		return fmt.Errorf("contract instance is not initialized")
	}

	nonce, err := contract.Client.PendingNonceAt(context.Background(), contract.FromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := contract.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(contract.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = contract.Instance.Save(auth, filePath, cid)
	return err
}

func (contract *Contract) Get(filePath string) (string, error) {
	// Check if the contract instance is properly initialized
	if contract.Instance == nil {
		return "", fmt.Errorf("contract instance is not initialized")
	}

	return contract.Instance.Get(&bind.CallOpts{}, filePath)
}
