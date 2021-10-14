package smart_contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	store "gethfun/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func WriteToSmartContract(client *ethclient.Client, instance *store.Store, key string, value string) {
	privateKey, err := crypto.HexToECDSA("5cdd771d2e0af548c308d4740d12cfb523a51be2f34ef839e7e3a960297df5fb")
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
	auth.GasLimit = uint64(1000000) // gas price increased, idk what limit is now but gas used was >300k see https://rinkeby.etherscan.io/tx/0xa26c63f190ccbf7e1a59efa85ba38554056f5760dfe42bef1229fbff00d79718
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
	fmt.Printf("tx sent: %x", tx.Hash())

}

func ReadSmartContract(instance *store.Store, key string) {
	keyBytes := [32]byte{}
	copy(keyBytes[:], key)
	result, err := instance.Items(nil, keyBytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result[:]))
}
