package smart_contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gethfun/accessories"
	store "gethfun/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func QuerySmartContract(instance *store.Store) {
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your version of Store is:", version)
}

func WriteToSmartContract(client *ethclient.Client, instance *store.Store, key string, value string) (*types.Transaction, error) {
	goenvprivkey := accessories.GoDotEnvVariable("privkey")
	privateKey, err := crypto.HexToECDSA(goenvprivkey)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("couldn't cast to ecdsa pub")
	}
	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(1000000) // gas price increased
	auth.GasPrice = gasPrice
	// store has a keystore method, but it's in byte32. we create our key and value here.
	keyBytes := [32]byte{}
	valueBytes := [32]byte{}
	copy(keyBytes[:], key)
	copy(valueBytes[:], value)
	// we create our tx in which we call the method w/ its necessary parameters here
	tx, err := instance.SetItem(auth, keyBytes, valueBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: https://rinkeby.etherscan.io/tx/%s\n", tx.Hash().Hex())
	return tx, nil

}

func ReadSmartContract(instance *store.Store, key string) (string, error) {
	keyBytes := [32]byte{}
	copy(keyBytes[:], key)
	result, err := instance.Items(nil, keyBytes)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(result[:]), nil
}
