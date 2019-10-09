package users

import (
	"fmt"
)

func (s *Storage) CreateUser (name string, interests []string) error{
	if len(name) < 0 {
		return fmt.Errorf("users' name is not provided")
	} else if len(interests) < 1 {
		return fmt.Errorf("users' interests are not provided")

	}

	//TODO might not work properly
	_, err := s.Db.Exec("INSERT INTO users (name, users_interests) VALUES ($1, $2)")
	fmt.Println("INSERT INTO users (name, users_interests) VALUES ($1, $2)")
	return err
}
