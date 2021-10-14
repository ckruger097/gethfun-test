package transactions

import (
	"context"
	"fmt"
	"gethfun/accessories"
	"gethfun/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metachris/eth-go-bindings/erc20"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// TestToken "0x014dF965e86d241b7CB1303C235CBDB197cf3e2e" banan
func TestToken(client *ethclient.Client) (string, error) {
	// have user enter smart contract address
	fmt.Println("Enter the token contract address:")
	userInputAddress := accessories.UserInputLine()
	address := common.HexToAddress(userInputAddress)
	// generate token bindings
	token, err := erc20.NewErc20(address, client)
	if err != nil {
		log.Fatal(err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatal(err)
	}
	// confirm token name
	fmt.Println("Your token's name is:", name)
	fmt.Println("Is this correct? [Y/n]")
	userResponse := accessories.UserInputLine()
	if strings.ToLower(userResponse) == "y" || userResponse == " " {
		fmt.Println("Continuing transaction flow...")
	} else {
		return "", fmt.Errorf("user responded incorrect token")
	}
	// choose keystore file
	fmt.Println("Files currently in keystore:")
	var files []string
	root := "./keystores"
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("Which keystore do you want to unlock?")
	fileName := accessories.UserInputLine()
	file := fmt.Sprintf("./keystores/%s", fileName)
	key, err := ioutil.ReadFile(file)
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// unlock keystore
	fmt.Println("What's the password?")
	userPass := accessories.UserInputLine()
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(key)), userPass, chainId) 
	if err != nil {
		log.Fatal(err)
	}
	// get address of receiver
	fmt.Println("Enter the address of the token receiver:")
	userReceiverAddress := accessories.UserInputLine()
	if !accounts.ValidEthAddress(userReceiverAddress){
		log.Fatal("invalid eth address")
	}
	// enter value
	fmt.Printf("Please enter value (in %s) you would like to send: ", name)
	valueUserInput, err := accessories.UserInputFloat()
	if err != nil {
		log.Fatal(err)
	}
	valueFloat64 := valueUserInput.String()
	float, err := strconv.ParseFloat(valueFloat64, 64)
	if err != nil {
		return "", err
	}
	value := accessories.EtherToWei(big.NewFloat(float))
	// make tx
	tx, err := token.Transfer(auth, common.HexToAddress(userReceiverAddress), value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("transfer pending: 0x%x\n", tx.Hash())
	return tx.Hash().Hex(), nil
}

