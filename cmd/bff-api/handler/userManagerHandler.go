package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ZecretBone/ips-bff/cmd/bff-api/mapper"
	userv1 "github.com/ZecretBone/ips-bff/internal/gen/proto/ips/user/v1"
	"github.com/ZecretBone/ips-bff/internal/models/request"
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

func (rs *userManagerHandler) GetSingleCoordinate(ctx *gin.Context) {

	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
		return
	}

	role := ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)]
	is_admin := false
	if role == 1 {
		is_admin = true
	}

	var body request.GetSingleCoordinateRequest
	if err := ctx.BindJSON(&body); err != nil {
		//err
		ctx.JSON(http.StatusBadRequest, err)
	}

	data := mapper.ToGetCoordinateRequest(body, userInfo.Name, is_admin)

	predictedCoordinate, err := rs.userManagerClient.GetCoordinate(ctx, data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, mapper.ToGetCoordinateResponse(predictedCoordinate))

}

func (rs *userManagerHandler) GetCoordinate(ctx *gin.Context) {
	userInfo, err := oidcmiddleware.GetUserInfoFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get user info"})
		return
	}

	role := ToUserV1RoleEnum[oidcmiddleware.MatchRole(userInfo.Groups)]
	is_admin := false
	if role == 1 {
		is_admin = true
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		if messageType == websocket.TextMessage {
			message, err := messageHandler(p)
			if err != nil {
				log.Println("Message parsing error:", err)
				continue
			}
			message.User = userInfo.Name
			message.IsAdmin = is_admin
			predictedCoordinate, err := rs.userManagerClient.GetCoordinate(ctx, message)
			if err != nil {
				log.Println("GetCoordinate error:", err)
				continue
			}
			jsonMessage, err := json.Marshal(predictedCoordinate)
			if err != nil {
				log.Println("JSON marshaling error:", err)
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				log.Println("WebSocket write error:", err)
				continue
			}
		}
	}
}
