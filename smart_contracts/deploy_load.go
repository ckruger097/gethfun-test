package smart_contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gethfun/accessories"
	store "gethfun/build"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func DeploySmartContract(version string, client *ethclient.Client) {
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
	auth.GasLimit = uint64(900000) // gas price increased,
	auth.GasPrice = gasPrice

	address, tx, instance, err := store.DeployStore(auth, client, version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("address is:", address)
	fmt.Println("tx hash is:", tx.Hash().Hex())
	_ = instance
}

func LoadSmartContract(client *ethclient.Client, contractAddress string) (*store.Store, error) {
	address := common.HexToAddress(contractAddress) // the addr you get from deploy
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract loaded")
	return instance, nil
}
