package handler

type Handlers struct {
	StatCollection  DataCollectionHandler
	UserTracking    UserTrackingHandler
	UserTrackingSSE UserTrackingSSEHandler
	UserManager     UserManagerHandler
	Map             MapHandler
}
