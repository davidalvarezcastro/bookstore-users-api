package users

import (
	"net/http"
	"strconv"

	"github.com/davidalvarezcastro/bookstore-oauth-go/oauth"
	"github.com/davidalvarezcastro/bookstore-users-api/models/users"
	"github.com/davidalvarezcastro/bookstore-users-api/services"
	errorsutils "github.com/davidalvarezcastro/bookstore-utils-go/rest_errors"
	"github.com/gin-gonic/gin"
)

func getUserID(userIDPram string) (int64, *errorsutils.RestErr) {
	userID, err := strconv.ParseInt(userIDPram, 10, 64)
	if err != nil {
		return 0, errorsutils.NewBadRequestError("user id should be a number")
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
		restErr := errorsutils.NewBadRequestError("invalid json body")
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
	// if we want to force to authenticate a user
	// if callerID := oauth.GetCallerID(c.Request); callerID == 0 {
	// 	err := errorsutils.RestErr{
	// 		Status:  http.StatusUnauthorized,
	// 		Message: "resource not available",
	// 	}
	// 	c.JSON(err.Status, err)
	// 	return
	// }

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

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

	if oauth.GetCallerID(c.Request) == user.ID {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
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
		restErr := errorsutils.NewBadRequestError("invalid json body")
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

// Login tries to login user
func Login(c *gin.Context) {
	var request users.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errorsutils.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
