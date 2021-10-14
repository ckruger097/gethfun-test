package smart_contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gethfun/accessories"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// ethereum-checks:
// make a smart contract that you can send to someone with a passphrase that unlocks pre-determined funds to be deposited to their address
// smart contract:
// needs to be able to 1. be deployed with an already hashed key, and the amount of funds to deposit.
// 2. have the ability to take in a passphrase through msg.sender
// 3. be able to use msg.sender to decrypt hash.
// 4. if hash == hash, deposit allotted amount into recipient address

func DeployEthCheck(client *ethclient.Client, passphrase string, value *big.Float) {
	goenvprivkey := accessories.GoDotEnvVariable("privkey")
	privateKey, err := crypto.HexToECDSA(goenvprivkey)
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("couldnt cast to ecdsa pub")
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
	auth.GasLimit = uint64(900000) // gas price increased, idk what limit is now but gas used was >300k see https://rinkeby.etherscan.io/tx/0xa26c63f190ccbf7e1a59efa85ba38554056f5760dfe42bef1229fbff00d79718
	auth.GasPrice = gasPrice
	keyBytes := []byte(passphrase)
	encryptedKey := crypto.Keccak256Hash(keyBytes)
	fmt.Println(encryptedKey)

}
