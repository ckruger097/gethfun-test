package main

import (
	"context"
	"fmt"
	"gethfun/accessories"
	"gethfun/accounts"
	"gethfun/smart_contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

func main() {
	rinkebyurl := accessories.GoDotEnvVariable("rinkeby")
	client, err := ethclient.Dial(rinkebyurl) // rinkeby
	if err != nil {
		log.Fatal(err)
	}

	//smart_contracts.DeploySmartContract(client)
	contract, err := smart_contracts.LoadSmartContract(client, "0xC25249Bb74dd83a0470A0df7e5C93BDaFFb300e9")
	if err != nil {
		log.Fatal(err)
	}
	ver, err := contract.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ver)
	tx, err := smart_contracts.WriteToSmartContract(client, contract, "harry", "33")
	if err != nil {
		log.Fatal(err)
	}
	//time.Sleep(30)
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	res, err := smart_contracts.ReadSmartContract(contract, "harry")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
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
		pk = fmt.Sprintf("0x%s", pk)
		newWallet = accounts.LoadWallet(pk)
	} else {
		fmt.Println("Invalid! Aborting program (are you happy?)")
		os.Exit(1)
	}
	address := accounts.GetCommonAddress(newWallet.WalletAddress)
	accounts.PrintBalanceInfo(client, address)
	return newWallet
}
