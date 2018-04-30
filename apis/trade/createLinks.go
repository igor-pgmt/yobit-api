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
		panic("createLinkTrade Pair hasn't been set")
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
		panic("createLinkActiveOrders Pair hasn't been set")
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
		panic("createLinkOrderInfo Pair hasn't been set")
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
		panic("createLinkCancelOrder Pair hasn't been set")
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
		values.Add("From", strconv.FormatUint(th.From, 10))
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
		panic("createLinkTradeHistory Pair hasn't been set")
	}

	return &values

}

func (api *API) createLinkGetDepositAddress(th *requests.GetDepositAddressSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"GetDepositAddress"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.CoinName != "" {
		values.Add("coinName", th.CoinName)
	} else {
		panic("createLinkGetDepositAddress Pair hasn't been set")
	}
	if th.NeedNew != 0 {
		values.Add("need_new", strconv.FormatUint(th.NeedNew, 10))
	}

	return &values

}

func (api *API) createLinkWithdrawCoinsToAddress(th *requests.WithdrawCoinsToAddressSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"WithdrawCoinsToAddress"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.CoinName != "" {
		values.Add("coinName", th.CoinName)
	} else {
		panic("createLinkWithdrawCoinsToAddress Pair hasn't been set")
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	} else {
		panic("createLinkWithdrawCoinsToAddress Amount hasn't been set")
	}
	if th.Address != "" {
		values.Add("address", th.Address)
	} else {
		panic("createLinkWithdrawCoinsToAddress Address hasn't been set")
	}

	return &values

}

func (api *API) createLinkCreateYobicode(th *requests.CreateYobicodeSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"CreateYobicode"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Currency != "" {
		values.Add("coinName", th.Currency)
	} else {
		panic("createLinkCreateYobicode Currency hasn't been set")
	}
	if th.Amount != 0 {
		values.Add("amount", strconv.FormatFloat(th.Amount, 'f', -1, 64))
	} else {
		panic("createLinkCreateYobicode Amount hasn't been set")
	}

	return &values

}

func (api *API) createLinkRedeemYobicode(th *requests.RedeemYobicodeSettings) *url.Values {

	nonce, err := api.GetNonce(settings.Key)
	settings.Check("Trade API.prepareRequest() Getting nonce", err)

	values := url.Values{
		"method": []string{"RedeemYobicode"},
		"nonce":  []string{strconv.Itoa(nonce)},
	}

	if th.Coupon != "" {
		values.Add("coupon", th.Coupon)
	} else {
		panic("createLinkRedeemYobicode Currency hasn't been set")
	}

	return &values

}
