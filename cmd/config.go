package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/bombamong/geo-cache/pkg/snowflake"
// )

// type Config struct {
// 	SnowConfig snowflake.Config
// }

// var config = Config{}

// func bindSecrets() {
// 	// SNOW config
// 	snowAccount := os.Getenv("SNOWSQL_ACCOUNT")
// 	snowRegion := os.Getenv("SNOWSQL_REGION")
// 	snowUser := os.Getenv("SNOWSQL_USER")
// 	snowDB := os.Getenv("SNOWSQL_DATABASE")
// 	snowPW := os.Getenv("SNOWSQL_PWD")
// 	snowSchema := os.Getenv("SNOWSQL_SCHEMA")
// 	snowWarehouse := os.Getenv("SNOWSQL_WAREHOUSE")
// 	snowNoProxy := os.Getenv("SNOWSQL_NO_PROXY")
// 	snowKeepAlive := os.Getenv("SNOWSQL_CLIENT_SESSION_KEEP_ALIVE")
// 	snowQueueTO := os.Getenv("SNOWSQL_QUEUE_TIMEOUT")
// 	snowStmtTO := os.Getenv("SNOWSQL_STATEMENT_TIMEOUT")
// 	snowConfig := snowflake.Config{
// 		Account:          fmt.Sprintf("%s.%s", snowAccount, snowRegion),
// 		Username:         snowUser,
// 		Password:         snowPW,
// 		Database:         snowDB,
// 		Schema:           snowSchema,
// 		Warehouse:        snowWarehouse,
// 		NoProxy:          snowNoProxy,
// 		KeepAlive:        snowKeepAlive,
// 		QueueTimeout:     snowQueueTO,
// 		StatementTimeout: snowStmtTO,
// 	}
// 	config.SnowConfig = snowConfig
// }
