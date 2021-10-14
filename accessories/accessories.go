package accessories

import (
	"bufio"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
	"log"
	"math/big"
	"os"
	"strings"
)

func UserInputLine() string {
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	response = strings.TrimSuffix(response, "\n")
	return response
}

func UserInputInteger() (*big.Int, error) {
	w := new(big.Int)
	_, err := fmt.Scan(w)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("bad int")
	}
	return w, nil
}

func UserInputFloat() (*big.Float, error) {
	w := new(big.Float)
	_, err := fmt.Scan(w)
	if err != nil {
		return big.NewFloat(0), fmt.Errorf("bad float")
	}
	return w, nil
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func WeiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
