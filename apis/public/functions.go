package public

import (
	"encoding/json"
	"io/ioutil"

	"github.com/igor-pgmt/yobit-api/responses"
	"github.com/igor-pgmt/yobit-api/settings"
)

// Info returns information about server time and active pairs.
func (api *API) Info() (responses.Info, error) {
	info := responses.NewInfo()
	resp, err := api.sendRequest("info", []string{})
	settings.Check("Public API.Info() Getting response", err)
	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Public API.Info() Reading response body", err)
	defer resp.Body.Close()

	err = json.Unmarshal(body, &info)
	settings.Check("Public API.Info() Unmarshalling response body", err)

	return info, err
}

// Depth returns information about lists of active orders for selected pairs.
func (api *API) Depth(pair string) (responses.Depth, error) {

	resp, err := api.sendRequest("depth", []string{pair})
	settings.Check("Public API.Depth() Getting response", err)
	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Public API.Depth() Reading response body", err)
	defer resp.Body.Close()

	depth := responses.NewDepth()
	err = json.Unmarshal(body, &depth.PairData)
	settings.Check("Public API.Depth() Unmarshalling response body", err)
	if err != nil {
		err = json.Unmarshal(body, &depth)
		settings.Check("Public API.Depth() Unmarshalling response error", err)
	}

	return depth, err
}

// Trades returns information about the last transactions of selected pairs.
func (api *API) Trades(pair string) (responses.Trades, error) {

	resp, err := api.sendRequest("trades", []string{pair})
	settings.Check("Public API.Trades() Getting response", err)
	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Public API.Trades() Reading response body", err)
	defer resp.Body.Close()

	trades := responses.NewTrades()
	err = json.Unmarshal(body, &trades.PairData)
	settings.Check("Public API.Trades() Unmarshalling response body", err)
	if err != nil {
		err = json.Unmarshal(body, &trades)
		settings.Check("Public API.Depth() Unmarshalling response error", err)
	}
	return trades, err
}
