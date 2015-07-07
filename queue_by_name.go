package rabbitmq

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) QueueByName(vhost, queue string) *QueueByNameService {
	return &QueueByNameService{c: c, vhost: vhost, queue: queue}
}

type QueueByNameService struct {
	c     *Client
	vhost string
	queue string
}

func (svc *QueueByNameService) Do() (*QueueInfo, error) {
	params := make(url.Values)

	path := fmt.Sprintf("/api/queues/%s/%s", url.QueryEscape(svc.vhost), url.QueryEscape(svc.queue))

	res, err := svc.c.Execute("GET", path, params, nil)
	if err != nil {
		return nil, err
	}

	var ret QueueInfo
	if err := json.Unmarshal(res.Body, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
