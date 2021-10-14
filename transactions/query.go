package transactions

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func QueryBlocks(client *ethclient.Client) *types.Block {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String())
	blockNumber := big.NewInt(13040280)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(block.Number().Uint64())
	//fmt.Println(block.Time())
	//fmt.Println(block.Difficulty().Uint64())
	//fmt.Println(block.Hash().Hex())
	//fmt.Println(len(block.Transactions()))
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("this block has", count,"transactions")
	return block
}

func QueryTx(client *ethclient.Client, block *types.Block) {
	//for _, tx := range block.Transactions() {
	//
	//	chainId, err := client.NetworkID(context.Background())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if msg, err := tx.AsMessage(types.NewLondonSigner(chainId), tx.GasPrice()); err != nil {
	//	fmt.Println(msg.From().Hex())}
	//	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(receipt.Logs)
	//}
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)       // false
}