package config

const (
	SUCCESS                   = 200
	ERROR                     = 500
)

var ExceptionMessage = map[int]string{
	SUCCESS:                   "成功",
	ERROR:                     "發生未知錯誤",
}

func GetExceptionMessage(StatusCode int) string {
	message, ok := ExceptionMessage[StatusCode]
	if ok {
		return message
	}

	return ExceptionMessage[ERROR]
}
