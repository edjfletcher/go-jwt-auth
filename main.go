package go_jwt_auth

type JWTPayload struct {
	Iss string
	Sub string
	Aud string
	Iat int
	Exp int
	Uid string
}

func IsValidJWT() bool {
	return true
}
