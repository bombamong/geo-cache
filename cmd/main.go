package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bombamong/geo-cache/pkg/cache"
	"github.com/bombamong/geo-cache/pkg/snowflake"

	"github.com/gin-gonic/gin"
)

func init() {
	bindSecrets()
}

type Config struct {
	SnowConfig snowflake.Config
}

var config = Config{}

func bindSecrets() {
	// SNOW config
	snowAccount := os.Getenv("SNOWSQL_ACCOUNT")
	snowRegion := os.Getenv("SNOWSQL_REGION")
	snowUser := os.Getenv("SNOWSQL_USER")
	snowDB := os.Getenv("SNOWSQL_DATABASE")
	snowPW := os.Getenv("SNOWSQL_PWD")
	snowSchema := os.Getenv("SNOWSQL_SCHEMA")
	snowWarehouse := os.Getenv("SNOWSQL_WAREHOUSE")
	snowNoProxy := os.Getenv("SNOWSQL_NO_PROXY")
	snowKeepAlive := os.Getenv("SNOWSQL_CLIENT_SESSION_KEEP_ALIVE")
	snowQueueTO := os.Getenv("SNOWSQL_QUEUE_TIMEOUT")
	snowStmtTO := os.Getenv("SNOWSQL_STATEMENT_TIMEOUT")
	snowConfig := snowflake.Config{
		Account:          fmt.Sprintf("%s.%s", snowAccount, snowRegion),
		Username:         snowUser,
		Password:         snowPW,
		Database:         snowDB,
		Schema:           snowSchema,
		Warehouse:        snowWarehouse,
		NoProxy:          snowNoProxy,
		KeepAlive:        snowKeepAlive,
		QueueTimeout:     snowQueueTO,
		StatementTimeout: snowStmtTO,
	}
	config.SnowConfig = snowConfig
}

func main() {
	snow := snowflake.NewSnowflake(config.SnowConfig)
	cash := cache.NewCacheLayer(snow)
	ctx := context.Background()
	err := cash.FillRows(ctx)
	if err != nil {
		log.Println(err)
	}

	r := gin.Default()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	r.GET("/:merchant", func(c *gin.Context) {
		p, ok := c.Params.Get("merchant")
		var _ = p

		if !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		cf := func(rd cache.RawData) bool {
			return true
		}
		md := cash.Query(cf)
		log.Println(md)
		c.JSON(http.StatusOK, md)
	})

	fmt.Println("Start server")
	err = server.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
