package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	//if the user is USER and not ADMIN he cannot access others resource so we check if uid == userId
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resoure")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, userType string) (err error) {
	role := c.GetString("user_type")
	err = nil

	//check if the role is correct or not
	if role != userType{
		err = errors.New("Unauthorized to access this resoure")
		return err
	}
	return err
}