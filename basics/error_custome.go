package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Error Message Custome"}
	}

	if id != "afan" {
		return &notFoundError{"data not found"}
	}

	// oke

	return nil
}

func main() {
	err := SaveData("afa", nil)

	if err != nil {
		// terjadi err
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundErr.Error())
		} else {
			fmt.Println("Error not found")
		}
	} else {
		// terjadi sukses
		fmt.Println("sukses")
	}
}