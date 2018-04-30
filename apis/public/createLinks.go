package public

import (
	"net/url"
	"strconv"

	"github.com/igor-pgmt/yobit-api/apis/requests"
	"github.com/igor-pgmt/yobit-api/settings"
)

func (api *API) createLinkInfo() (*url.Values, string) {

	values := url.Values{}
	link := settings.PublicApiLink + "info"

	return &values, link

}

func (api *API) createLinkTicker(th *requests.TickerSettings) (*url.Values, string) {

	values := url.Values{}
	link := settings.PublicApiLink + "ticker" + "/" + th.Pair

	return &values, link

}

func (api *API) createLinkDepth(th *requests.DepthSettings) (*url.Values, string) {

	values := url.Values{}

	if th.Limit != 0 {
		values.Add("limit", strconv.FormatUint(th.Limit, 10))
	}

	link := settings.PublicApiLink + "depth" + "/" + th.Pair

	return &values, link

}

func (api *API) createLinkTrades(th *requests.TradesSettings) (*url.Values, string) {

	values := url.Values{}

	if th.Limit != 0 {
		values.Add("limit", strconv.FormatUint(th.Limit, 10))
	}

	link := settings.PublicApiLink + "trades" + "/" + th.Pair

	return &values, link

}
