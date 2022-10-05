package main

import (
	"errors"
	"fmt"
	"os"
)

type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

func obtainError(status int) (code int, err error) {
	if status >= 400 {
		return 500, &MyCustomError{
			StatusCode: 500,
			Message:    "Algo salió mal",
		}
	}
	return 200, nil
}

var errTest = errors.New("error Test 1")

func getError() error {
	return fmt.Errorf("informacion extra: %w", errTest)
}

type ErrType2 struct{}

func (e ErrType2) Error() string {
	return "soy el error 2"
}

func main() {
	status, err := obtainError(400)
	if status > 399 {
		err := errors.New("la petición ha fallado")
		fmt.Println(err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err1 := &MyCustomError{
		StatusCode: 502,
		Message:    "Soy el error 1",
	}

	err2 := &MyCustomError{
		StatusCode: 406,
		Message:    "Soy el error 2",
	}
	bothErrorsAreEqual := errors.As(err1, &err2)

	fmt.Println(bothErrorsAreEqual)
	error := getError()
	coincidence := errors.Is(error, errTest)
	fmt.Println(coincidence)
	err4 := ErrType2{}
	err3 := fmt.Errorf("soy el error 3, contengo al 4 %w", err4)

	fmt.Println(
		errors.Unwrap(err3),
	)

	fmt.Println(
		errors.Unwrap(err4),
	)
}
