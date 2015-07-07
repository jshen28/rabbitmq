package rabbitmq

import (
	"encoding/json"
	"net/url"
)

func (c *Client) Overview() *OverviewService {
	return &OverviewService{c: c}
}

type OverviewService struct {
	c *Client
}

func (svc *OverviewService) Do() (*OverviewResponse, error) {
	params := make(url.Values)

	res, err := svc.c.Execute("GET", "/api/overview", params, nil)
	if err != nil {
		return nil, err
	}

	var ret OverviewResponse
	if err := json.Unmarshal(res.Body, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

type OverviewResponse struct {
	ManagementVersion      string          `json:"management_version"`
	RatesMode              string          `json:"rates_mode"`
	RabbitMQVersion        string          `json:"rabbitmq_version"`
	ClusterName            string          `json:"cluster_name"`
	Node                   string          `json:"node"`
	ErlangVersion          string          `json:"erlang_version"`
	ErlangFullVersion      string          `json:"erlang_full_version"`
	MessageStats           *MessageStats   `json:"message_stats"`
	QueueTotals            *QueueStats     `json:"queue_totals"`
	ObjectTotals           *ObjectStats    `json:"object_totals"`
	StatisticsDBNode       string          `json:"statistics_db_node"`
	StatisticsDBEventQueue int             `json:"statistics_db_event_queue"`
	ExchangeTypes          []*ExchangeType `json:"exchange_types"`
	Listeners              []*Listener     `json:"listeners"`
	Contexts               []*Context      `json:"contexts"`
}

type MessageStats struct {
	Publish           int64         `json:"publish"`
	PublishDetails    *StatsDetails `json:"publish_details"`
	Ack               int64         `json:"ack"`
	AckDetails        *StatsDetails `json:"ack_details"`
	DeliverGet        int64         `json:"deliver_get"`
	DeliverGetDetails *StatsDetails `json:"deliver_get_details"`
	Redeliver         int64         `json:"redeliver"`
	RedeliverDetails  *StatsDetails `json:"redeliver_details"`
	Deliver           int64         `json:"deliver"`
	DeliverDetails    *StatsDetails `json:"deliver_details"`
	Get               int64         `json:"get"`
	GetDetails        *StatsDetails `json:"get_details"`
	DiskReads         int64         `json:"disk_reads"`
	DiskReadsDetails  *StatsDetails `json:"disk_reads_details"`
	DiskWrites        int64         `json:"disk_writes"`
	DiskWritesDetails *StatsDetails `json:"disk_writes_details"`
}

type QueueStats struct {
	Messages                      int64         `json:"messages"`
	MessagesDetails               *StatsDetails `json:"messages_details"`
	MessagesReady                 int64         `json:"messages_ready"`
	MessagesReadyDetails          *StatsDetails `json:"messages_ready_details"`
	MessagesUnacknowledged        int64         `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails *StatsDetails `json:"messages_unacknowledged_details"`
}

type ObjectStats struct {
	Consumers   int `json:"consumers"`
	Queues      int `json:"queues"`
	Exchanges   int `json:"exchanges"`
	Connections int `json:"connections"`
	Channels    int `json:"channels"`
}

type StatsDetails struct {
	Rate float64 `json:"rate"`
}

type ExchangeType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type Listener struct {
	Node      string `json:"node"`
	Protocol  string `json:"protocol"`   // e.g. amqp, mqtt, or stomp
	IPAddress string `json:"ip_address"` // e.g. 127.0.0.1 or ::
	Port      int    `json:"port"`
}

type Context struct {
	Node        string `json:"node"`
	Description string `json:"description"`
	Path        string `json:"path"`
	Port        string `json:"port"` // it's a string here
}
