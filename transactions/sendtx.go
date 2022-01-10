package transactions

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"gethfun/accessories"
	"gethfun/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func SendTxFlowPrivKey(client *ethclient.Client) (bool, error) {
	// enter our privkey convert to ecdsa
	fmt.Println("Would you like to send a transaction? [Y/N]")
	userInput := accessories.UserInputLine()
	if strings.ToLower(userInput) == "y" {
		fmt.Println("Please enter the address to send Ethereum to:")
		// define recipient address
		toAddress := accessories.UserInputLine()
		if !accounts.ValidEthAddress(toAddress) {
			return false, fmt.Errorf("invalid recepient")
		}
		recpient := common.HexToAddress(toAddress)
		fmt.Println("Enter value of transaction being sent in ETH:")
		userInputValue, err := accessories.UserInputFloat()
		if err != nil {
			return false, err
		}
		value := accessories.EtherToWei(userInputValue)
		// get gas price & gas limit from env
		gasLimit := uint64(64000)
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Enter your private key:")
		privKeyInput := accessories.UserInputLine()[2:]
		privKey, err := crypto.HexToECDSA(privKeyInput)
		if err != nil {
			return false, err
		}
		// grab pubkey from priv, make pubkey ecdsa
		pubKey := privKey.Public()
		pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("couldnt cast pubkey to ecdsa")
		}
		// define the sender addy
		fromAddress := crypto.PubkeyToAddress(*pubKeyECDSA)
		// pull pending nonce
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		// compile transaction
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      gasLimit,
			To:       &recpient,
			Value:    value,
			Data:     nil,
			V:        nil,
			R:        nil,
			S:        nil,
		})
		// define chain ID
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		// sign transaction
		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privKey)
		if err != nil {
			log.Fatal(err)
		}
		// broadcast transaction
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("tx sent: %s", signedTx.Hash())
		return true, nil
	} else {
		fmt.Println("Transaction cancelled")
		return false, nil
	}
}

func SendTxFlowKeyStore(client *ethclient.Client) (bool, error) {
	acc, ks, err := accounts.GetAccountAndKs()
	fmt.Println("Please enter the address to send Ethereum to:")
	// define recipient address
	toAddress := accessories.UserInputLine()
	if !accounts.ValidEthAddress(toAddress) {
		return false, fmt.Errorf("invalid recepient")
	}
	recipient := common.HexToAddress(toAddress)
	addrBalance := accounts.GetAddressBalance(client, acc.Address)
	fmt.Printf("Your current balance is: %s ETH worth %f USD\n", addrBalance.EthBalance.String(), addrBalance.BalanceUSD)
	fmt.Println("Enter value of transaction being sent in ETH:")
	userInputValue, err := accessories.UserInputFloat()
	if err != nil {
		return false, err
	}
	value := accessories.EtherToWei(userInputValue)
	// get gas price & gas limit from env
	gasLimit := uint64(64000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// pull pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), acc.Address)
	if err != nil {
		log.Fatal(err)
	}
	// compile transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &recipient,
		Value:    value,
		Data:     nil,
		V:        nil,
		R:        nil,
		S:        nil,
	})
	// define chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// sign Tx
	fmt.Printf("Please enter the password to send %s Ethereum to %s.\n", accessories.WeiToEther(value).String(), toAddress)
	txPass := accessories.UserInputLine()
	// define recipient address
	signedTx, err := ks.SignTxWithPassphrase(acc, txPass, tx, chainID)
	if err != nil {
		log.Fatal(err)
	}
	// broadcast transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash())
	return true, nil

}
