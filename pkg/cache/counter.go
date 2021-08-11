package cache

type TimeCountMap map[string]*TimeCount

func (oldD TimeCountMap) Add(newD TimeCountMap) {
	for k, v := range newD {
		if _, ok := oldD[k]; ok {
			oldD[k].Add(*v)
		} else {
			oldD[k] = v
		}
	}
}

type TimeCount struct {
	ValueSum         int
	EventCount       int
	DistinctCountMap DistinctCountMap
}

func (oldD *TimeCount) Add(newD TimeCount) {
	oldD.ValueSum += newD.ValueSum
	oldD.EventCount += newD.EventCount
	oldD.DistinctCountMap.Add(newD.DistinctCountMap)
}

type DistinctCountMap map[string]int

func (oldD DistinctCountMap) Add(newD DistinctCountMap) DistinctCountMap {
	for k, v := range newD {
		oldD[k] += v
	}
	return oldD
}

type MerchCount struct {
	ValueSum         int
	EventCount       int
	DistinctCountMap DistinctCountMap
	TimeCountMap     TimeCountMap
}

func (oldD *MerchCount) Add(newD MerchCount) {
	oldD.ValueSum += newD.ValueSum
	oldD.EventCount += newD.EventCount
	oldD.DistinctCountMap.Add(newD.DistinctCountMap)
	oldD.TimeCountMap.Add(newD.TimeCountMap)
}

func NewMerchCount(event RawData) *MerchCount {
	return &MerchCount{
		ValueSum:   event.Value,
		EventCount: 1,
		DistinctCountMap: DistinctCountMap{
			event.ABID: 1,
		},
		TimeCountMap: TimeCountMap{
			event.Timeframe: &TimeCount{
				ValueSum:   event.Value,
				EventCount: 1,
				DistinctCountMap: DistinctCountMap{
					event.ABID: 1,
				},
			},
		},
	}
}

type MerchCountMap map[string]*MerchCount

func (oldD MerchCountMap) Add(newD MerchCountMap) {
	for k, v := range newD {
		if _, ok := oldD[k]; ok {
			oldD[k].Add(*v)
		} else {
			oldD[k] = v
		}
	}
}

type Counter struct {
	MerchCountMap MerchCountMap
}

func (c Counter) ToMerchants() Merchants {
	merch := make(Merchants)
	for merchName, mdata := range c.MerchCountMap {
		md := &MerchantData{
			PurchaserCount:   len(mdata.DistinctCountMap),
			TransactionCount: mdata.EventCount,
			TransactionSum:   mdata.ValueSum,
			TimeData:         TimeData{},
			validK3:          true,
		}
		for timeFrame, timeData := range mdata.TimeCountMap {
			md.TimeData[timeFrame] = TimeTotals{
				PurchaserCount:   len(timeData.DistinctCountMap),
				TransactionCount: timeData.EventCount,
				TransactionSum:   timeData.ValueSum,
			}
		}
		merch[merchName] = md
	}
	return merch
}
