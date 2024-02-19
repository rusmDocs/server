package exceptionCodes

import "fmt"

const (
	EntityNotFound = "%s not found"
	EntityExists   = "%s already exists"
	EntityInvalid  = "%s is invalid"
)

func MakeException(code string, entity string) string {
	return fmt.Sprintf(code, entity)
}
