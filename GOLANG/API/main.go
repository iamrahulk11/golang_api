package main

import (
	"errors"
	config "example/API/config"
	"example/API/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	USER_ID  string `json:"user_id"`
	USERNAME string `json:"username"`
	USER_IMG string `json:"user_img"`
}
type db_user struct {
	ID          int        `json:"id"`
	USER_ID     string     `json:"user_id"`
	USERNAME    string     `json:"username"`
	INSERTED_ON time.Time  `json:"inserted_on"`
	UPDATED_ON  *time.Time `json:"updated_on"`
}

var users = []user{
	{USER_ID: "1001", USERNAME: "John Do", USER_IMG: "/static/images/avatars/001.jpg"},
	{USER_ID: "1002", USERNAME: "John", USER_IMG: "/static/images/avatars/001.jpg"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, users)
}

func getUserById(c *gin.Context) {
	user_id := c.Param("user_id")

	userFound, err := getUser(user_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, userFound)
}

func db_getUser(c *gin.Context) {
	var db_users []db_user

	cfg := config.GetConfiguration()
	db, err := database.CreateConnection(cfg)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to connect db"})
		return
	}

	if err := db.Ping(); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to connect db"})
		return
	}

	query := `select id, user_id, username,inserted_on,updated_on from tbl_golang_user`

	rows, err := db.Query(query)
	defer db.Close()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "failed to get users"})
		return
	}

	for rows.Next() {
		db_user := db_user{}
		err = rows.Scan(&db_user.ID, &db_user.USERNAME, &db_user.USER_ID, &db_user.INSERTED_ON, &db_user.UPDATED_ON)
		if err != nil {
			// handle this error
			panic(err)
		}
		db_users = append(db_users, db_user)
	}

	c.IndentedJSON(http.StatusOK, db_users)
}

func getUser(user_id string) (*user, error) {
	for i, u := range users {
		if u.USER_ID == user_id {
			return &users[i], nil
		}
	}
	return nil, errors.New("no user found with that id")
}

func deleteUser(c *gin.Context) {
	user_id, ok := c.GetQuery("user_id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing query param 'id'"})
		return
	}

	if len(users) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no user present"})
		return
	}

	userFound, err := getUser(user_id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no such user present"})
		return
	}

	NewUsers := removerUser(userFound.USER_ID)

	c.IndentedJSON(http.StatusOK, NewUsers)
}

func removerUser(user_id string) *[]user {
	var NewUsers []user
	for i, u := range users {
		if u.USER_ID != user_id {
			NewUsers = append(NewUsers, users[i])
		}
	}
	return &NewUsers
}

func bulkInsert(c *gin.Context) {
	var multipleUser []user

	if err := c.Bind(&multipleUser); err != nil {
		return
	}

	users = append(users, multipleUser...)

	c.IndentedJSON(http.StatusCreated, users)
}

func main() {

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/user/:user_id", getUserById)
	router.DELETE("/removeUser", deleteUser)
	router.POST("/bulkinsert", bulkInsert)

	router.GET("/db_users", db_getUser)
	router.Run("localhost:3001")
}
