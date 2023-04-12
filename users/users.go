package users

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http"
	"strconv" // package used to covert string into int type

	"github.com/anirudhmpai/database"
	"github.com/anirudhmpai/middleware"
	"github.com/gin-gonic/gin"

	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
)

// CreateUser create a user in the postgres db
func CreateUser(c *gin.Context) {

	// create an empty user of type User
	var user User

	// decode the json request to user
	err := c.BindJSON(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := insertUser(user)

	// format a middleware.Response object
	res := middleware.Response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the middleware.Response
	c.IndentedJSON(http.StatusOK, res)
}

// GetUser will return a single user by its id
func GetUser(c *gin.Context) {
	// get the userid from the request params, key is "id"
	param := c.Param("id")
	// convert the id type from string to int
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	user, err := getUser(int64(id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}

	// send the middleware.Response
	c.IndentedJSON(http.StatusOK, user)
}

// GetAllUser will return all the users
func GetAllUser(c *gin.Context) {

	// get all the users in the db
	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// send all the users as middleware.Response
	c.IndentedJSON(http.StatusOK, users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(c *gin.Context) {

	// get the userid from the request params, key is "id"
	param := c.Param("id")
	// convert the id type from string to int
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty user of type User
	var user User

	// decode the json request to user
	err = json.NewDecoder(c.Request.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update user to update the user
	updatedRows := updateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the middleware.Response message
	res := middleware.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the middleware.Response
	c.IndentedJSON(http.StatusOK, res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(c *gin.Context) {

	// get the userid from the request params, key is "id"
	param := c.Param("id")
	// convert the id type from string to int
	id, err := strconv.Atoi(param)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := middleware.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the middleware.Response
	c.IndentedJSON(http.StatusOK, res)
}

// ------------------------- handler functions ----------------
// insert one user in the DB
func insertUser(user User) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one user from the DB by its userid
func getUser(id int64) (User, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create a user of User type
	var user User

	// create the select sql query
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

// get one user from the DB by its userid
func getAllUsers() ([]User, error) {
	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	var usersData []User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var user User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		usersData = append(usersData, user)

	}

	// return empty user on error
	return usersData, err
}

// update user in the DB
func updateUser(id int64, user User) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE users SET name=$2, email=$3, age=$4 WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Name, user.Email, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteUser(id int64) int64 {

	// create the postgres db connection
	db := database.CreateConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM users WHERE id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
