package models

import (
	"github.com/google/uuid"
)

type Song struct {
	Id       uuid.UUID
	Title    string
	Duration uint32
}
