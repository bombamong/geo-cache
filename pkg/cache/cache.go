package cache

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Cache struct {
	Data []Row
}

type MerchantData struct {
	Purchasers     map[string]int
	Transactions   int
	TransactionSum int
}

func (c Cache) Query(merchant string) MerchantData {
	start := time.Now()

	md := MerchantData{
		Purchasers: make(map[string]int),
	}
	for _, v := range c.Data {
		if v.Merchant == merchant {
			if _, ok := md.Purchasers[v.ABID]; !ok {
				md.Purchasers[v.ABID] = 1
			} else {
				md.Purchasers[v.ABID] += 1
			}
			md.Transactions += 1
			md.TransactionSum += v.Value
		}
	}

	// by time
	//purchaser_count
	//transaction_count
	//transaction_sum
	// then total up

	elapsed := time.Since(start)
	log.Printf("Query time for %s: %v\n", merchant, elapsed)

	return md
}

type Row struct {
	Timeframe string
	Merchant  string
	Value     int
	ABID      string
	Longitude float64
	Latitude  float64
}

func (r *Row) ParseLine(line string) {
	line = strings.Trim(line, "\n")
	arr := strings.Split(line, ",")

	if len(arr) < 5 {
		return
	}
	r.Timeframe = arr[0]
	r.Merchant = arr[1]
	i, err := strconv.ParseInt(arr[2], 10, 64)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	r.Value = int(i)
	r.ABID = arr[3]
	f, err := strconv.ParseFloat(arr[4], 64)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	r.Longitude = f
	f, err = strconv.ParseFloat(arr[5], 64)
	if err != nil {
		log.Printf("error %v\n", err)
	}
	r.Latitude = f
}
