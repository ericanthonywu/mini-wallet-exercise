package Model

import (
	"github.com/google/uuid"
	"time"
)

type WalletResponse struct {
	ID        uuid.UUID `json:"id"`
	OwnedBy   uuid.UUID `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   uint64    `json:"balance"`
}

type DepositResponse struct {
	ID          uuid.UUID `json:"id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Amount      uint64    `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

type WithdrawalResponse struct {
	ID          uuid.UUID `json:"id"`
	WithdrawnBy uuid.UUID `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Amount      string    `json:"amount"`
	ReferenceId string    `json:"reference_id"`
}

type WalletDisabledResponse struct {
	ID         uuid.UUID `json:"id"`
	OwnedBy    uuid.UUID `json:"owned_by"`
	Status     string    `json:"status"`
	DisabledAt time.Time `json:"disabled_at"`
	Balance    uint64    `json:"balance"`
}
