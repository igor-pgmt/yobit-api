package trade

import (
	"encoding/json"

	"github.com/igor-pgmt/yobit-api/apis/requests"
	"github.com/igor-pgmt/yobit-api/responses"
	"github.com/igor-pgmt/yobit-api/settings"
)

// GetInfo shows info about account's balance.
func (api *API) GetInfo() (responses.GetInfo, error) {

	values := api.createLinkGetInfo()

	body, err := api.sendRequest(values)
	settings.Check("Trade API.Trade() Sending request", err)

	balance := responses.NewBalance()
	err = json.Unmarshal(body, &balance)
	settings.Check("Trade API.GetInfo() Unmarshalling response body", err)

	return balance, err
}

// Trade allows creating new orders.
func (api *API) Trade(t *requests.TradeSettings) (responses.Trade, error) {

	values := api.createLinkTrade(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.Trade() Sending request", err)

	trade := responses.NewTrade()
	err = json.Unmarshal(body, &trade)
	settings.Check("Trade API.Trade() Unmarshalling response body", err)

	return trade, err
}

// ActiveOrders returns list of user's active orders.
func (api *API) ActiveOrders(t *requests.ActiveOrdersSettings) (responses.ActiveOrders, error) {

	values := api.createLinkActiveOrders(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.ActiveOrders() Sending request", err)

	activeOrders := responses.NewActiveOrders()
	err = json.Unmarshal(body, &activeOrders)
	settings.Check("Trade API.ActiveOrders() Unmarshalling response body", err)

	return activeOrders, err
}

// OrderInfo returns detailed information about the chosen order.
func (api *API) OrderInfo(t *requests.OrderInfoSettings) (responses.OrderInfo, error) {

	values := api.createLinkOrderInfo(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.OrderInfo() Sending request", err)

	orderInfo := responses.NewOrderInfo()
	err = json.Unmarshal(body, &orderInfo)
	settings.Check("Trade API.OrderInfo() Unmarshalling response body", err)

	return orderInfo, err
}

// CancelOrder cancells the chosen order.
func (api *API) CancelOrder(t *requests.CancelOrderSettings) (responses.CancelOrder, error) {

	values := api.createLinkCancelOrder(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.CancelOrder() Sending request", err)

	cancelOrder := responses.NewCancelOrder()
	err = json.Unmarshal(body, &cancelOrder)
	settings.Check("Trade API.CancelOrder() Unmarshalling response body", err)

	return cancelOrder, err
}

// TradeHistory returns transaction history.
func (api *API) TradeHistory(t *requests.TradeHistorySettings) (responses.TradeHistory, error) {

	values := api.createLinkTradeHistory(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.TradeHistory() Sending request", err)

	tradeHistory := responses.NewTradeHistory()
	err = json.Unmarshal(body, &tradeHistory)
	settings.Check("Trade API.TradeHistory() Unmarshalling response body", err)

	return tradeHistory, err
}

// GetDepositAddress returns deposit address.
func (api *API) GetDepositAddress(t *requests.GetDepositAddressSettings) (responses.GetDepositAddress, error) {

	values := api.createLinkGetDepositAddress(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.GetDepositAddress() Sending request", err)

	getDepositAddress := responses.NewGetDepositAddress()
	err = json.Unmarshal(body, &getDepositAddress)
	settings.Check("Trade API.GetDepositAddress() Unmarshalling response body", err)

	return getDepositAddress, err
}

// WithdrawCoinsToAddress creates withdrawal request.
func (api *API) WithdrawCoinsToAddress(t *requests.WithdrawCoinsToAddressSettings) (responses.WithdrawCoinsToAddress, error) {

	values := api.createLinkWithdrawCoinsToAddress(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.WithdrawCoinsToAddress() Sending request", err)

	getDepositAddress := responses.NewWithdrawCoinsToAddress()
	err = json.Unmarshal(body, &getDepositAddress)
	settings.Check("Trade API.WithdrawCoinsToAddress() Unmarshalling response body", err)

	return getDepositAddress, err
}

// CreateYobicode allows you to create Yobicodes (coupons).
func (api *API) CreateYobicode(t *requests.CreateYobicodeSettings) (responses.CreateYobicode, error) {

	values := api.createLinkCreateYobicode(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.CreateYobicode() Sending request", err)

	createYobicode := responses.NewCreateYobicode()
	err = json.Unmarshal(body, &createYobicode)
	settings.Check("Trade API.CreateYobicode() Unmarshalling response body", err)

	return createYobicode, err
}

// RedeemYobicode is used to redeem Yobicodes (coupons).
func (api *API) RedeemYobicode(t *requests.RedeemYobicodeSettings) (responses.RedeemYobicode, error) {

	values := api.createLinkRedeemYobicode(t)

	body, err := api.sendRequest(values)
	settings.Check("Trade API.RedeemYobicode() Sending request", err)

	redeemYobicode := responses.NewRedeemYobicode()
	err = json.Unmarshal(body, &redeemYobicode)
	settings.Check("Trade API.RedeemYobicode() Unmarshalling response body", err)

	return redeemYobicode, err
}
