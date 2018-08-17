package go_iexclient

type TOPS struct {
	// Refers to the stock ticker.
	Symbol        string
	// Refers to IEX’s percentage of the market in the stock.
	MarketPercent float64
	// Refers to amount of shares on the bid on IEX.
	BidSize       int
	// Refers to the best bid price on IEX.
	BidPrice      float64
	// Refers to amount of shares on the ask on IEX.
	AskSize       int
	// Refers to the best ask price on IEX.
	AskPrice      float64
	// Refers to shares traded in the stock on IEX.
	Volume        int
	// Refers to last sale price of the stock on IEX. (Refer to the attribution section above.)
	LastSalePrice float64
	// Refers to last sale size of the stock on IEX.
	LastSaleSize  int
	// Refers to last sale time of the stock on IEX.
	LastSaleTime  Time
	// Refers to the last update time of the data.
	// If the value is the zero Time, IEX has not quoted the symbol in
	// the trading day.
	LastUpdated   Time
}

//type Batch struct {
//	Symbol           string
//	CompanyName      string
//	PrimaryExchange  string
//	Sector           string
//	CalculationPrice string
//	Open             float64
//	OpenTime         Time
//	Close            float64
//	CloseTime        Time
//	High             float64
//	Low              float64
//	LatestPrice      float64
//	LatestSource     string
//	LatestTime       string
//	LatestUpdate     Time
//	LatestVolume     uint64
//	IexRealtimePrice float64
//	IexRealtimeSize  uint64
//	IexLastUpdated   string
//	DelayedPrice     float64
//	DelayedPriceTime Time
//	PreviousClose    float64
//	Change           float64
//	ChangePercent    float64
//	AvgTotalVolume   uint64
//	MarketCap        uint64
//	PeRatio          float64
//	Week52High       float64
//	Week52Low        float64
//	YtdChange        float64
//}

type Batch struct {
	Symbol      string
	CompanyName string
	Exchange    string
	Industry    string
	Website     string
	Description string
	CEO         string
	IssueType   string
	Sector      string
	LatestPrice float64
}