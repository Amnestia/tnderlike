package paseto

// Payload struct containing auth data
type Payload struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	TokenType string `json:"token_type"`
}

// CtxKey context key for auth
type CtxKey string

// claims constant
const (
	AccessToken  = "Access-Token"
	RefreshToken = "Refresh-Token"

	audience = "Tnderlike - User"
	issuer   = "Tnderlike - Auth"
	jti      = "Tnderlike"
	footer   = "Tnderlike - Auth Token"

	payloadKey   = "Payload"
	tokenTypeKey = "Token-Type"

	AuthData CtxKey = "AuthData"
)
