package enum

import "errors"

var ErrUnknownExecutionReportStatus = errors.New("Unknown executionReportStatus value")

type ExecutionReportStatus int

const (
	ERS_SUCCESS ExecutionReportStatus = iota // Order processed successfully
	ERS_FAILURE                              // Order failed.

	// The order itself has been accepted, but at least one (possibly all) actions have generated errors.
	// This error only occurs for replaceOrders, cancelOrders and updateOrders operations.
	// The placeOrders operation will not return PROCESSED_WITH_ERRORS status as it is an atomic operation.
	ERS_PROCESSED_WITH_ERRORS
	ERS_TIMEOUT // Order timed out.
)

var ersItems = [...]string{
	"SUCCESS",
	"FAILURE",
	"PROCESSED_WITH_ERRORS",
	"TIMEOUT",
}

var ersMap = map[string]ExecutionReportStatus{
	ersItems[ERS_SUCCESS]:               ERS_SUCCESS,
	ersItems[ERS_FAILURE]:               ERS_FAILURE,
	ersItems[ERS_PROCESSED_WITH_ERRORS]: ERS_PROCESSED_WITH_ERRORS,
	ersItems[ERS_TIMEOUT]:               ERS_TIMEOUT,
}

func (code ExecutionReportStatus) String() string {
	return ersItems[code]
}

func (code *ExecutionReportStatus) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := ersMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownExecutionReportStatus
	}
	return err

}
