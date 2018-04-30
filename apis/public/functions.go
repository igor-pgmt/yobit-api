package public

import (
	"encoding/json"

	"github.com/igor-pgmt/yobit-api/apis/requests"
	"github.com/igor-pgmt/yobit-api/responses"
	"github.com/igor-pgmt/yobit-api/settings"
)

// Info returns information about server time and active pairs.
func (api *API) Info() (responses.Info, error) {

	values, link := api.createLinkInfo()

	body, err := api.sendRequest(values, link)
	settings.Check("Public API.Info() Getting response body", err)

	info := responses.NewInfo()
	err = json.Unmarshal(body, &info)
	settings.Check("Public API.Info() Unmarshalling response body", err)

	return info, err
}

// Ticker provides statistic data for the last 24 hours.
func (api *API) Ticker(t *requests.TickerSettings) (responses.Ticker, error) {

	values, link := api.createLinkTicker(t)

	body, err := api.sendRequest(values, link)
	settings.Check("Public API.Ticker() Getting response body", err)

	ticker := responses.NewTicker()
	err = json.Unmarshal(body, &ticker.PairData)
	settings.Check("Public API.Ticker() Unmarshalling response body", err)

	return ticker, err
}

// Depth returns information about lists of active orders for selected pairs.
func (api *API) Depth(t *requests.DepthSettings) (responses.Depth, error) {

	values, link := api.createLinkDepth(t)

	body, err := api.sendRequest(values, link)
	settings.Check("Public API.Depth() Getting response body", err)

	depth := responses.NewDepth()
	err = json.Unmarshal(body, &depth.PairData)
	settings.Check("Public API.Depth() Unmarshalling response body", err)

	return depth, err
}

// Trades returns information about the last transactions of selected pairs.
func (api *API) Trades(t *requests.TradesSettings) (responses.Trades, error) {

	values, link := api.createLinkTrades(t)

	body, err := api.sendRequest(values, link)
	settings.Check("Public API.Trades() Getting response body", err)

	trades := responses.NewTrades()
	err = json.Unmarshal(body, &trades.PairData)
	settings.Check("Public API.Trades() Unmarshalling response body", err)

	return trades, err
}
