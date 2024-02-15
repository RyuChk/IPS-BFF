package handler

type UserManagerHandler interface {
}

type userManagerHandler struct {
}

func ProvideUserManagerHandler() UserManagerHandler {
	return &userManagerHandler{}
}
