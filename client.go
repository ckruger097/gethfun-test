package main

import (
	"fmt"
	"gethfun/accessories"
	"gethfun/accounts"
	"gethfun/event_logs"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	rinkebyurl := accessories.GoDotEnvVariable("rinkeby")
	client, err := ethclient.Dial(rinkebyurl) // rinkeby
	if err != nil {
		log.Fatal(err)
	}
	rinkebywss := accessories.GoDotEnvVariable("rinkebywss")
	clientwss, err := ethclient.Dial(rinkebywss)
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	//contract, err := smart_contracts.LoadStore(client, accessories.GoDotEnvVariable("storeContract"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	event_logs.ReadStoreLog(clientwss, accessories.GoDotEnvVariable("storeContract"), big.NewInt(9964990),
		big.NewInt(9974533))
}

func loadCreateWallet(client *ethclient.Client) accounts.EthWallet {
	fmt.Println("Create a new wallet, use a pre-existing one, or use env? --- [1] [2] [3] ---")
	var newWallet accounts.EthWallet
	response := accessories.UserInputLine()
	if response == "1" {
		fmt.Println("What will you name your wallet?")
		wn := accessories.UserInputLine()
		newWallet = accounts.GenerateWallet(wn)
	} else if response == "2" {
		fmt.Println("Enter your private key")
		pk := accessories.UserInputLine()
		pk = fmt.Sprintf("0x%s", pk)
		newWallet = accounts.LoadWallet(pk)
	} else {
		pk := accessories.GoDotEnvVariable("privkey")
		pk = fmt.Sprintf("0x%s", pk)
		newWallet = accounts.LoadWallet(pk)
	}
	//address := accounts.GetCommonAddress(newWallet.WalletAddress)
	//accounts.PrintBalanceInfo(client, address)
	return newWallet
}
