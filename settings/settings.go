package settings

import "io/ioutil"
import "strconv"
import "runtime/debug"
import "log"

// Place your keys at your main.go
var Key string
var Secret string

const (
	PublicApiLink = "https://yobit.net/api/3/"
	TradeApiLink  = "https://yobit.io/tapi/"
)

// Shared function for error checking
func Check(err error) {
	if err != nil {
		log.Println(err, string(debug.Stack()))
	}
}

// Maintenance function for storing counter
func GetNonce(Key string) (nonce int, err error) {
	nonceFileName := "nonce." + Key[0:8] + ".txt"
	nonceBytes, err := ioutil.ReadFile(nonceFileName)
	if err == nil {
		nonce, _ = strconv.Atoi(string(nonceBytes))
	}
	nonce++
	err = ioutil.WriteFile(nonceFileName, []byte(strconv.Itoa(nonce)), 0644)
	Check(err)

	return
}
