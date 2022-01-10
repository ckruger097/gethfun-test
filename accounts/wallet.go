package accounts

import (
	"crypto/ecdsa"
	"fmt"
	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"log"
)

type EthWallet struct {
	WalletName    string
	PrivateKey    []byte
	PublicKey     []byte
	WalletAddress string
	Balance       AddressBalance
}

func LoadWallet(privKey string) EthWallet {
	var newWallet EthWallet
	// TODO: database? or else walletname is kinda useless
	newWallet.WalletName = "wallet"
	privateKeyBytes, err := hexutil.Decode(privKey)
	if err != nil {
		log.Fatal(err)
	}
	newWallet.PrivateKey = privateKeyBytes
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal("couldn't decode private key")
	}
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("could not cast public key to ECDSA")
	}
	// convert pubkey to byte form
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	newWallet.PublicKey = publicKeyBytes
	// address assignment
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	newWallet.WalletAddress = address
	return newWallet
}

func GenerateWallet(newWalletName string) EthWallet {
	var newWallet EthWallet
	// generates wallet
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	newWallet.WalletName = newWalletName
	fmt.Println("Your new wallet name will be:", newWallet.WalletName)
	// converts it to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	newWallet.PrivateKey = privateKeyBytes
	fmt.Println("Your private key is:", hexutil.Encode(newWallet.PrivateKey))
	// pull out pubkey from privkey
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("could not cast public key to ECDSA")
	}
	// convert pubkey to byte form
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	newWallet.PublicKey = publicKeyBytes
	fmt.Println("Public key is:", hexutil.Encode(newWallet.PublicKey))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	newWallet.WalletAddress = address
	fmt.Println("Address to that public key is:", address)
	return newWallet
}
