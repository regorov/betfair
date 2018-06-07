package betfair

type EMarketProjection string

const (
	MP_COMPETITION        EMarketProjection = "COMPETITION"
	MP_EVENT              EMarketProjection = "EVENT"
	MP_EVENT_TYPE         EMarketProjection = "EVENT_TYPE"
	MP_MARKET_START_TIME  EMarketProjection = "MARKET_START_TIME"
	MP_MARKET_DESCRIPTION EMarketProjection = "MARKET_DESCRIPTION"
	MP_RUNNER_DESCRIPTION EMarketProjection = "RUNNER_DESCRIPTION"
	MP_RUNNER_METADATA    EMarketProjection = "RUNNER_METADATA"

//	ErrUnknownMarketProjection EMarketProjection = "Unknown marketProjection value"
)

type EPriceData string

const (
	PD_SP_AVAILABLE   EPriceData = "SP_AVAILABLE"
	PD_SP_TRADED      EPriceData = "SP_TRADED"
	PD_EX_BEST_OFFERS EPriceData = "EX_BEST_OFFERS"
	PD_EX_ALL_OFFERS  EPriceData = "EX_ALL_OFFERS"
	PD_EX_TRADED      EPriceData = "EX_TRADED"
	//ErrUnknownPriceData EPriceData = "Unknown priceData value"
)

type EMatchProjection string

const (
	MP_NO_ROLLUP              EMatchProjection = "NO_ROLLUP"
	MP_ROLLED_UP_BY_PRICE     EMatchProjection = "ROLLED_UP_BY_PRICE"
	MP_ROLLED_UP_BY_AVG_PRICE EMatchProjection = "ROLLED_UP_BY_AVG_PRICE"

//	ErrUnknownMatchProjection EMatchProjection = "Unknown matchProjection value"
)

type EOrderProjection string

const (
	OP_ALL                EOrderProjection = "ALL"
	OP_EXECUTABLE         EOrderProjection = "EXECUTABLE"
	OP_EXECUTION_COMPLETE EOrderProjection = "EXECUTION_COMPLETE"

//	ErrUnknownOrderProjection EOrderProjection = "Unknown orderProjection value"
)

type EMarketStatus string

const (
	MS_INACTIVE  EMarketStatus = "INACTIVE"
	MS_OPEN      EMarketStatus = "OPEN"
	MS_SUSPENDED EMarketStatus = "SUSPENDED"
	MS_CLOSED    EMarketStatus = "CLOSED"

//	ErrUnknowMarketStatus EMarketStatus = "Unknown marketStatus value"
)

type ERunnerStatus string

const (
	RS_ACTIVE         ERunnerStatus = "ACTIVE"
	RS_WINNER         ERunnerStatus = "WINNER"
	RS_LOSER          ERunnerStatus = "LOSER"
	RS_PLACED         ERunnerStatus = "PLACED"
	RS_REMOVED_VACANT ERunnerStatus = "REMOVED_VACANT"
	RS_REMOVED        ERunnerStatus = "REMOVED"
	RS_HIDDEN         ERunnerStatus = "HIDDEN"

//	ErrUnknownRunnerStatus ERunnerStatus = "Unknown runnerStatus value"
)

type ETimeGranularity string

const (
	TG_DAYS    ETimeGranularity = "DAYS"
	TG_HOURS   ETimeGranularity = "HOURS"
	TG_MINUTES ETimeGranularity = "MINUTES"

//	ErrUnknownTimeGranularity ETimeGranularity = "Unknown timeGranularity value"
)

type ESide string

const (
	S_BACK ESide = "BACK"
	S_LAY  ESide = "LAY"
	//ErrUnknownSide ESide = "Unknown side value. Supported BACK or LAY only"
)

type EOrderStatus string

const (
	OS_PENDING            EOrderStatus = "PENDING"
	OS_EXECUTION_COMPLETE EOrderStatus = "EXECUTION_COMPLETE"
	OS_EXECUTABLE         EOrderStatus = "EXECUTABLE"
	OS_EXPIRED            EOrderStatus = "EXPIRED"
	//ErrOrderStatus        EOrderStatus = "Unknown orderStatus value"
)

type EOrderBy string

const (
	OB_BY_BET          EOrderBy = "BY_BET"
	OB_BY_MARKET       EOrderBy = "BY_MARKET"
	OB_BY_MATCH_TIME   EOrderBy = "BY_MATCH_TIME"
	OB_BY_PLACE_TIME   EOrderBy = "BY_PLACE_TIME"
	OB_BY_SETTLED_TIME EOrderBy = "BY_SETTLED_TIME"
	OB_BY_VOID_TIME    EOrderBy = "BY_VOID_TIME"

//	ErrUnknownOrderBy  EOrderBy = "Unknown orderBy value"
)

type ESortDir string

const (
	SD_EARLIEST_TO_LATEST ESortDir = "EARLIEST_TO_LATEST"
	SD_LATEST_TO_EARLIEST ESortDir = "LATEST_TO_EARLIEST"
)

type EOrderType string

const (
	OT_LIMIT           EOrderType = "LIMIT"
	OT_LIMIT_ON_CLOSE  EOrderType = "LIMIT_ON_CLOSE"
	OT_MARKET_ON_CLOSE EOrderType = "MARKET_ON_CLOSE"
)

type EMarketSort string

const (
	MS_MINIMUM_TRADED    EMarketSort = "MINIMUM_TRADED"
	MS_MAXIMUM_TRADED    EMarketSort = "MAXIMUM_TRADED"
	MS_MINIMUM_AVAILABLE EMarketSort = "MINIMUM_AVAILABLE"
	MS_MAXIMUM_AVAILABLE EMarketSort = "MAXIMUM_AVAILABLE"
	MS_FIRST_TO_START    EMarketSort = "FIRST_TO_START"
	MS_LAST_TO_START     EMarketSort = "LAST_TO_START"

//	ErrUnknownMarketSort EMarketSort = "Unknown marketSort value"
)

type EMarketBettingType string
type EMarketBettingTypes []EMarketBettingType

const (
	MBT_ODDS                       EMarketBettingType = "ODDS"
	MBT_LINE                       EMarketBettingType = "LINE"
	MBT_ASIAN_HANDICAP_DOUBLE_LINE EMarketBettingType = "ASIAN_HANDICAP_DOUBLE_LINE"
	MBT_ASIAN_HANDICAP_SINGLE_LINE EMarketBettingType = "ASIAN_HANDICAP_SINGLE_LINE"
	MBT_FIXED_ODDS                 EMarketBettingType = "FIXED_ODDS"
	//ErrUnknownMarketBettingType    EMarketBettingType = "Unknown marketBettingType value"
)

type EExecutionReportStatus string

const (
	ERS_SUCCESS               EExecutionReportStatus = "SUCCESS"
	ERS_FAILURE               EExecutionReportStatus = "FAILURE"
	ERS_PROCESSED_WITH_ERRORS EExecutionReportStatus = "PROCESSED_WITH_ERRORS"
	ERS_TIMEOUT               EExecutionReportStatus = "TIMEOUT"
)

type EExecutionReportErrorCode string

const (
	EREC_ERROR_IN_MATCHER            EExecutionReportErrorCode = "ERROR_IN_MATCHER"
	EREC_PROCESSED_WITH_ERRORS       EExecutionReportErrorCode = "PROCESSED_WITH_ERRORS"
	EREC_BET_ACTION_ERROR            EExecutionReportErrorCode = "BET_ACTION_ERROR"
	EREC_INVALID_ACCOUNT_STATE       EExecutionReportErrorCode = "INVALID_ACCOUNT_STATE"
	EREC_INVALID_WALLET_STATUS       EExecutionReportErrorCode = "INVALID_WALLET_STATUS"
	EREC_INSUFFICIENT_FUNDS          EExecutionReportErrorCode = "INSUFFICIENT_FUNDS"
	EREC_LOSS_LIMIT_EXCEEDED         EExecutionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	EREC_MARKET_SUSPENDED            EExecutionReportErrorCode = "MARKET_SUSPENDED"
	EREC_MARKET_NOT_OPEN_FOR_BETTING EExecutionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	EREC_DUPLICATE_TRANSACTION       EExecutionReportErrorCode = "DUPLICATE_TRANSACTION"
	EREC_INVALID_ORDER               EExecutionReportErrorCode = "INVALID_ORDER"
	EREC_INVALID_MARKET_ID           EExecutionReportErrorCode = "INVALID_MARKET_ID"
	EREC_PERMISSION_DENIED           EExecutionReportErrorCode = "PERMISSION_DENIED"
	EREC_DUPLICATE_BETIDS            EExecutionReportErrorCode = "DUPLICATE_BETIDS"
	EREC_NO_ACTION_REQUIRED          EExecutionReportErrorCode = "NO_ACTION_REQUIRED"
	EREC_SERVICE_UNAVAILABLE         EExecutionReportErrorCode = "SERVICE_UNAVAILABLE"
	EREC_REJECTED_BY_REGULATOR       EExecutionReportErrorCode = "REJECTED_BY_REGULATOR"
	EREC_NO_CHASING                  EExecutionReportErrorCode = "NO_CHASING"
	EREC_REGULATOR_IS_NOT_AVAILABLE  EExecutionReportErrorCode = "REGULATOR_IS_NOT_AVAILABLE"
	EREC_TOO_MANY_INSTRUCTIONS       EExecutionReportErrorCode = "TOO_MANY_INSTRUCTIONS"
	EREC_INVALID_MARKET_VERSION      EExecutionReportErrorCode = "INVALID_MARKET_VERSION"
	//ErrUnknownExecutionReportErrorCode EExecutionReportErrorCode = "Unknown executionReportErrorCode value"
)

type EPersistenceType string

const (
	PT_LAPSE           EPersistenceType = "LAPSE"
	PT_PERSIST         EPersistenceType = "PERSIST"
	PT_MARKET_ON_CLOSE EPersistenceType = "MARKET_ON_CLOSE"
	//ErrUnknownPersistenceType EPersistenceType = "Unknown persistenceType value"
)

type EInstructionReportStatus string

const (
	IRS_SUCCESS EInstructionReportStatus = "SUCCESS"
	IRS_FAILURE EInstructionReportStatus = "FAILURE"
	IRS_TIMEOUT EInstructionReportStatus = "TIMEOUT"
)

type EInstructionReportErrorCode string

const (
	IREC_INVALID_BET_SIZE                EInstructionReportErrorCode = "INVALID_BET_SIZE"
	IREC_INVALID_RUNNER                  EInstructionReportErrorCode = "INVALID_RUNNER"
	IREC_BET_TAKEN_OR_LAPSED             EInstructionReportErrorCode = "BET_TAKEN_OR_LAPSED"
	IREC_BET_IN_PROGRESS                 EInstructionReportErrorCode = "BET_IN_PROGRESS"
	IREC_RUNNER_REMOVED                  EInstructionReportErrorCode = "RUNNER_REMOVED"
	IREC_MARKET_NOT_OPEN_FOR_BETTING     EInstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BETTING"
	IREC_LOSS_LIMIT_EXCEEDED             EInstructionReportErrorCode = "LOSS_LIMIT_EXCEEDED"
	IREC_MARKET_NOT_OPEN_FOR_BSP_BETTING EInstructionReportErrorCode = "MARKET_NOT_OPEN_FOR_BSP_BETTING"
	IREC_INVALID_PRICE_EDIT              EInstructionReportErrorCode = "INVALID_PRICE_EDIT"
	IREC_INVALID_ODDS                    EInstructionReportErrorCode = "INVALID_ODDS"
	IREC_INSUFFICIENT_FUNDS              EInstructionReportErrorCode = "INSUFFICIENT_FUNDS"
	IREC_INVALID_PERSISTENCE_TYPE        EInstructionReportErrorCode = "INVALID_PERSISTENCE_TYPE"
	IREC_ERROR_IN_MATCHER                EInstructionReportErrorCode = "ERROR_IN_MATCHER"
	IREC_INVALID_BACK_LAY_COMBINATION    EInstructionReportErrorCode = "INVALID_BACK_LAY_COMBINATION"
	IREC_ERROR_IN_ORDER                  EInstructionReportErrorCode = "ERROR_IN_ORDER"
	IREC_INVALID_BID_TYPE                EInstructionReportErrorCode = "INVALID_BID_TYPE"
	IREC_INVALID_BET_ID                  EInstructionReportErrorCode = "INVALID_BET_ID"
	IREC_CANCELLED_NOT_PLACED            EInstructionReportErrorCode = "CANCELLED_NOT_PLACED"
	IREC_RELATED_ACTION_FAILED           EInstructionReportErrorCode = "RELATED_ACTION_FAILED"
	IREC_NO_ACTION_REQUIRED              EInstructionReportErrorCode = "NO_ACTION_REQUIRED"
	IREC_TIME_IN_FORCE_CONFLICT          EInstructionReportErrorCode = "TIME_IN_FORCE_CONFLICT"
	IREC_UNEXPECTED_PERSISTENCE_TYPE     EInstructionReportErrorCode = "UNEXPECTED_PERSISTENCE_TYPE"
	IREC_INVALID_ORDER_TYPE              EInstructionReportErrorCode = "INVALID_ORDER_TYPE"
	IREC_UNEXPECTED_MIN_FILL_SIZE        EInstructionReportErrorCode = "UNEXPECTED_MIN_FILL_SIZE"
	IREC_INVALID_CUSTOMER_ORDER_REF      EInstructionReportErrorCode = "INVALID_CUSTOMER_ORDER_REF"
	IREC_INVALID_MIN_FILL_SIZE           EInstructionReportErrorCode = "INVALID_MIN_FILL_SIZE"
	//ErrUnknownInstructionReportStatus    EInstructionReportStatus    = "Unknown instructionReportStatus value"
)

type ERollupModel string

const (
	RM_STAKE             ERollupModel = "STAKE"
	RM_PAYOUT            ERollupModel = "PAYOUT"
	RM_MANAGED_LIABILITY ERollupModel = "MANAGED_LIABILITY"
	RM_NONE              ERollupModel = "NONE"
	//ErrUnknownRollupModel ERollupModel = "Unknown rollupModel value"
)

type EGroupBy string

const (
	GB_EVENT_TYPE EGroupBy = "EVENT_TYPE"
	GB_EVENT      EGroupBy = "EVENT"
	GB_MARKET     EGroupBy = "MARKET"
	GB_SIDE       EGroupBy = "SIDE"
	GB_BET        EGroupBy = "BET"
	//ErrUnknownGroupBy EGroupBy = "Unknown groupBy value"
)

type EBetStatus string

const (
	BS_SETTLED   EBetStatus = "SETTLED"
	BS_VOIDED    EBetStatus = "VOIDED"
	BS_LAPSED    EBetStatus = "LAPSED"
	BS_CANCELLED EBetStatus = "CANCELLED"
	//ErrUnknownBetStatus EBetStatus = "Unknown betStatus value"
)

type EMarketType string

const (
	MT_A              EMarketType = "A"
	MT_L              EMarketType = "L"
	MT_O              EMarketType = "O"
	MT_R              EMarketType = "R"
	MT_NOT_APPLICABLE EMarketType = "NOT_APPLICABLE"
)

type ETimeInForce string

const (
	TIF_FILL_OR_KILL ETimeInForce = "FILL_OR_KILL"
)

type EBetTargetType string

const (
	BTT_BACKERS_PROFIT EBetTargetType = "BACKERS_PROFIT"
	BTT_PAYOUT         EBetTargetType = "PAYOUT"
)

type EPriceLadderType string

const (
	PLT_CLASSIC    EPriceLadderType = "CLASSIC"
	PLT_FINEST     EPriceLadderType = "FINEST"
	PLT_LINE_RANGE EPriceLadderType = "LINE_RANGE"
)
