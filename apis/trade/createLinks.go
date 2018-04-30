package trade

import (
	"net/url"
	"strconv"

	"github.com/igor-pgmt/yobit-api/apis/requests"
	"github.com/igor-pgmt/yobit-api/settings"
)

func (api *API) createLinkGetInfo() *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"getInfo"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	return &values

}

func (api *API) createLinkTrade(th *requests.TradeSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"Trade"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		panic("createLinkTrade Pair not set")
	}
	if th.Type != "" {
		values.Add("type", th.Type)
	}
	if th.Rate != 0 {
		values.Add("rate", strconv.FormatFloat(th.Rate, 'f', -1, 64))
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	}

	return &values

}

func (api *API) createLinkActiveOrders(th *requests.ActiveOrdersSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"ActiveOrders"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		panic("createLinkActiveOrders Pair not set")
	}

	return &values

}

func (api *API) createLinkOrderInfo(th *requests.OrderInfoSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"OrderInfo"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.OrderID != 0 {
		values.Add("order_id", strconv.FormatUint(th.OrderID, 10))
	} else {
		panic("createLinkOrderInfo Pair not set")
	}

	return &values

}

func (api *API) createLinkCancelOrder(th *requests.CancelOrderSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"CancelOrder"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.OrderID != 0 {
		values.Add("order_id", strconv.FormatUint(th.OrderID, 10))
	} else {
		panic("createLinkCancelOrder Pair not set")
	}

	return &values

}

func (api *API) createLinkTradeHistory(th *requests.TradeHistorySettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"TradeHistory"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.From != 0 {
		values.Add("From", string(th.From))
	}
	if th.Count != 0 {
		values.Add("Count", strconv.FormatUint(th.Count, 10))
	}
	if th.FromID != 0 {
		values.Add("FromID", strconv.FormatUint(th.FromID, 10))
	}
	if th.EndID != 0 {
		values.Add("EndID", strconv.FormatUint(th.EndID, 10))
	}
	if th.Order != "" {
		values.Add("Order", th.Order)
	}
	if th.Since != 0 {
		values.Add("Since", strconv.FormatUint(th.Since, 10))
	}
	if th.End != 0 {
		values.Add("End", strconv.FormatUint(th.End, 10))
	}
	if th.Pair != "" {
		values.Add("pair", th.Pair)
	} else {
		panic("createLinkTradeHistory Pair not set")
	}

	return &values

}
