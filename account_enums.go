// Account enums
// http://docs.developer.betfair.com/docs/display/1smk3cen4v3lu3yomq5qye0ni/Accounts+Enums
package betfair

type EWallet string

const (
	W_UK EWallet = "UK"
)

type ESubscriptionStatus string

const (
	// Any subscription status
	SS_ALL ESubscriptionStatus = "ALL"

	// Only activated subscriptions
	SS_ACTIVATED ESubscriptionStatus = "ACTIVATED"

	// Only unactivated subscriptions
	SS_UNACTIVATED ESubscriptionStatus = "UNACTIVATED"

	// Only cancelled subscriptions
	SS_CANCELLED ESubscriptionStatus = "CANCELLED"

	// Only expired subscriptions
	SS_EXPIRED ESubscriptionStatus = "EXPIRED"
)

type EIncludeItem string

const (
	IT_ALL                  EIncludeItem = "ALL"
	IT_DEPOSITS_WITHDRAWALS EIncludeItem = "DEPOSITS_WITHDRAWALS"
	IT_EXCHANGE             EIncludeItem = "EXCHANGE"
	IT_POKER_ROOM           EIncludeItem = "POKER_ROOM"
)

type EItemClass string

const (
	IC_UNKNOWN EItemClass = "UNKNOWN"
)

type ELoginStatus string

const (
	LS_SUCCESS                                 ELoginStatus = "SUCCESS"
	LS_INVALID_USERNAME_OR_PASSWORD            ELoginStatus = "INVALID_USERNAME_OR_PASSWORD"
	LS_ACCOUNT_NOW_LOCKED                      ELoginStatus = "ACCOUNT_NOW_LOCKED"
	LS_ACCOUNT_ALREADY_LOCKED                  ELoginStatus = "ACCOUNT_ALREADY_LOCKED"
	LS_PENDING_AUTH                            ELoginStatus = "PENDING_AUTH"
	LS_TELBET_TERMS_CONDITIONS_NA              ELoginStatus = "TELBET_TERMS_CONDITIONS_NA"
	LS_DUPLICATE_CARDS                         ELoginStatus = "DUPLICATE_CARDS"
	LS_SECURITY_QUESTION_WRONG_3X              ELoginStatus = "SECURITY_QUESTION_WRONG_3X"
	LS_KYC_SUSPEND                             ELoginStatus = "KYC_SUSPEND"
	LS_SUSPENDED                               ELoginStatus = "SUSPENDED"
	LS_CLOSED                                  ELoginStatus = "CLOSED"
	LS_SELF_EXCLUDED                           ELoginStatus = "SELF_EXCLUDED"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK    ELoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_DK"
	LS_NOT_AUTHORIZED_BY_REGULATOR_DK          ELoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_DK"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT    ELoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_IT"
	LS_NOT_AUTHORIZED_BY_REGULATOR_IT          ELoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_IT"
	LS_SECURITY_RESTRICTED_LOCATION            ELoginStatus = "SECURITY_RESTRICTED_LOCATION"
	LS_BETTING_RESTRICTED_LOCATION             ELoginStatus = "BETTING_RESTRICTED_LOCATION"
	LS_TRADING_MASTER                          ELoginStatus = "TRADING_MASTER"
	LS_TRADING_MASTER_SUSPENDED                ELoginStatus = "TRADING_MASTER_SUSPENDED"
	LS_AGENT_CLIENT_MASTER                     ELoginStatus = "AGENT_CLIENT_MASTER"
	LS_AGENT_CLIENT_MASTER_SUSPENDED           ELoginStatus = "AGENT_CLIENT_MASTER_SUSPENDED"
	LS_DANISH_AUTHORIZATION_REQUIRED           ELoginStatus = "DANISH_AUTHORIZATION_REQUIRED"
	LS_SPAIN_MIGRATION_REQUIRED                ELoginStatus = "SPAIN_MIGRATION_REQUIRED"
	LS_DENMARK_MIGRATION_REQUIRED              ELoginStatus = "DENMARK_MIGRATION_REQUIRED"
	LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED       ELoginStatus = "SPANISH_TERMS_ACCEPTANCE_REQUIRED"
	LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED    ELoginStatus = "ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED"
	LS_CERT_AUTH_REQUIRED                      ELoginStatus = "CERT_AUTH_REQUIRED"
	LS_CHANGE_PASSWORD_REQUIRED                ELoginStatus = "CHANGE_PASSWORD_REQUIRED"
	LS_PERSONAL_MESSAGE_REQUIRED               ELoginStatus = "PERSONAL_MESSAGE_REQUIRED"
	LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED ELoginStatus = "INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED"
	LS_EMAIL_LOGIN_NOT_ALLOWED                 ELoginStatus = "EMAIL_LOGIN_NOT_ALLOWED"
	LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL     ELoginStatus = "MULTIPLE_USERS_WITH_SAME_CREDENTIAL"
	LS_ACCOUNT_PENDING_PASSWORD_CHANGE         ELoginStatus = "ACCOUNT_PENDING_PASSWORD_CHANGE"
	LS_TEMPORARY_BAN_TOO_MANY_REQUESTS         ELoginStatus = "TEMPORARY_BAN_TOO_MANY_REQUESTS"
	ErrUnknownLoginStatus                      ELoginStatus = "Unknown loginStatus value"
)
