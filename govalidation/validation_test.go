package govalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func TestAliasTag(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		ID     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	req := LoginRequest{
		Username: "NAGACHI",
		Password: "Password",
	}
	err := validate.Struct(req)

	assert.Nil(t, err)
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().Interface().(string)
	if !regexNumber.MatchString(value) {
		return false
	}
	return len(value) == length
}

func TestCustomValidationParam(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	req := Login{
		Phone: "8232193103",
		Pin:   "123456",
	}

	err := validate.Struct(req)
	assert.Nil(t, err)
}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_eq_ic", MustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,field_eq_ic=Email|field_eq_ic=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "nagachi@mail.com",
		Email:    "nagachi@mail.com",
		Phone:    "62878782382",
		Name:     "Nagachi",
	}
	err := validate.Struct(user)
	assert.Nil(t, err)
}

type RegisterReq struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegister(level validator.StructLevel) {
	req := level.Current().Interface().(RegisterReq)

	if req.Username == req.Email || req.Username == req.Phone {

	} else {
		// report on username tag in username field
		level.ReportError(req.Username, "username", "username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegister, RegisterReq{})

	req := RegisterReq{
		Username: "nagachi@mail.com",
		Email:    "nagachi@mail.com",
		Phone:    "8223273273",
		Password: "password",
	}

	err := validate.Struct(req)
	assert.Nil(t, err)
}
