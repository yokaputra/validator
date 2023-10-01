// main.go

package main

import (
	"fmt"

	"github.com/yokaputra/validator"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func main() {
	// Create instance Validator
	validator := validator.NewValidator()

	// Create instance User struct
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	if err := validator.Validate(user); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Println("Validation passed")
	}
}
