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

	result, saveErr := services.Create(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Get returns the information of a user
func Get(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
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

	result, updateErr := services.Update(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete removes a user
func Delete(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	deleteErr := services.Delete(userID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search finds users
func Search(c *gin.Context) {
	status := c.Query("status")

	users, searchErr := services.Search(status)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}

	c.JSON(http.StatusOK, users)
}
