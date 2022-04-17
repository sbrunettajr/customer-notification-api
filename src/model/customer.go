package model

import "github.com/google/uuid"

type Customer struct {
	ID          int       `json:"-"`
	UUID        uuid.UUID `json:"uuid"`
	CNPJ        string    `json:"cnpj"`
	FantasyName string    `json:"fantasy_name"`
}
