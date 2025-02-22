package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transaction,alias:t"`

	ID         uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	CustomerID uuid.UUID `bun:"customer_id,notnull" json:"customer_id"`
	Type       string    `bun:"type,notnull" json:"type"`
	Nominal    float64   `bun:"nominal" json:"nominal"`
	CreatedAt  time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	IsDeleted  bool      `bun:"is_deleted" json:"-"`

	Customer *Customer `bun:"rel:has-one,join:customer_id=id" json:"customer,omitempty"`
}
