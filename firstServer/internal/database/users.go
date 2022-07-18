package database

import (
	"errors"
	"time"
)

// User -
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, error := c.readDB()

	if error != nil {
		return User{}, error
	}

	if _, ok := db.Users[email]; ok {
		return User{}, errors.New("User already exists")
	}

	user := User{Email: email, Password: password, Name: name, Age: age, CreatedAt: time.Now().UTC()}

	db.Users[email] = user

	error = c.updateDB(db)

	return user, error
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {

	db, error := c.readDB()


	if error != nil {
		return User{}, error
	}

	if _, ok := db.Users[email]; !ok {
		return User{}, errors.New("user doesn't exist")
	}


	u := db.Users[email]
	u.Age = age
	u.Password = password
	u.Name = name

	db.Users[email] = u


	error = c.updateDB(db)

	return u, error
}

func (c Client) GetUser(email string) (User, error) {

	db, err := c.readDB()

	if err != nil {
		return User{}, err
	}

	if user, ok := db.Users[email]; ok {
		return user, nil
	}

	return User{}, errors.New("User not found")
}

func (c Client) DeleteUser(email string) error {

	db, err := c.readDB()

	if err != nil {
		return err
	}

	if _, ok := db.Users[email]; ok {
		delete(db.Users, email)
		err = c.updateDB(db)
		return err
	}



	return errors.New("User not found")
}

