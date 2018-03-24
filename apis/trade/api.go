package trade

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/igor-pgmt/yobit-api/settings"
)

// API is the Trade API that included in the main client
type API struct{}

// NewAPI creates and returns the Trade API to the main client
func NewAPI() *API {
	return &API{}
}

func (api *API) sendRequest(method string, parameters map[string]interface{}) (*http.Response, error) {

	req, _ := api.prepareRequest(method, parameters)
	resp, err := api.sendPost(req)

	return resp, err
}

func (api *API) sendPost(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	settings.Check(err)

	return resp, err
}

func (api *API) prepareRequest(method string, parameters map[string]interface{}) (*http.Request, error) {
	nonce, err := settings.GetNonce(settings.Key)
	settings.Check(err)

	values := url.Values{
		"method": []string{method},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	for paramName, paramValue := range parameters {

		paramName = strings.ToLower(paramName)
		if s, ok := paramValue.(string); ok {
			values.Add(paramName, s)
		} else {
			values.Add(paramName, strconv.FormatFloat(paramValue.(float64), 'f', -1, 64))
		}
	}

	requestString := values.Encode()
	fmt.Println("requestString", requestString)

	sign := hmac.New(sha512.New, []byte(settings.Secret))
	sign.Write([]byte(requestString))

	req, err := http.NewRequest("POST", settings.TradeApiLink, strings.NewReader(requestString))
	req.Header.Add("Key", settings.Key)
	req.Header.Add("Sign", hex.EncodeToString(sign.Sum(nil)))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, err

}
