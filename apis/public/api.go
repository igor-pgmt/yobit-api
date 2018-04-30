package public

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/igor-pgmt/yobit-api/settings"
)

// API is the Public API that included in the main client.
type API struct{}

// NewAPI creates and returns the Public API to the main client.
func NewAPI() *API {
	return &API{}
}

// sendRequest prepares and sends request to server by calling objective functions and returns the body of response
func (api *API) sendRequest(values *url.Values, link string) ([]byte, error) {

	req, err := api.prepareRequest(values, link)
	settings.Check("Public API.sendRequest() Preparing request", err)

	resp, err := api.sendPost(req)
	settings.Check("Public API.sendRequest() Sending POST request", err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Public API.sendRequest() Reading response body", err)

	return body, err
}

// prepareRequest creates link and prepares request to send
func (api *API) prepareRequest(values *url.Values, link string) (*http.Request, error) {

	requestString := values.Encode()

	req, err := http.NewRequest("POST", link, strings.NewReader(requestString))
	settings.Check("Public API.prepareRequest() Creating request", err)

	return req, err
}

// sendPost sends POST request to the API server
func (api *API) sendPost(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	settings.Check("Public API.sendPost() Doing request", err)

	return resp, err
}
