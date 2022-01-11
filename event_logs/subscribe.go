package event_logs

import (
	"context"
	"fmt"
	store "gethfun/build"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func SubscribeToEvent(client ethclient.Client, contractAddress string) {
	addr := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}
	// make a channel to subscribe to the log
	logs := make(chan types.Log)
	// subscribe
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

func ReadStoreLog(client *ethclient.Client, contractAddress string, blockBegin *big.Int, blockEnd *big.Int) {
	addr := common.HexToAddress(contractAddress)
	query := ethereum.FilterQuery{
		FromBlock: blockBegin,
		ToBlock:   blockEnd,
		Addresses: []common.Address{
			addr,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(event.Key[:]))
		fmt.Println(string(event.Value[:]))

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0])
	}

	eventSig := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSig)
	fmt.Println(hash.Hex())
}
