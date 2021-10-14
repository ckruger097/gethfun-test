package accounts

import (
	"context"
	"fmt"
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
	//fmt.Println("is 0x323b5d4c32345ced77393b3530b1eed0f346429d a valid address?", checkIfValid("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // true
	//fmt.Println("is 0xZYXb5d4c32345ced77393b3530b1eed0f346429d a valid address?", checkIfValid("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // false
}

func ValidSmartContract(client *ethclient.Client, addressString string) {
	address := common.HexToAddress(addressString)
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: ??
}


