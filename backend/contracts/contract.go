package contracts

import (
	"context"
	"crypto/ecdsa"
	cfg "file-registory/config"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	Instance *Storage
	Auth     *bind.TransactOpts
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.ContractAddr)
	instance, err := NewStorage(address, client)
	if err != nil {
		return nil, err
	}

	return &Contract{
		Instance: instance,
		Auth:     auth,
	}, nil
}

func (contract *Contract) Save(filePath string, cid string) error {
	_, err := contract.Instance.Save(contract.Auth, filePath, cid)
	return err
}

func (contract *Contract) Get(filePath string) (string, error) {
	return contract.Instance.Get(&bind.CallOpts{}, filePath)
}
