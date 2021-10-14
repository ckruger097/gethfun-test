package transactions

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"gethfun/accessories"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func CreateRawTx(client *ethclient.Client) ([]byte, error){
	privateKey, err := crypto.HexToECDSA("5cdd771d2e0af548c308d4740d12cfb523a51be2f34ef839e7e3a960297df5fb")
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privateKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("couldnt cast to ecdsa pub")
	}
	fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := accessories.EtherToWei(big.NewFloat(0.69))
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x546E6aD2e7AA8C047AE5D022D183c5918D7b158B")
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	})
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	var buff bytes.Buffer
	err = signedTx.EncodeRLP(&buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("your raw tx: 0x%x\n", buff.Bytes())
	return buff.Bytes(), nil
}
