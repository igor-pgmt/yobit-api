# Golang wrapper for Yobit API

This wrapper allows to trade on [Yobit.io](https://Yobit.io)

## Available methods:

Public API:
  - [X] info
  - [X] depth
  - [X] ticker
  - [X] trades

Trade API:
  - [X] getInfo
  - [X] ActiveOrders
  - [X] Trade
  - [X] OrderInfo
  - [X] CancelOrder
  - [X] TradeHistory
  - [X] GetDepositAddress
  - [X] WithdrawCoinsToAddress
  - [X] CreateYobicode
  - [X] RedeemYobicode

### How to use
Download the repository from github to your src folder:
```bash
go get -v github.com/igor-pgmt/yobit-api
```
Import package into your project (add these lines into your code):
```go
import	"github.com/igor-pgmt/yobit-api" // API client
import	"github.com/igor-pgmt/yobit-api/settings" // Settings for API keys
import  "github.com/igor-pgmt/yobit-api/apis/requests" // Requests settings for API
```
Set Your keys:
```go
	settings.Key = "0F353F10C44250F49C499ACDD0D5FDD3"
	settings.Secret = "2c0019b0578b2059705830f9b3a527a1"
```

### Examples
Your main.go:
```go
package main

import (
	"fmt"
	"github.com/igor-pgmt/yobit-api"
	"github.com/igor-pgmt/yobit-api/settings"
	"github.com/igor-pgmt/yobit-api/apis/requests"
)

func main() {
	// Set your keys
	settings.Key = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	settings.Secret = "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	// Create new client
	client := trader.NewClient()
	// Pair
	pair := "doge_btc"

	// Public API

	fmt.Println("/---------------------------------/")
	
	info, err := client.Public.Info()
	if err != nil {
		fmt.Println("Cannot load info")
	}

	fmt.Println("Public API: Info", pair, info.Pairs)       // all the pairs
	fmt.Println("Public API: Info", pair, info.Pairs[pair]) // current pair

	fmt.Println("/---------------------------------/")

	tis := &requests.TickerSettings{
		Pair: pair,
	}

	ticker, err := client.Public.Ticker(tis)
	if err != nil {
		fmt.Println("Cannot load ticker")
	}
	tdata := ticker.PairData[pair]

	fmt.Println("Public API: Ticker", pair, tdata)

	fmt.Println("/---------------------------------/")

	des := &requests.DepthSettings{
		Pair: pair,
	}

	depth, err := client.Public.Depth(des)
	if err != nil {
		fmt.Println("Cannot load depth")
	}
	data := depth.PairData[pair].Asks[:5] // show just five orders

	fmt.Println("Public API: Depth", pair, data)

	fmt.Println("/---------------------------------/")

	tss := &requests.TradesSettings{
		Pair: pair,
	}

	trades, err := client.Public.Trades(tss)
	if err != nil {
		fmt.Println("Cannot load trades")
	}
	tradesData := trades.PairData[pair][:5] // show just five transactions

	fmt.Println("Public API: Trades", pair, tradesData)

	fmt.Println("/---------------------------------/")

	// Trade API

	tradeInfo, err := client.Trade.GetInfo()
	if err != nil {
		fmt.Println("Cannot load tradeInfo")
	}
	fmt.Println("Trade API: GetInfo", tradeInfo)

	fmt.Println("/---------------------------------/")

	aos := &requests.ActiveOrdersSettings{
		Pair: pair,
	}

	activeOrders, err := client.Trade.ActiveOrders(aos)
	if err != nil {
		fmt.Println("Cannot load ActiveOrders")
	}
	fmt.Println("Trade API: ActiveOrders", activeOrders)

	fmt.Println("/---------------------------------/")

	ts := &requests.TradeSettings{
		Pair:   pair,
		Type:   "sell",       // or "buy"
		Rate:   0.00001339,   // e.g. desired "price"
		Amount: 444.44444444, // amount of your currency
	}

	trade, err := client.Trade.Trade(ts)
	if err != nil {
		fmt.Println("Cannot make Trade")
	}
	fmt.Println("Trade API: Trade", trade)

	fmt.Println("/---------------------------------/")

	ois := &requests.OrderInfoSettings{
		OrderID: 123456789, // <--- YOUR ORDER ID HERE
	}

	orderInfo, err := client.Trade.OrderInfo(ois)
	if err != nil {
		fmt.Println("Cannot load OrderInfo")
	}
	fmt.Println("Trade API: Order Info", orderInfo)

	fmt.Println("/---------------------------------/")

	cos := &requests.CancelOrderSettings{
		OrderID: 123456789, // <--- YOUR ORDER ID HERE
	}

	cancelOrder, err := client.Trade.CancelOrder(cos)
	if err != nil {
		fmt.Println("Cannot load OrderInfo")
	}
	fmt.Println("Trade API: CancelOrder", cancelOrder)

	fmt.Println("/---------------------------------/")

	ths := &requests.TradeHistorySettings{
		Pair: pair,
	}

	TradeHistory, err := client.Trade.TradeHistory(ths)
	if err != nil {
		fmt.Println("Cannot load TradeHistory")
	}
	fmt.Println("Trade API: TradeHistory", TradeHistory)

	fmt.Println("/---------------------------------/")

	gdas := &requests.GetDepositAddressSettings{
		CoinName: "BTC",
	}

	GetDepositAddress, err := client.Trade.GetDepositAddress(gdas)
	if err != nil {
		fmt.Println("Cannot load GetDepositAddress")
	}
	fmt.Println("Trade API: GetDepositAddress", GetDepositAddress)

	fmt.Println("/---------------------------------/")

	wctas := &requests.WithdrawCoinsToAddressSettings{
		CoinName: "BTC",
		Amount:   1.00001023,
		Address:  "addr", // <<-- SET ADDRESS HERE
	}

	WithdrawCoinsToAddress, err := client.Trade.WithdrawCoinsToAddress(wctas)
	if err != nil {
		fmt.Println("Cannot load WithdrawCoinsToAddress")
	}
	fmt.Println("Trade API: WithdrawCoinsToAddress", WithdrawCoinsToAddress)

	fmt.Println("/---------------------------------/")

	cys := &requests.CreateYobicodeSettings{
		Currency: "BTC",
		Amount:   1.00001023,
	}

	CreateYobicode, err := client.Trade.CreateYobicode(cys)
	if err != nil {
		fmt.Println("Cannot load CreateYobicode")
	}
	fmt.Println("Trade API: CreateYobicode", CreateYobicode)

	fmt.Println("/---------------------------------/")

	rys := &requests.RedeemYobicodeSettings{
		Coupon: "YOBITUZ0HHSTB...OQX3H01BTC", // <<-- SET YOUR COUPON HERE
	}

	RedeemYobicode, err := client.Trade.RedeemYobicode(rys)
	if err != nil {
		fmt.Println("Cannot load RedeemYobicode")
	}
	fmt.Println("Trade API: RedeemYobicode", RedeemYobicode)

	fmt.Println("/---------------------------------/")
}
```

