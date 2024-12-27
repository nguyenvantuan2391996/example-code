package constants

const (
	SomethingWentWrong          = "Something went wrong"
	ErrAuthorizationHeaderEmpty = "authorization header is empty"
	ErrInvalidAuthorizationType = "invalid authorization type"
	ErrInvalidToken             = "invalid token"
	ErrNotAuthorized            = "not authorized"
)

const (
	FormatGetEntityErr    = "get %v is failed with err %v"
	FormatCreateEntityErr = "create %v is failed with err %v"
	FormatUpdateEntityErr = "update %v is failed with err %v"
	FormatDeleteEntityErr = "delete %v is failed with err %v"
	FormatTaskErr         = "task %v is failed with err %v"
)

const (
	FormatBeginTask = "begin task %v with input %v"
	FormatBeginAPI  = "begin api %v..."
)
