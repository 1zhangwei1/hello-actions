package code

type ErrorCode int64

const (
	ERROR_CODE_OK             ErrorCode = 0 //fail
	ERROR_CODE_INVALID_PARAMS ErrorCode = 1 //参数错误
	ERRER_CODE_TIMEOUT        ErrorCode = 2 //超时
)
