package errcode

type ErrorCode struct {
	Message string `json:"message"`
}

var (
	InternalServerError  = ErrorCode{Message: "Internal Server Error"}
	RegisteredEmail      = ErrorCode{Message: "Email has been registered"}
	InvalidRequest       = ErrorCode{Message: "Invalid Request"}
	ValidationError      = ErrorCode{Message: "Validation error"}
	InvalidEncryptedText = ErrorCode{Message: "Invalid encrypted text"}
	UserNotFound         = ErrorCode{Message: "User not found"}
	FailedGetUser        = ErrorCode{Message: "Failed to get user"}
	RegisteredUsername   = ErrorCode{Message: "Username has been used"}
	TokenNotFound        = ErrorCode{Message: "Token not found"}
	InvalidToken         = ErrorCode{Message: "Invalid Token"}
	FileError            = ErrorCode{Message: "File error"}
)
