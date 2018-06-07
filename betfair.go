package betfair

import (
	"time"
)

// BetfairRestURL type of all betfair rest urls
type BetfairRestURL string

const (
	CertURL                     = "https://identitysso-api.betfair.com/api/certlogin"
	KeepAliveURL                = "https://identitysso.betfair.com/api/keepAlive"
	AccountURL   BetfairRestURL = "https://api.betfair.com/exchange/account/rest/v1.0"
	BettingURL   BetfairRestURL = "https://api.betfair.com/exchange/betting/rest/v1.0"
	ScoresURL                   = "https://api.betfair.com/exchange/scores/json-rpc/v1"
)

type Betfair struct {
	*Client
	*Betting
}

func NewBetfair(apikey string) *Betfair {

	client := Client{ApiKey: apikey}

	return &Betfair{
		Client:  &client,
		Betting: &Betting{&client},
	}
}

// TODO: Deprecate
var NewBet = NewBetfair

type Time time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	tmp, err := time.Parse(`"2006-01-02T15:04:05.999Z0700"`, string(data))
	if err != nil {
		return err
	}
	*t = Time(tmp)

	return nil
}
