package betfair

type Filter struct {
	Wallet                  EWallet              `json:"wallet,omitempty"`
	Locale                  string               `json:"locale,omitempty"`
	FromRecord              int                  `json:"fromRecord,omitempty"`
	RecordCount             int                  `json:"recordCount,omitempty"`
	ItemTimeRange           *TimeRange           `json:"itemTimeRange,omitempty"`
	IncludeItem             EIncludeItem         `json:"recordCount,omitempty"`
	FromCurrency            string               `json:"fromCurrency,omitempty"`
	From                    EWallet              `json:"from,omitempty"`
	To                      EWallet              `json:"to,omitempty"`
	Amount                  float64              `json:"amount,omitempty"`
	BetIDs                  []string             `json:"betIds,omitempty"`
	MarketIDs               []string             `json:"marketIds,omitempty"`
	PriceProjection         *PriceProjection     `json:"priceProjection,omitempty"`
	OrderProjection         EOrderProjection     `json:"orderProjection,omitempty"`
	MarketProjection        *[]EMarketProjection `json:"marketProjection,omitempty"`
	TimeRange               *TimeRange           `json:"TimeRange,omitempty"`
	OrderBy                 EOrderBy             `json:"orderBy,omitempty"`
	SortDir                 ESortDir             `json:"sortDir,omitempty"`
	Sort                    EMarketSort          `json:"sort,omitempty"`
	MarketFilter            *MarketFilter        `json:"filter,omitempty"`
	BetStatus               EBetStatus           `json:"betStatus,omitempty"`
	EventTypeIDs            []string             `json:"eventTypeIds,omitempty"`
	EventIDs                []string             `json:"eventIds,omitempty"`
	RunnerIDs               []RunnerID           `json:"runnerIds,omitempty"`
	Side                    ESide                `json:"side,omitempty"`
	SettledTimeRange        *TimeRange           `json:"settledTimeRange,omitempty"`
	GroupBy                 EGroupBy             `json:"groupBy,omitempty"`
	IncludeItemDescription  bool                 `json:"includeItemDescription,omitempty"`
	MaxResults              int                  `json:"maxResults,omitempty"`
	IncludeSettledBets      bool                 `json:"includeSettledBets,omitempty"`
	TimeGranularity         ETimeGranularity     `json:"granularity,omitempty"`
	PlaceOrdersMarketID     string               `json:"marketId,omitempty"`
	PlaceOrdersInstructions []PlaceInstruction   `json:"instructions,omitempty"`
	CustomerOrderRefs       []string             `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs    []string             `json:"customerStrategyRefs,omitempty"`
}
