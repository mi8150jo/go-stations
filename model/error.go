package model 

import "fmt"

type ErrNotFound struct{
	message string
}

func(err *ErrNotFound) Error() string{
	return fmt.Sprintf("Error Not Found: %s", err.message)
}

