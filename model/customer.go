package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"table:customer,alias:c"`

	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	NIK           string    `bun:"nik" json:"niik"`
	Name          string    `bun:"name" json:"name"`
	PhoneNumber   string    `bun:"phone_number" json:"phone_number"`
	AccountNumber string    `bun:"account_number" json:"account_number"`
	Saldo         float64   `bun:"saldo" json:"saldo"`
	IsDeleted     bool      `bun:"is_deleted" json:"-"`
}
