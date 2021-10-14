package transactions

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func SubscribeBlocks(client *ethclient.Client) {
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("block hash:", block.Hash().Hex())
			fmt.Println("block number:", block.Number().Uint64())
			fmt.Println("block time:", block.Time())
			fmt.Println("block nonce:", block.Nonce())
			fmt.Println("block num of tx", len(block.Transactions()))
		}
	}
}
