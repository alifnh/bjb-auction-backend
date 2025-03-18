package constant

const (
	FailedToRegister              = "failed to register"
	EntityNotFoundErrorMessage    = "%s not found"
	InternalServerErrorMessage    = "currently our server is facing unexpected error, please try again later"
	UserAlreadyRegisteredError    = "another user already registered with the same email"
	ValidationError               = "input validation error"
	InvalidJsonUnmarshallError    = "invalid JSON format"
	JsonSyntaxError               = "invalid JSON syntax"
	InvalidJsonValueTypeError     = "invalid value for %s"
	FailedToHashPasword           = "failed to hash password"
	FailedToSaveVerificationToken = "failed to save verification token"
	InvalidVerificationToken      = "invalid verification token"
	FailedToSendVerificationEmail = "failed to send verification email"
	PasswordNotMatch              = "password not match"
	FailedToForgotPassword        = "failed to forgot password"
	FailedToResetPassword         = "failed to reset password"
	InvalidOrExpiredToken         = "invalid or expired token"
	InvalidCredentials            = "invalid credentials"
	FailedToLogin                 = "failed to login"
	Unauthorized                  = "unauthorized"
	UserAlreadyVerified           = "user already verified"
	FailedToExchangeToken         = "failed to exchange token"
	FailedToGetUserInfo           = "failed to get user info"
	FailedToDecodeUserInfo        = "failed to decode user info"
	RoleNotPermitted              = "role not permitted"
	FailedToGetAssetInfo          = "failed to get asset info"
)
