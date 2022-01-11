package smart_contracts

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ReadBytecode(client *ethclient.Client, contract string) (string, error) {
	address := common.HexToAddress(contract)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytecode), nil
}
