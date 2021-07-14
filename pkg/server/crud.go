package server

import (
	"context"
	"errors"
)

/*

CRUD-functions:

	addUser(usr User) error --- adds the User to the database, returns a non-nil error in case of failure.
		Can duplicate existing users.

	getUser(name string) (User, error) --- returns the User if he exists in database. Otherwise returns a
		non-nil error and an empty User structure. In case of duplicates, selects only the first row and
		ignores the rest.

*/

type User struct {
	Name string `json:"Name"`
}

func addUser(usr User) error {
	db := GetDB()
	query := "INSERT INTO users (name) VALUES ($1);"

	_, err := db.Exec(context.Background(), query, usr.Name)
	if err != nil {
		return err
	}

	return nil
}

func getUser(name string) (User, error) {
	var usr User
	db := GetDB()
	query := "SELECT * FROM users WHERE name=$1;"

	row := db.QueryRow(context.Background(), query, name)
	row.Scan(&usr.Name)
	if usr.Name == "" {
		return User{}, errors.New("error: user not found")
	}

	return usr, nil
}