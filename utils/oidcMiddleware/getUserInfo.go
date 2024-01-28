package oidcmiddleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserInfoFromContext(c *gin.Context) (UserInfo, error) {
	userInfo, exist := c.Get("UserInfo")
	if !exist {
		return UserInfo{}, errors.New("there is no UserInfo from middleware")
	}

	if asserted, ok := userInfo.(UserInfo); ok {
		return asserted, nil
	} else {
		return UserInfo{}, errors.New("input in UserInfo is not type UserInfo")
	}
}
