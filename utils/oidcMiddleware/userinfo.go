package oidcmiddleware

type UserInfo struct {
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"email_verified"`
	Name              string   `json:"name"`
	GivenName         string   `json:"given_name"`
	PreferredUsername string   `json:"preferred_username"`
	Nickname          string   `json:"nickname"`
	Groups            []string `json:"groups"`
	Sub               string   `json:"sub"`
}

func (u UserInfo) IsAdmin() bool {
	return MatchRole(u.Groups) == "ADMIN"
}

func (u UserInfo) IsUser() bool {
	return MatchRole(u.Groups) == "USER"
}
