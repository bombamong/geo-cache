package snowflake

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/snowflakedb/gosnowflake"
)

type Snowflake struct {
	db     *sql.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSnowflake(snowConfig Config) Snowflake {
	ctx, cancel := context.WithCancel(context.Background())
	db, err := sql.Open("snowflake", snowConfig.stringify())
	if err != nil {
		log.Fatalln("Error opening db")
	}
	if db == nil {
		log.Fatalln("DB does not exist")
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Printf("Error establishing snowflake connection: %s\n", err.Error())
	}
	return Snowflake{
		db:     conn,
		ctx:    ctx,
		cancel: cancel,
	}
}

type IterableRows interface {
	Next() bool
	Scan(dest ...interface{}) error
}

func (s *Snowflake) Query(query string) (*sql.Rows, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// func (s *Snowflake) Query(query string) (response.Merchants, error) {
// 	merchants := response.Merchants{}
// 	ctx, cancel := context.WithTimeout(s.ctx, time.Second*6)

// 	rows, err := s.db.QueryContext(ctx, query)
// 	cancel()
// 	if err != nil {
// 		log.Printf("Error while querying: %s\n", err.Error())
// 		return nil, err
// 	}

// 	select {
// 	case <-ctx.Done():
// 		cancel()
// 		return nil, errors.New("timeout error")
// 	default:
// 	}

// 	for rows.Next() {
// 		merchantData := new(response.MerchantData)
// 		timeMap := make(response.TimeData)
// 		timeMap.Reset()

// 		var transactionCount sql.NullInt64
// 		var purchaserCount sql.NullInt64
// 		var transactionAmount sql.NullInt64
// 		var merchantName sql.NullString
// 		var transactionTime sql.NullString

// 		err := rows.Scan(&transactionCount, &purchaserCount, &transactionAmount, &merchantName, &transactionTime)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		if transactionCount.Valid {
// 			merchantData.TransactionCount = int(transactionCount.Int64)
// 		}
// 		if purchaserCount.Valid {
// 			merchantData.PurchaserCount = int(purchaserCount.Int64)
// 		}
// 		if transactionAmount.Valid {
// 			merchantData.TransactionSum = int(transactionAmount.Int64)
// 		}
// 		if transactionTime.Valid {
// 			timeMap[transactionTime.String] = response.TimeTotals{
// 				PurchaserCount:   merchantData.PurchaserCount,
// 				TransactionCount: merchantData.TransactionCount,
// 				TransactionSum:   merchantData.TransactionSum,
// 			}
// 			merchantData.TimeData = timeMap
// 		}

// 		oldMerch, ok := merchants[merchantName.String]
// 		if ok {
// 			oldMerch.Add(*merchantData)
// 			merchants[merchantName.String] = oldMerch
// 		} else {
// 			merchants[merchantName.String] = merchantData
// 		}
// 	}
// 	for _, merch := range merchants {
// 		merch.ApplyK3()
// 	}

// 	return merchants, nil
// }

func (s *Snowflake) Close() {
	if err := s.db.Close(); err != nil {
		log.Fatalf("Error closing snowflake: %s\n", err.Error())
	}
}

func (s *Snowflake) Cancel() {
	s.cancel()
}
