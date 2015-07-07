package rabbitmq

import (
	"encoding/json"
	"net/url"
)

func (c *Client) Exchanges() *ExchangesService {
	return &ExchangesService{c: c}
}

type ExchangesService struct {
	c *Client
}

func (svc *ExchangesService) Do() ([]*ExchangeInfo, error) {
	params := make(url.Values)

	res, err := svc.c.Execute("GET", "/api/exchanges", params, nil)
	if err != nil {
		return nil, err
	}

	ret := make([]*ExchangeInfo, 0)
	if err := json.Unmarshal(res.Body, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

type ExchangeInfo struct {
	Name         string                 `json:"name"`
	VHost        string                 `json:"vhost"`
	Type         string                 `json:"type"`
	Durable      bool                   `json:"durable"`
	AutoDelete   bool                   `json:"auto_delete"`
	Internal     bool                   `json:"internal"`
	Arguments    map[string]interface{} `json:"arguments"`
	MessageStats *ExchangeMessageStats  `json:"message_stats"`
}

type ExchangeMessageStats struct {
	PublishIn         int64         `json:"publish_in"`
	PublishInDetails  *StatsDetails `json:"publish_in_details"`
	PublishOut        int64         `json:"publish_out"`
	PublishOutDetails *StatsDetails `json:"publish_out_details"`
}
