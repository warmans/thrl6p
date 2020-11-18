package model

import "time"

type Patch struct {
	Id          string     `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Patch       string     `db:"data"`
	CreatedAt   *time.Time `db:"created_at"`
}
