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
	"strings"
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
	//am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return account, nil
}

func GetAccountAndKs() (accounts.Account, *keystore.KeyStore, error) {
	var files []string
	fmt.Println("Files currently in keystore:")
	root := "./keystores"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(strings.TrimPrefix(file, "keystores/"))
	}
	fmt.Println("Which keystore do you want to unlock?")
	fileName := accessories.UserInputLine()
	file := fmt.Sprintf("./keystores/%s", fileName)
	ks := keystore.NewKeyStore("./keystores", keystore.StandardScryptN, keystore.StandardScryptP)
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
	//am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	return impAcc, ks, nil
}
