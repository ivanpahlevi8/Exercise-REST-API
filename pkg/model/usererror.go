package model

type ErrorUser struct {
	ErrMessage string `json:"error"`
	TimeStamp  string `json:"time_stamp"`
}

func (e *ErrorUser) GetErrMessage() string {
	return e.ErrMessage
}

func (e *ErrorUser) GetTimeStampe() string {
	return e.TimeStamp
}

func (e *ErrorUser) SetErrMessage(errMessage string) {
	e.ErrMessage = errMessage
}

func (e *ErrorUser) SetTimeStamp(timeStamp string) {
	e.TimeStamp = timeStamp
}
