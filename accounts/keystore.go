package accounts

import (
	"fmt"
	"gethfun/accessories"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CreateKs(newPassword string) (accounts.Account, error) {
	ks := keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.StandardScryptP)
	//am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	account, err := ks.NewAccount(newPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your new address:", account.Address.Hex())
	_, err = ks.Export(account, newPassword, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	return account, nil
}

func ImportKs() (accounts.Account, error){
	var files []string
	fmt.Println("Files currently in keystore:")
	root := "./keystores"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
	fmt.Println("Which keystore do you want to unlock?")
	fileName := accessories.UserInputLine()
	file := fmt.Sprintf("./keystores/%s", fileName)
	ks := keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.StandardScryptP)
	//am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("What's the password?")
	userPass := accessories.UserInputLine()
	impAcc, _ := ks.Import(jsonBytes, userPass, userPass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Your address is:", impAcc.Address.Hex())
	fmt.Println("Your key data is:", jsonBytes)
	return impAcc, nil
}
