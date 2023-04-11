package users

// user represents data about a users data.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
}
