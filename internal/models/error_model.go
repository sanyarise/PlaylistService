package models

type ErrorNotFound struct {
}

func (e ErrorNotFound) Error() string {
	return ""
}

type ErrorAlreadyPlaying struct {
}

func (e ErrorAlreadyPlaying) Error() string {
	return ""
}
