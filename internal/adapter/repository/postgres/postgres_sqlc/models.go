// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgres_sqlc

import (
	"time"
)

type UserTask struct {
	WalletAddress string
	Amount        int64
	Point         int32
	EventName     string
	CreateTime    time.Time
	UpdateTime    time.Time
}
