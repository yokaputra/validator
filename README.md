# Go Custom Validator

Go Custom Validator is a package used for data validation within your application. This package utilizes the [go-playground/validator](https://pkg.go.dev/github.com/go-playground/validator/v10) library to implement validation based on struct tags in your data structures.

## Installation

To install this package, you can use the go get command:

```shell
go get -u github.com/yokaputra/validator
```

Then import the validator package into your own code.

```go
import "github.com/yokaputra/validator"
```

## Usage

Here's an example of how to use this package:

```go
// Create instance Validator
validator := validator.NewValidator()

// Create instance User struct
user := User{
    Name:  "John Doe",
    Email: "johndoe@example.com",
}

// Validate user
err := validator.Validate(user);
```

## Contribution

---

To contrib to this project, you can open a PR or an issue.

## License

This project is licensed under the Apache License 2.0. Please refer to the LICENSE file for more information.
