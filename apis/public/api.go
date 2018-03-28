package public

import (
	"net/http"

	"github.com/igor-pgmt/yobit-api/settings"
)

// API is the Public API that included in the main client.
type API struct{}

// NewAPI creates and returns the Public API to the main client.
func NewAPI() *API {
	return &API{}
}

func (api *API) sendRequest(method string, parameters []string) (*http.Response, error) {

	req, err := api.prepareRequest(method, parameters)
	settings.Check("Public API.sendRequest() Preparing request", err)

	resp, err := api.sendPost(req)
	settings.Check("Public API.sendRequest() Sending POST request", err)

	return resp, err

}

func (api *API) prepareRequest(method string, parameters []string) (*http.Request, error) {
	apiLink := settings.PublicApiLink + method
	for _, paramName := range parameters {
		apiLink += "/" + paramName
	}
	// TODO: next line is a hardcode. Need to place this parameter into request structs.
	apiLink += "?limit=2000"
	req, err := http.NewRequest("POST", apiLink, nil)
	settings.Check("Public API.prepareRequest() Creating request", err)

	return req, err
}

func (api *API) sendPost(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	settings.Check("Public API.sendPost() Doing request", err)

	return resp, err
}
