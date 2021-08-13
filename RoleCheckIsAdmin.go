package go_jwt_auth

func RoleCheckIsAdmin(role string) bool {
	return role == "admin"
}
