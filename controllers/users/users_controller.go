package users

import (
	"net/http"
	"strconv"

	"github.com/davidalvarezcastro/bookstore-users-api/models/users"
	"github.com/davidalvarezcastro/bookstore-users-api/services"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDPram string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userIDPram, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userID, nil
}

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

	result, saveErr := services.UserService.Create(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Get returns the information of a user
func Get(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UserService.Get(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

// Update updates info of a user
func Update(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.ID = userID
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UserService.Update(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// Delete removes a user
func Delete(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	deleteErr := services.UserService.Delete(userID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search finds users
func Search(c *gin.Context) {
	status := c.Query("status")

	users, searchErr := services.UserService.Search(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
