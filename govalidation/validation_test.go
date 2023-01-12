package govalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	var val *validator.Validate = validator.New()
	assert.NotNil(t, val)
}

func TestValidationTwoVariable(t *testing.T) {
	val := validator.New()
	password := "password"
	wrongPassword := "wrong"

	err := val.VarWithValue(password, wrongPassword, "eqfield")
	assert.NotNil(t, err)
}

func TestValidationTagParameter(t *testing.T) {
	val := validator.New()
	user := "12121"
	err := val.Var(user, "required,min=5,max=40")
	assert.Nil(t, err)
}

func TestValidationStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginReq := LoginRequest{
		Username: "nagachi@mail.com",
		Password: "nagachi",
	}
	err := validate.Struct(loginReq)
	assert.Equal(t, err, nil)
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginReq := LoginRequest{
		Username: "nagachi",
		Password: "naga",
	}
	err := validate.Struct(loginReq)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			fmt.Println("error", fieldErr.Field(), "on tag", fieldErr.Tag(), "with error", fieldErr.Error())
		}
	}
	assert.NotNil(t, err)
}

func TestCrossField(t *testing.T) {
	type RegisterReq struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	registerReq := RegisterReq{
		Username:        "nagachi@mail.com",
		Password:        "nagac",
		ConfirmPassword: "nagac",
	}
	err := validate.Struct(registerReq)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldErr := range validationErrors {
			fmt.Println("error", fieldErr.Field(), "on tag", fieldErr.Tag(), "with error", fieldErr.Error())
		}
	}
	assert.Nil(t, err)
}
