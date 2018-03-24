package public

import (
	"net/http"

	"github.com/igor-pgmt/yobit-api/settings"
)

// API is the Public API that included in the main client.
type API struct {
}

// NewAPI creates and returns the Public API to the main client.
func NewAPI() *API {
	return &API{}
}

func (api *API) sendRequest(method string, parameters []string) (*http.Response, error) {

	req, _ := api.prepareRequest(method, parameters)

	resp, err := api.sendPost(req)

	return resp, err

}

func (api *API) prepareRequest(method string, parameters []string) (*http.Request, error) {
	apiLink := settings.PublicApiLink + method
	for _, paramName := range parameters {
		apiLink += "/" + paramName
	}
	apiLink += "?limit=2000"
	req, err := http.NewRequest("POST", apiLink, nil)

	return req, err
}

func (api *API) sendPost(req *http.Request) (*http.Response, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	settings.Check(err)

	return resp, err
}
