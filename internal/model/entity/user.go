package entity

import "github.com/dddsphere/quanta/internal/core"

type (
	User struct {
		core.ID
		username string
		email    string
		meta     Meta
		Audit
	}
)
