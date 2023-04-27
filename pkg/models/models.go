// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package drive

import (
	"encoding/json"
	"time"
)

type WebappUser struct {
	UserID       int64           `db:"user_id" json:"userID"`
	UserName     string          `db:"user_name" json:"userName"`
	PassWordHash string          `db:"pass_word_hash" json:"passWordHash"`
	Name         string          `db:"name" json:"name"`
	Config       json.RawMessage `db:"config" json:"config"`
	CreatedAt    time.Time       `db:"created_at" json:"createdAt"`
	IsEnabled    bool            `db:"is_enabled" json:"isEnabled"`
}
