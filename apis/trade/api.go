package trade

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/igor-pgmt/yobit-api/settings"
)

// API is the Trade API that included in the main client
type API struct {
	VirtualNonce bool // is for saving to file or not (false = to file by edfault)
	Nonce        int  // current nonce parameter
}

// NewAPI creates and returns the Trade API to the main client
func NewAPI() *API {
	return &API{}
}

// sendRequest prepares and sends request to server by calling objective functions and returns the body of response
func (api *API) sendRequest(values *url.Values) ([]byte, error) {

	req, err := api.prepareRequest(values)
	settings.Check("Trade API.sendRequest() Preparing request", err)

	resp, err := api.sendPost(req)
	settings.Check("Trade API.sendRequest() Getting response", err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Trade API.sendRequest() Reading response body", err)

	return body, err
}

// prepareRequest creates link and prepares request to send
func (api *API) prepareRequest(values *url.Values) (*http.Request, error) {

	requestString := values.Encode()

	sign := hmac.New(sha512.New, []byte(settings.Secret))
	sign.Write([]byte(requestString))

	req, err := http.NewRequest("POST", settings.TradeApiLink, strings.NewReader(requestString))
	settings.Check("Trade API.prepareRequest() Creating request", err)
	req.Header.Add("Key", settings.Key)
	req.Header.Add("Sign", hex.EncodeToString(sign.Sum(nil)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err
}

// sendPost sends POST request to the API server
func (api *API) sendPost(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	settings.Check("Trade API.sendPost() Doing request", err)

	return resp, err
}

// GetNonce is a maintenance function for getting and storing nonce counter
func (api *API) GetNonce(Key string) (int, error) {
	nonceFileName := "nonce." + Key[0:8] + ".txt"
	if api.VirtualNonce != true {
		nonceBytes, err := ioutil.ReadFile(nonceFileName)
		if err == nil {
			api.Nonce, _ = strconv.Atoi(string(nonceBytes))
		}
	}
	api.Nonce++
	err := api.WriteNonce(api.Nonce, nonceFileName)
	settings.Check("WriteNonce", err)
	return api.Nonce, err
}

// WriteNonce writes nonce to the file
func (api *API) WriteNonce(nonce int, nonceFileName string) (err error) {
	err = ioutil.WriteFile(nonceFileName, []byte(strconv.Itoa(nonce)), 0644)
	settings.Check("Get Nonce write file error", err)
	return
}
