package users

import (
	"net/http"
	"strconv"

	"github.com/davidalvarezcastro/bookstore-users-api/models/users"
	"github.com/davidalvarezcastro/bookstore-users-api/services"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

// Create creates a new user in our app
func Create(c *gin.Context) {
	var user users.User

	// equivalent to:
	// 	bytes, err := ioutil.ReadAll(c.Request.Body)
	// 	if err != nil {}
	// 	if err := json.Unmarshal(bytes, &user); err != nil {}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get returns the information of a user
func Get(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("userid should be a number")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.Get(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Search finds users
func Search(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
