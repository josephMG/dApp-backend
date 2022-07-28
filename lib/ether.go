package lib

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"hardhat-backend/config"
	"hardhat-backend/lib/loggers"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	greeter "hardhat-backend/contracts"
)

// EtherClient modal
type EtherClient struct {
	*ethclient.Client
	loggers.Logger
	Env *config.Env
}

// NewEther creates a new ehter client instance
func NewEther(env *config.Env, logger loggers.Logger) EtherClient {
	client, err := ethclient.Dial(env.EthereumURL)
	if err != nil {
		logger.Error(err.Error())
	}
	defer client.Close()

	logger.Info("we have a connection")
	_ = client // we'll use this in the upcoming sections
	return EtherClient{
		Client: client,
		Logger: logger,
		Env:    env,
	}
}

// GetBalance get balance from address
func (c EtherClient) GetBalance(address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := c.Client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		c.Logger.Error(err.Error())
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}

// GetGreeting get greeting string from contract function address
func (c EtherClient) GetGreeting() string {
	blockHeader, _ := c.Client.HeaderByNumber(context.Background(), nil)

	contractAddr := common.HexToAddress(c.Env.ContractAddr)
	data, _ := hexutil.Decode("0xcfae3217")
	callMsg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}
	res, err := c.Client.CallContract(context.Background(), callMsg, blockHeader.Number)
	if err != nil {
		c.Logger.Fatalf("Error calling contract: %v", err)
	}

	s, _ := hex.DecodeString(common.Bytes2Hex(res[:]))
	c.Logger.Info(res[:])
	c.Logger.Info(string(s))
	res = bytes.Trim(res[:], "\x00 \x0c")
	return strings.TrimSpace(string(res[:]))
}

// GetGreeting get greeting string from contract function address
func (c EtherClient) GetGreetingFromContract() string {
	contractAddr := common.HexToAddress(c.Env.ContractAddr)
	instance, err := greeter.NewGreeter(contractAddr, c.Client)

	tx, err := instance.Greet(new(bind.CallOpts))
	if err != nil {
		c.Logger.Fatal(err)
	}
	return tx
}

// GetGreeting get greeting string from contract function address
func (c EtherClient) PostGreeting(greeting string) string {
	auth, err := GetAuth(c, c.Env.AccountPrivateKey)

	contractAddr := common.HexToAddress(c.Env.ContractAddr)
	instance, err := greeter.NewGreeter(contractAddr, c.Client)

	tx, err := instance.SetGreeting(auth, greeting)
	if err != nil {
		c.Logger.Fatal(err)
	}

	c.Logger.Infof("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	return "ok"
}

func GetAuth(c EtherClient, accountAddress string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		c.Logger.Fatal(err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		c.Logger.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.Logger.Fatal(err)
		return nil, err
	}

	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		c.Logger.Fatal(err)
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}
