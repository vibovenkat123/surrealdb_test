package main

import (
	"fmt"
	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	db, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}
	if _, err = db.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	}); err != nil {
		panic(err)
	}
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}

	user := User{
		Name:    "Vaibhav",
		Surname: "Venkat",
	}
	data, err := db.Create("user", user)
	if err != nil {
		panic(err)
	}
	createdUser := make([]User, 1)
	err = surrealdb.Unmarshal(data, &createdUser)
	if err != nil {
		panic(err)
	}
	data, err = db.Select(createdUser[0].ID)
	if err != nil {
		panic(err)
	}

	selectedUser := new(User)
	err = surrealdb.Unmarshal(data, &selectedUser)
	if err != nil {
		panic(err)
	}
	changes := map[string]string{"name": "Vibo"}
	if _, err = db.Change(selectedUser.ID, changes); err != nil {
		panic(err)
	}

	if _, err = db.Query("SELECT * FROM $rec", map[string]interface{}{
		"rec": createdUser[0].ID,
	}); err != nil {
		panic(err)
	}
	if _, err = db.Delete(selectedUser.ID); err != nil {
		panic(err)
	}
	fmt.Println(*selectedUser)
}
