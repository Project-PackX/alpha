package exceptions

type BaseException struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
}
