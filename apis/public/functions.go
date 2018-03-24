package public

import (
	"encoding/json"
	"io/ioutil"

	"github.com/igor-pgmt/yobit-api/responses"
	"github.com/igor-pgmt/yobit-api/settings"
)

// Info shows information about server time and active pairs.
func (api *API) Info() (responses.Info, error) {
	info := responses.NewInfo()
	resp, err := api.sendRequest("info", []string{})
	settings.Check(err)
	body, err := ioutil.ReadAll(resp.Body)
	settings.Check(err)
	defer resp.Body.Close()

	err = json.Unmarshal(body, &info)
	settings.Check(err)

	return info, err
}

// Depth returns information about lists of active orders for selected pairs.
func (api *API) Depth(pair string) (responses.Depth, error) {
	depth := responses.NewDepth()
	fullData := make(map[string]responses.PData)

	resp, err := api.sendRequest("depth", []string{pair})
	settings.Check(err)
	body, err := ioutil.ReadAll(resp.Body)
	settings.Check(err)
	defer resp.Body.Close()

	err = json.Unmarshal(body, &fullData)

	settings.Check(err)
	depth.PairData = fullData

	return depth, err
}
