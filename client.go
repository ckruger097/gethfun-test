package main

import (
	"fmt"
	"gethfun/accessories"
	"gethfun/accounts"
	"gethfun/smart_contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	rinkebyurl := accessories.GoDotEnvVariable("rinkeby")
	client, err := ethclient.Dial(rinkebyurl) // rinkeby
	if err != nil {
		log.Fatal(err)
	}

	smart_contracts.DeployEthCheck(client, "bababooie", big.NewFloat(1.0))

	//_, err = transactions.SendRawTx(client)
	//smart_contracts.DeploySmartContract(client)
	//instance, err := smart_contracts.LoadSmartContract(client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//smart_contracts.QuerySmartContract(instance)
	//smart_contracts.WriteToSmartContract(client, instance, "bobo", "bear")
	//smart_contracts.ReadSmartContract(instance, "bobo")

	//txhash, err := transactions.TestToken(client)
	//fmt.Println("Yay our Tx hash is:", txhash)

	//myWallet := loadCreateWallet(client)

	//success, err := transactions.SendTokenFlow(client)
	//if success {
	//	fmt.Println("\nNice!")
	//}else{
	//	log.Fatal(err)
	//}

}

func loadCreateWallet(client *ethclient.Client) accounts.EthWallet {
	fmt.Println("Create a new wallet or use a pre-existing one? --- [1] [2] ---")
	var newWallet accounts.EthWallet
	response := accessories.UserInputLine()
	if response == "1" {
		fmt.Println("What will you name your wallet?")
		wn := accessories.UserInputLine()
		newWallet = accounts.GenerateWallet(wn)
	} else if response == "2" {
		fmt.Println("Enter your private key")
		pk := accessories.UserInputLine()
		newWallet = accounts.LoadWallet(pk)
	} else {
		fmt.Println("Invalid! Aborting program (are you happy?)")
		os.Exit(1)
	}
	address := accounts.GetCommonAddress(newWallet.WalletAddress)
	accounts.PrintBalanceInfo(client, address)
	return newWallet
}

func clientKeyStore() {
	accy, err := accounts.CreateKs("okletsgo")
	fmt.Println(accy)
	time.Sleep(10000)
	accy, err = accounts.ImportKs()
	if err != nil {
		log.Fatal(err)
	}
}
