package go_iexclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

const baseEndpoint = "https://api.iextrading.com/1.0"

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{client}
}

func (c *Client) GetTOPS(symbols []string) ([]*TOPS, error) {
	req := &topsRequest{symbols}
	var result []*TOPS
	err := c.getJSON("/tops", req, &result)
	return result, err
}

type topsRequest struct {
	Symbols []string `url:"symbols,comma,omitempty"`
}

func (c *Client) GetBatch(symbol string,types string, ranges string) (*BatchQuota, error) {
	req := &batchRequest{types, ranges}
	var result *BatchQuota
	err := c.getJSON("/stock/" + symbol + "/batch", req, &result)
	return result, err
}

type batchRequest struct {
	Types string `url:"types"`
	Range string `url:"range"`
}

type BatchQuota struct{
	Quote Batch `json:"quote"`
}

func (c *Client) GetSymbols() (*[]models.Symbol, error) {
	var req interface{}
	var result *[]models.Symbol
	err := c.getJSON("/ref-data/symbols", req, &result)
	return result, err
}


//func (c *Client) GetBatch(symbol string,types string, ranges string) (*BatchQuota, error) {
//	req := &batchRequest{types, ranges}
//	var result *BatchQuota
//	err := c.getJSON("/stock/" + symbol + "/batch", req, &result)
//	return result, err
//}

//type batchRequest struct {
//	Types string `url:"types"`
//	Range string `url:"range"`
//}

func (c *Client) getJSON(route string, request interface{}, response interface{}) error {
	url := c.endpoint(route)

	values, err := query.Values(request)
	if err != nil {
		return err
	}
	queryString := values.Encode()
	if queryString != "" {
		url = url + "?" + queryString
	}

	logrus.Infoln(url)

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%v: %v", resp.Status, string(body))
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(response)
}

func (c *Client) endpoint(route string) string {
	return baseEndpoint + route
}
