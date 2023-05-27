package entity

import (
	"time"

	"github.com/dddsphere/quanta/internal/core"
	"github.com/dddsphere/quanta/internal/model/vo"
)

type (
	List struct {
		core.ID
		description string
		items       []Item
		owner       User
		Audit
	}

	Item struct {
		core.ID
		description string
		dueDate     time.Time
		category    vo.Category
		tags        []vo.Tag
		places      []vo.Place
		status      []Status
		Audit
	}

	Status struct {
		core.ID
		status map[time.Time]vo.Status
		Audit
	}
)
