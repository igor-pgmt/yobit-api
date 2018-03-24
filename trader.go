package trader

import "github.com/igor-pgmt/yobit-api/apis/trade"
import "github.com/igor-pgmt/yobit-api/apis/public"

// Client is a main struct for executing all the operations of the API
type Client struct {
	Public *public.API
	Trade  *trade.API
}

// NewClient is a constructor for the Client
func NewClient() *Client {
	client := &Client{}
	client.Trade = trade.NewAPI()
	client.Public = public.NewAPI()
	return client
}
