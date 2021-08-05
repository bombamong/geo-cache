package snowflake

import (
	"fmt"
)

type Config struct {
	Account          string `json:"account"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"database"`
	Schema           string `json:"schema"`
	Warehouse        string `json:"warehouse"`
	NoProxy          string `json:"no_proxy"`
	KeepAlive        string `json:"keep_alive"`
	QueueTimeout     string `json:"queue_timeout"`
	StatementTimeout string `json:"statement_timeout"`
}

func (c *Config) stringify() string {
	return fmt.Sprintf("%s:%s@%s/%s/%s?warehouse=%s&noproxy=%s&client_session_keep_alive=%v&statement_queued_timeout_in_seconds=%v&statement_timeout_in_seconds=%v",
		c.Username,
		c.Password,
		c.Account,
		c.Database,
		c.Schema,
		c.Warehouse,
		c.NoProxy,
		c.KeepAlive,
		c.QueueTimeout,
		c.StatementTimeout,
	)
}
