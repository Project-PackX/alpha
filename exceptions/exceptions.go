package exceptions

import "time"

type baseException struct {
	Message   string `json:"message"`
	TimeStamp string `json:"timestamp"`
}

func CreateBaseException(message string) baseException {
	return baseException{
		Message:   message,
		TimeStamp: time.Now().String(),
	}
}

type InvalidInputException struct {
	baseException
}

func CreateInvalidInputException(message string) InvalidInputException {
	return InvalidInputException{
		baseException: CreateBaseException(message),
	}
}

type UserAlreadyExistsException struct {
	baseException
}

func CreateUserAlreadyExistsException(message string) UserAlreadyExistsException {
	return UserAlreadyExistsException{
		baseException: CreateBaseException(message),
	}
}
