package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg string
	err error
}

func (ce CustomError) Error() string {
	return ce.Msg + ": " + ce.err.Error()
}

func temp1() {
	ce := CustomError{Msg: "This is the custom error", err: errors.New("New kind of error")}
	fmt.Println(ce)
}
