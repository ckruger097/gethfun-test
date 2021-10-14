package transactions

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

func SendRawTx(client *ethclient.Client) (common.Hash, error){
	rawTxBytes, err := CreateRawTx(client)
	if err != nil {
		log.Fatal(err)
	}
	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), tx)

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return tx.Hash(), nil

}
