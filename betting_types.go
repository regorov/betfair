package betfair

type MarketFilter struct {
	TextQuery          *string               `json:"textQuery,omitempty"`
	ExchangeIDs        []string              `json:"exchangeIds,omitempty"`
	EventTypeIDs       []string              `json:"eventTypeIds,omitempty"`
	EventIDs           []string              `json:"eventIds,omitempty"`
	CompetitionIDs     []string              `json:"competitionIds,omitempty"`
	MarketIDs          []string              `json:"marketIds,omitempty"`
	Venues             []string              `json:"venues,omitempty"`
	BspOnly            *bool                 `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  *bool                 `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         *bool                 `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []EMarketBettingTypes `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string              `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string              `json:"marketTypeCodes,omitempty"`
	MarketStartTime    *TimeRange            `json:"marketStartTime,omitempty"`
	WithOrders         []EOrderStatus        `json:"withOrders,omitempty"`
	RaceTypes          []string              `json:"raceTypes,omitempty"`
}

type MarketCatalogue struct {
	MarketID        string             `json:"marketId"`
	MarketName      string             `json:"marketName"`
	MarketStartTime *Time              `json:"marketStartTime,omitempty"`
	Description     *MarketDescription `json:"description,omitempty"`
	TotalMatched    *float64           `json:"totalMatched,omitempty"`
	Runners         []RunnerCatalog    `json:"runners,omitempty"`
	EventType       *EventType         `json:"eventType,omitempty"`
	Competition     *Competition       `json:"competition,omitempty"`
	Event           *Event             `json:"event,omitempty"`
}

type MarketBook struct {
	MarketID              string              `json:"marketId"`
	IsMarketDataDelayed   bool                `json:"isMarketDataDelayed"`
	Status                *EMarketStatus      `json:"status,omitempty"`
	BetDelay              *int                `json:"betDelay,omitempty"`
	BspReconciled         *bool               `json:"bspReconciled,omitempty"`
	Complete              *bool               `json:"complete,omitempty"`
	Inplay                *bool               `json:"inplay,omitempty"`
	NumberOfWinners       *int                `json:"numberOfWinners,omitempty"`
	NumberOfRunners       *int                `json:"numberOfRunners,omitempty"`
	NumberOfActiveRunners *int                `json:"numberOfActiveRunners,omitempty"`
	LastMatchTime         *Time               `json:"lastMatchTime,omitempty"`
	TotalMatched          *float64            `json:"totalMatched,omitempty"`
	TotalAvailable        *float64            `json:"totalAvailable,omitempty"`
	CrossMatching         *bool               `json:"crossMatching,omitempty"`
	RunnersVoidable       *bool               `json:"runnersVoidable,omitempty"`
	Version               *int64              `json:"version,omitempty"`
	Runners               []Runner            `json:"runners,omitempty"`
	KeyLineDescription    *KeyLineDescription `json:"keyLineDescription,omitempty"`
}

type RunnerCatalog struct {
	SelectionID  int64             `json:"selectionId"`
	RunnerName   string            `json:"runnerName"`
	Handicap     float64           `json:"handicap"`
	SortPriority int               `json:"sortPriority"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type Runner struct {
	SelectionID       int64              `json:"selectionId"`
	Handicap          float64            `json:"handicap"`
	Status            ERunnerStatus      `json:"status"`
	AdjustmentFactor  float64            `json:"adjustmentFactor"`
	LastPriceTraded   *float64           `json:"lastPriceTraded,omitempty"`
	TotalMatched      *float64           `json:"totalMatched,omitempty"`
	RemovalDate       *Time              `json:"removalDate,omitempty"`
	SP                *StartingPrices    `json:"sp,omitempty"`
	EX                *ExchangePrices    `json:"ex,omitempty"`
	Orders            []Order            `json:"orders,omitempty"`
	Matches           []Match            `json:"matches,omitempty"`
	MatchesByStrategy map[string][]Match `json:"matchesByStrategy,omitempty"`
}

type StartingPrices struct {
	NearPrice         *float64    `json:"nearPrice,omitempty"`
	FarPrice          *float64    `json:"farPrice,omitempty"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          *float64    `json:"actualSP,omitempty"`
}

type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

type Event struct {
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Timezone    *string `json:"timezone,omitempty"`
	Venue       *string `json:"venue,omitempty"`
	OpenDate    *Time   `json:"openDate,omitempty"`
}

type EventResult struct {
	Event       *Event `json:"event,omitempty"`
	MarketCount *int   `json:"marketCount,omitempty"`
}

type Competition struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CompetitionResult struct {
	Competition       *Competition `json:"competition,omitempty"`
	MarketCount       *int         `json:"marketCount,omitempty"`
	CompetitionRegion *string      `json:"competitionRegion,omitempty"`
}

type EventType struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type EventTypeResult struct {
	EventType   *EventType `json:"eventType,omitempty"`
	MarketCount *int       `json:"marketCount,omitempty"`
}

type MarketTypeResult struct {
	MarketType  *string `json:"marketType,omitempty"`
	MarketCount *int    `json:"marketCount,omitempty"`
}

type CountryCodeResult struct {
	CountryCode *string `json:"countryCode,omitempty"`
	MarketCount *int    `json:"marketCount,omitempty"`
}

type VenueResult struct {
	Venue       *string `json:"venue,omitempty"`
	MarketCount *int    `json:"marketCount,omitempty"`
}

type TimeRange struct {
	From *Time `json:"from,omitempty"`
	To   *Time `json:"to,omitempty"`
}

type TimeRangeResult struct {
	TimeRange   *TimeRange `json:"timeRange,omitempty"`
	MarketCount *int       `json:"marketCount,omitempty"`
}

type Order struct {
	BetID               string           `json:"betId"`
	OrderType           EOrderType       `json:"orderType"`
	Status              EOrderStatus     `json:"orderStatus"`
	PersistenceType     EPersistenceType `json:"persistenceType"`
	Side                ESide            `json:"side"`
	Price               float64          `json:"price"`
	Size                float64          `json:"size"`
	BspLiability        float64          `json:"bspLiability"`
	PlacedDate          Time             `json:"placedDate"`
	AvgPriceMatched     *float64         `json:"avgPriceMatched,omitempty"`
	SizeMatched         *float64         `json:"sizeMatched,omitempty"`
	SizeRemaining       *float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed          *float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled       *float64         `json:"sizeCancelled,omitempty"`
	SizeVoided          *float64         `json:"sizeVoided,omitempty"`
	CustomerOrderRef    *string          `json:"customerOrderRef,omitempty"`
	CustomerStrategyRef *string          `json:"customerStrategyRef,omitempty"`
}

// An individual bet Match, or rollup by price or avg price. Rollup depends on the requested MatchProjection.
type Match struct {
	BetID     *string `json:"betId,omitempty"`
	MatchID   *string `json:"matchId,omitempty"`
	Side      ESide   `json:"side"`
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	MatchDate *Time   `json:"matchDate,omitempty"`
}

type MarketVersion struct {
	Version *int64 `json:"version,omitempty"`
}

type MarketDescription struct {
	PersistenceEnabled     bool                    `json:"persistenceEnabled"`
	BspMarket              bool                    `json:"bspMarket"`
	MarketTime             Time                    `json:"marketTime"`
	SuspendTime            Time                    `json:"suspendTime"`
	SettleTime             *Time                   `json:"settleTime,omitempty"`
	BettingType            EMarketBettingType      `json:"bettingType"`
	TurnInPlayEnabled      bool                    `json:"turnInPlayEnabled"`
	MarketType             string                  `json:"marketType"`
	Regulator              string                  `json:"regulator"`
	MarketBaseRate         float64                 `json:"marketBaseRate"`
	DiscountAllowed        bool                    `json:"discountAllowed"`
	Wallet                 *string                 `json:"wallet,omitempty"`
	Rules                  *string                 `json:"rules,omitempty"`
	RulesHasDate           *bool                   `json:"rulesHasDate,omitempty"`
	EachWayDivisor         *float64                `json:"eachWayDivisor,omitempty"`
	Clarifications         *string                 `json:"clarifications,omitempty"`
	LineRangeInfo          *MarketLineRangeInfo    `json:"lineRangeInfo,omitempty"`
	RaceType               *string                 `json:"raceType,omitempty"`
	PriceLadderDescription *PriceLadderDescription `json:"priceLadderDescription,omitempty"`
}

type MarketRates struct {
	MarketBaseRate  float64 `json:"marketBaseRate"`
	DiscountAllowed bool    `json:"discountAllowed"`
}

type MarketLicence struct {
	Wallet         string  `json:"wallet"`
	Rules          *string `json:"rules,omitempty"`
	RulesHasDate   *bool   `json:"rulesHasDate,omitempty"`
	Clarifications *string `json:"clarifications,omitempty"`
}

type MarketLineRangeInfo struct {
	MaxUnitValue float64 `json:"maxUnitValue"`
	MinUnitValue float64 `json:"minUnitValue"`
	Interval     float64 `json:"interval"`
	MarketUnit   string  `json:"marketUnit"`
}

type PriceSize struct {
	Price float64 `json:"price"`
	Size  float64 `json"size"`
}

type ClearedOrderSummary struct {
	EventTypeID         *string           `json:"eventTypeId,omitempty"`
	EventID             *string           `json:"eventId,omitempty"`
	MarketID            *string           `json:"marketId,omitempty"`
	SelectionID         *int64            `json:"selectionId,omitempty"`
	Handicap            *float64          `json:"handicap,omitempty"`
	BetID               *string           `json:"betId,omitempty"`
	PlacedDate          *Time             `json:"placedDate,omitempty"`
	PersistenceType     *EPersistenceType `json:"persistenceType,omitempty"`
	OrderType           *EOrderType       `json:"orderType,omitempty"`
	Side                *ESide            `json:"side,omitempty"`
	ItemDescription     *ItemDescription  `json:"itemDescription,omitempty"`
	BetOutcome          *string           `json:"betOutcome,omitempty"`
	PriceRequested      *float64          `json:"priceRequested,omitempty"`
	SettledDate         *Time             `json:"settledDate,omitempty"`
	LastMatchedDate     *Time             `json:"lastMatchedDate,omitempty"`
	BetCount            *int              `json:"betCount,omitempty"`
	Commission          *float64          `json:"commission,omitempty"`
	PriceMatched        *float64          `json:"priceMatched,omitempty"`
	PriceReduced        *bool             `json:"priceReduced,omitempty"`
	SizeSettled         *float64          `json:"sizeSettled,omitempty"`
	Profit              *float64          `json:"profit,omitempty"`
	SizeCancelled       *float64          `json:"sizeCancelled,omitempty"`
	CustomerOrderRef    *string           `json:"customerOrderRef,omitempty"`
	CustomerStrategyRef *string           `json:"customerStrategyRef,omitempty"`
}

type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

type ItemDescription struct {
	EventTypeDesc   *string  `json:"eventTypeDesc,omitempty"`
	EventDesc       *string  `json:"eventDesc,omitempty"`
	MarketDesc      *string  `json:"marketDesc,omitempty"`
	MarketType      *string  `json:"marketType,omitempty"`
	MarketStartTime *Time    `json:"marketStartTime,omitempty"`
	RunnerDesc      *string  `json:"runnerDesc,omitempty"`
	NumberOfWinners *int     `json:"numberOfWinners,omitempty"`
	EachWayDivisor  *float64 `json:"eachWayDivisor,omitempty"`
}

type RunnerID struct {
	MarketID    *string  `json:"marketId,omitempty"`
	SelectionID *int64   `json:"selectionId,omitempty"`
	Handicap    *float64 `json:"handicap,omitempty"`
}

type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

type CurrentOrderSummary struct {
	BetID               string           `json:"betId"`
	MarketID            string           `json:"marketId"`
	SelectionID         int64            `json:"selectionId"`
	Handicap            float64          `json:"handicap"`
	PriceSize           PriceSize        `json:"priceSize"`
	BspLiability        float64          `json:"bspLiability"`
	Side                ESide            `json:"side"`
	Status              EOrderStatus     `json:"status"`
	PersistenceType     EPersistenceType `json:"persistenceType"`
	OrderType           EOrderType       `json:"orderType"`
	PlacedDate          Time             `json:"placedDate"`
	MatchedDate         Time             `json:"matchedDate"`
	AveragePriceMatched *float64         `json:"averagePriceMatched,omitempty"`
	SizeMatched         *float64         `json:"sizeMatched,omitempty"`
	SizeRemaining       *float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed          *float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled       *float64         `json:"sizeCancelled,omitempty"`
	SizeVoided          *float64         `json:"sizeVoided,omitempty"`
	RegulatorAuthCode   *string          `json:"regulatorAuthCode,omitempty"`
	RegulatorCode       *string          `json:"regulatorCode,omitempty"`
	CustomerOrderRef    *string          `json:"customerOrderRef,omitempty"`
	CustomerStrategyRef *string          `json:"customerStrategyRef,omitempty"`
}

type PlaceInstruction struct {
	OrderType          EOrderType          `json:"orderType"`
	SelectionID        int64               `json:"selectionId"`
	Handicap           *float64            `json:"handicap,omitempty"`
	Side               ESide               `json:"side"`
	LimitOrder         *LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   *string             `json:"customerOrderRef,omitempty"`
}

type PlaceExecutionReport struct {
	CustomerRef        *string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus     `json:"status"`
	ErrorCode          *EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           *string                    `json:"marketId,omitempty"`
	InstructionReports []PlaceInstructionReport   `json:"instructionReports,omitempty"`
}

type LimitOrder struct {
	Size            float64          `json:"size"`
	Price           float64          `json:"price"`
	PersistenceType EPersistenceType `json:"persistenceType"`
	TimeInForce     *ETimeInForce    `json:"timeInForce,omitempty"`
	MinFillSize     *float64         `json:"minFillSize,omitempty"`
	BetTargetType   *EBetTargetType  `json:"betTargetType,omitempty"`
	BetTargetSize   *float64         `json:"betTargetSize,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability float64 `json:"liability"`
	Price     float64 `json:"price"`
}

type MarketOnCloseOrder struct {
	Liability float64 `json:"liability"`
}

type PlaceInstructionReport struct {
	Status              EInstructionReportStatus     `json:"status"`
	ErrorCode           *EInstructionReportErrorCode `json:"errorCode,omitempty"`
	OrderStatus         *EOrderStatus                `json:"orderStatus,omitempty"`
	Instruction         PlaceInstruction             `json:"instruction"`
	BetID               *string                      `json:"betId,omitempty"`
	PlacedDate          *Time                        `json:"placedDate,omitempty"`
	AveragePriceMatched *float64                     `json:"averagePriceMatched,omitempty"`
	SizeMatched         *float64                     `json:"sizeMatched,omitempty"`
}

type CancelInstruction struct {
	BetId         string   `json:"betId"`
	SizeReduction *float64 `json:"sizeReduction,omitempty"`
}

type CancelExecutionReport struct {
	CustomerRef        *string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus     `json:"status"`
	ErrorCode          *EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           *string                    `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport  `json:"instructionReports,omitempty"`
}

type ReplaceInstruction struct {
	BetId    string  `json:"betId"`
	NewPrice float64 `json:"newPrice"`
}

type ReplaceExecutionReport struct {
	CustomerRef        *string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus     `json:"status"`
	ErrorCode          *EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           *string                    `json:"marketId,omitempty"`
	InstructionReports []ReplaceInstructionReport `json:"instructionReports,omitempty"`
}

type ReplaceInstructionReport struct {
	Status                  EInstructionReportStatus     `json:"status"`
	ErrorCode               *EInstructionReportErrorCode `json:"errorCode,omitempty"`
	CancelInstructionReport *CancelInstructionReport     `json:"cancelInstructionReport,omitempty"`
	PlaceInstructionReport  *PlaceInstructionReport      `json:"placeInstructionReport,omitempty"`
}

type CancelInstructionReport struct {
	Status        EInstructionReportStatus     `json:"status"`
	ErrorCode     *EInstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction   *CancelInstruction           `json:"instruction,omitempty"`
	SizeCancelled float64                      `json:"sizeCancelled"`
	CancelledDate *Time                        `json:"cancelledDate,omitempty"`
}

type UpdateInstruction struct {
	BetId              string           `json:"betId"`
	NewPersistenceType EPersistenceType `json:"newPersistenceType"`
}

type UpdateExecutionReport struct {
	CustomerRef        *string                    `json:"customerRef,omitempty"`
	Status             EExecutionReportStatus     `json:"status"`
	ErrorCode          *EExecutionReportErrorCode `json:"errorCode,omitempty"`
	MarketID           *string                    `json:"marketId,omitempty"`
	InstructionReports []UpdateInstructionReport  `json:"instructionReports,omitempty"`
}

type UpdateInstructionReport struct {
	Status      EInstructionReportStatus     `json:"status"`
	ErrorCode   *EInstructionReportErrorCode `json:"errorCode,omitempty"`
	Instruction UpdateInstruction            `json:"instruction"`
}

type PriceProjection struct {
	PriceData             []EPriceData           `json:"priceData,omitempty"`
	EXBestOffersOverrides *EXBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            *bool                  `json:"virtualise,omitempty"`
	RolloverStakes        *bool                  `json:"rolloverStakes,omitempty"`
}

type EXBestOffersOverrides struct {
	BestPricesDepth          *int          `json:"bestPricesDepth,omitempty"`
	RollupModel              *ERollupModel `json:"rollupModel,omitempty"`
	RollupLimit              *int          `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold *float64      `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    *int          `json:"rollupLiabilityFactor,omitempty"`
}

type MarketProfitAndLoss struct {
	MarketID          *string               `json:"marketId,omitempty"`
	CommissionApplied *float64              `json:"commissionApplied,omitempty"`
	ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses,omitempty"`
}

type RunnerProfitAndLoss struct {
	SelectionID *int64   `json:"selectionId,omitempty"`
	IfWin       *float64 `json:"ifWin,omitempty"`
	IfLose      *float64 `json:"ifLose,omitempty"`
	IfPlace     *float64 `json:"ifPlace,omitempty"`
}

type PriceLadderDescription struct {
	Type *EPriceLadderType `json:"type,omitempty"`
}

type KeyLineSelection struct {
	SelectionID *int64   `json:"selectionId,omitempty"`
	Handicap    *float64 `json:"handicap,omitempty"`
}

type KeyLineDescription struct {
	KeyLine []KeyLineSelection `json:"keyLine,omitempty"`
}
