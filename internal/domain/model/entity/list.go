package entity

import (
	"time"

	"github.com/dddsphere/quanta/internal/core"
	vo2 "github.com/dddsphere/quanta/internal/domain/model/vo"
)

type (
	List struct {
		core.ID
		description string
		items       []Item
		owner       User
		meta        Meta
		Audit
	}

	Item struct {
		core.ID
		description string
		dueDate  time.Time
		category vo2.Category
		tags     []vo2.Tag
		places   []vo2.Place
		status   []Status
		Audit
	}

	Status struct {
		core.ID
		status map[time.Time]vo2.Status
		Audit
	}
)


func (l *List)
