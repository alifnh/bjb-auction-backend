package constant

type ContextKey string

const (
	ContextUserInfoKey ContextKey = "ctx-user-info"

	RoleUser  = "user"
	RoleAdmin = "admin"
)

const (
	ResponseMsgSuccessRegister       = "user registered successfully"
	ResponseMsgSuccessLogin          = "user logged in successfully"
	ResponseMsgSuccessForgotPassword = "if email exist, a reset password link will be sent"
	ResponseMsgSuccessResetPassword  = "reset password sucessfully"
	MessageResponseSuccess           = "success"
)
