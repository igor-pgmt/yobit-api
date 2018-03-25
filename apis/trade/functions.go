package trade

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/fatih/structs"
	"github.com/igor-pgmt/yobit-api/apis/requests"
	"github.com/igor-pgmt/yobit-api/responses"
	"github.com/igor-pgmt/yobit-api/settings"
)

// GetInfo shows info about account's balance
func (api *API) GetInfo() (responses.GetInfo, error) {
	balance := responses.NewBalance()
	resp, err := api.sendRequest("getInfo", map[string]interface{}{})
	settings.Check("Trade API.GetInfo() Getting response", err)

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Trade API.GetInfo() Reading response body", err)
	defer resp.Body.Close()

	err = json.Unmarshal(body, &balance)
	settings.Check("Trade API.GetInfo() Unmarshalling response body", err)

	return balance, err
}

// Trade allows creating new orders
func (api *API) Trade(t *requests.TradeSettings) (responses.Trade, error) {
	trade := responses.NewTrade()
	tradeMap := structs.Map(t)
	resp, err := api.sendRequest("Trade", tradeMap)
	settings.Check("Trade API.Trade() Getting response", err)

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Trade API.Trade() Reading response body", err)
	defer resp.Body.Close()

	err = json.Unmarshal(body, &trade)
	settings.Check("Trade API.Trade() Unmarshalling response body", err)

	return trade, err
}

// ActiveOrders returns list of user's active orders
func (api *API) ActiveOrders(t *requests.ActiveOrdersSettings) (responses.ActiveOrders, error) {
	activeOrdersMap := structs.Map(t)
	resp, err := api.sendRequest("ActiveOrders", activeOrdersMap)
	settings.Check("Trade API.ActiveOrders() Getting response", err)

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Trade API.ActiveOrders() Reading response body", err)
	defer resp.Body.Close()
	fmt.Println("BODY", string(body))

	activeOrders := responses.NewActiveOrders()
	err = json.Unmarshal(body, &activeOrders)
	settings.Check("Trade API.ActiveOrders() Unmarshalling response body", err)

	return activeOrders, err
}

// OrderInfo returns detailed information about the chosen order
func (api *API) OrderInfo(t *requests.OrderInfoSettings) (responses.OrderInfo, error) {
	orderInfoMap := structs.Map(t)
	// TODO: 2 next lines seem not very good. Need to find a way to convert struct's variables into needed names
	orderInfoMap["order_id"] = orderInfoMap["OrderID"]
	delete(orderInfoMap, "OrderID")

	resp, err := api.sendRequest("OrderInfo", orderInfoMap)
	settings.Check("Trade API.OrderInfo() Getting response", err)

	body, err := ioutil.ReadAll(resp.Body)
	settings.Check("Trade API.OrderInfo() Reading response body", err)
	defer resp.Body.Close()
	fmt.Println("BODY", string(body))

	orderInfo := responses.NewOrderInfo()
	err = json.Unmarshal(body, &orderInfo)
	settings.Check("Trade API.OrderInfo() Unmarshalling response body", err)

	return orderInfo, err
}
