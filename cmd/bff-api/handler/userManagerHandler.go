package handler

import (
	"encoding/json"
	"log"
	"net/http"

	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	usermanagerclient "github.com/ZecretBone/ips-bff/internal/repository/grpc/userManagerClient"
	oidcmiddleware "github.com/ZecretBone/ips-bff/utils/oidcMiddleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type UserManagerHandler interface {
	GetCoordinate(ctx *gin.Context)
}

type userManagerHandler struct {
	userManagerClient usermanagerclient.Service
}

func ProvideUserManagerHandler(userManagerClient usermanagerclient.Service) UserManagerHandler {
	return &userManagerHandler{
		userManagerClient: userManagerClient,
	}
}

func messageHandler(msg []byte) (*userv1.GetCoordinateRequest, error) {

	var message userv1.GetCoordinateRequest
	err := json.Unmarshal(msg, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (rs *userManagerHandler) GetCoordinate(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	role := ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)]
	is_admin := false
	if role == 1 {
		is_admin = true
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if messageType == websocket.TextMessage {
			message, err := messageHandler(p)
			if err != nil {
				continue
			}
			message.User = userInfo.Name
			message.IsAdmin = is_admin
			predictedCoordinate, err := rs.userManagerClient.GetCoordinate(ctx, message)
			if err != nil {
				continue
			}
			jsonMessage, err := json.Marshal(predictedCoordinate)
			if err != nil {
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				continue
			}
		}

	}

	//ctx.JSON(http.StatusOK, "success")
}
