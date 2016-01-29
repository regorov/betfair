package enum

import "errors"

var ErrUnknownInstructionReportStatus = errors.New("Unknown instructionReportStatus value")

type InstructionReportStatus int

const (
	IRS_SUCCESS InstructionReportStatus = iota
	IRS_FAILURE
	IRS_TIMEOUT
)

var instructionReportStatusItems = [...]string{
	"SUCCESS",
	"FAILURE",
	"TIMEOUT",
}

var instructionReportStatusMap = map[string]InstructionReportStatus{
	instructionReportStatusItems[IRS_SUCCESS]: IRS_SUCCESS,
	instructionReportStatusItems[IRS_FAILURE]: IRS_FAILURE,
	instructionReportStatusItems[IRS_TIMEOUT]: IRS_TIMEOUT,
}

func (code InstructionReportStatus) String() string {
	return instructionReportStatusItems[code]
}

func (code *InstructionReportStatus) UnmarshalJSON(buf []byte) error {
	var err error
	val, ok := instructionReportStatusMap[string(buf)]
	if ok {
		*code = val
	} else {
		err = ErrUnknownInstructionReportStatus
	}
	return err
}
