package cache

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/bombamong/geo-cache/pkg/snowflake"
)

type CacheLayer interface {
	FillRows(ctx context.Context) error
	IsReady() bool
	Query(cf CompareFunc) Merchants
}

func NewCacheLayer(src snowflake.Snowflake) CacheLayer {
	return (CacheLayer)(&cacheLayer{
		Source: src,
		Rows:   make([]RawData, 0),
		Ready:  false,
	})
}

type cacheLayer struct {
	Source snowflake.Snowflake
	Rows   []RawData
	Ready  bool
}

func (cl *cacheLayer) FillRows(ctx context.Context) error {
	start := time.Now()
	rows, err := cl.Source.Query(
		`select 
			transaction_time, 
			merchant, 
			value, 
			abid, 
			longitude, 
			latitude, 
			point 
		from test.fractional_events`)
	if err != nil {
		return err
	}

	for rows.Next() {
		var timeFrame sql.NullString
		var merchant sql.NullString
		var value sql.NullInt32
		var abid sql.NullString
		var lng sql.NullFloat64
		var lat sql.NullFloat64
		var point sql.NullString

		argList := []interface{}{&timeFrame, &merchant, &value, &abid, &lng, &lat, &point}
		err := rows.Scan(argList...)
		if err != nil {
			log.Printf("scan %v\n", err)
		}

		np := Point{}
		err = json.Unmarshal([]byte(point.String), &np)
		if err != nil {
			log.Print("error unmarshalling point", err)
		}

		nd := RawData{
			Timeframe: timeFrame.String,
			Merchant:  merchant.String,
			Value:     int(value.Int32),
			ABID:      abid.String,
			Longitude: lng.Float64,
			Latitude:  lat.Float64,
			Point:     np,
		}
		cl.Rows = append(cl.Rows, nd)
	}

	cl.Ready = true
	end := time.Since(start)
	log.Printf("time to fill rows %v\n", end)
	return nil
}

func (cl *cacheLayer) IsReady() bool {
	return cl.Ready
}

type CompareFunc func(rd RawData) bool

func (cl *cacheLayer) Query(cf CompareFunc) Merchants {
	counter := Counter{
		MerchCountMap: MerchCountMap{},
	}
	for _, event := range cl.Rows {
		if !cf(event) {
			continue
		}
		// merchant not enlisted
		if _, ok := counter.MerchCountMap[event.Merchant]; !ok {
			counter.MerchCountMap[event.Merchant] = NewMerchCount(event)
		} else {
			counter.MerchCountMap[event.Merchant].Add(*NewMerchCount(event))
		}

	}
	merchants := counter.ToMerchants()
	return merchants
}

// TEMP DATA STRUCTS
type RawData struct {
	Timeframe string
	Merchant  string
	Value     int
	ABID      string
	Longitude float64
	Latitude  float64
	Point     Point
}

type Point struct {
	Coordinates []float64 `json:"coordinates"`
	Type        string    `json:"type"`
}

type Merchants map[string]*MerchantData

type MerchantData struct {
	PurchaserCount   int      `json:"purchaser_count"`
	TransactionCount int      `json:"transaction_count"`
	TransactionSum   int      `json:"transaction_sum"`
	TimeData         TimeData `json:"time_data"`
	validK3          bool
}

func (md MerchantData) GetValidK3() bool {
	return md.validK3
}

type TimeData map[string]TimeTotals

type TimeTotals struct {
	PurchaserCount   int `json:"purchaser_count"`
	TransactionCount int `json:"transaction_count"`
	TransactionSum   int `json:"transaction_sum"`
}
