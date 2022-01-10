package accounts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func GetCommonAddress(givenAddr string) common.Address {
	address := common.HexToAddress(givenAddr)
	//fmt.Println(address.Hex())
	//fmt.Println(address.Hash().Hex())
	//fmt.Println(address.Bytes())
	return address
}

func ValidEthAddress(givenAddr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(givenAddr)

}

func ValidSmartContract(client *ethclient.Client, addressString string) bool {
	address := common.HexToAddress(addressString)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0
	//fmt.Printf("is contract: %v\n", isContract) // is contract: ??
	return isContract
}
