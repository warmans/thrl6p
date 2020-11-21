package model

import (
	v1 "github.com/warmans/thrl6p/server/gen/api/v1"
	"time"
)

type Patch struct {
	Id          string     `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Patch       string     `db:"data"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (p *Patch) Proto(permaLink string) *v1.Patch {
	return &v1.Patch{
		Id:          p.Id,
		Patch:       p.Patch,
		Name:        p.Name,
		Description: p.Description,
		Permalink:   permaLink,
	}
}
