package oidcmiddleware

var (
	Admin = "authentik Admins"
	User  = "User"
)

func MatchRole(groups []string) string {
	for _, r := range groups {
		if r == Admin {
			return "ADMIN"
		}
	}

	return "USER"
}
