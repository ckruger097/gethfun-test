package accounts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
	"math"
	"math/big"
)

type AddressBalance struct {
	rawBalance *big.Int
	EthBalance *big.Float
	BalanceUSD float64
}

func GetAddressBalance(client *ethclient.Client, addy common.Address) AddressBalance {
	var addrBal AddressBalance
	var err error
	addrBal.rawBalance, err = client.BalanceAt(context.Background(), addy, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(addrBal.rawBalance.String())
	addrBal.EthBalance = new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	ebfl, _ := addrBal.EthBalance.Float64()
	epfl, _ := getEthPrice().Float64()
	addrBal.BalanceUSD = ebfl * epfl

	return addrBal
}

func getEthPrice() *big.Float {
	cg := gecko.NewClient(nil)
	price, err := cg.SimpleSinglePrice("ethereum", "usd")
	if err != nil {
		log.Fatal(err)
	}
	return big.NewFloat(float64(price.MarketPrice))
}

func PrintBalanceInfo(client *ethclient.Client, address common.Address) {
	addressBalance := GetAddressBalance(client, address)
	fmt.Println("Your raw Balance is:", addressBalance.rawBalance)
	fmt.Println("Your value is:", addressBalance.EthBalance, "ETH")
	fmt.Println("This Balance is currently worth:$", addressBalance.BalanceUSD, "USD")
}
