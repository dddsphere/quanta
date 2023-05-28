package repo

import (
	"context"

	"github.com/dddsphere/quanta/internal/model/entity"
	"github.com/dddsphere/quanta/internal/model/vo"
)

type (
	User interface {
		GetByID(ctx context.Context, id string) error
		GetByUsername(ctx context.Context, username string) error
		GetByEmail(ctx context.Context, email string) error
	}

	List interface {
		GetByID(ctx context.Context, id string) (list entity.List, err error)
		GetBySlug(ctx context.Context, slug string) (list entity.List, err error)
		Save(ctx context.Context, list entity.List) (updated entity.List, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	Item interface {
		GetByID(ctx context.Context, id string) (list entity.Item, err error)
		GetBySlug(ctx context.Context, slug string) (list entity.Item, err error)
		Save(ctx context.Context, list entity.List) (updated entity.Item, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	Category interface {
		GetByName(ctx context.Context, name string) (category vo.Category, err error)
		Save(ctx context.Context, category vo.Category) (updated vo.Category, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	Place interface {
		GetByName(ctx context.Context, name string) (place vo.Place, err error)
		Save(ctx context.Context, place vo.Place) (updated vo.Place, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	Status interface {
		GetByName(ctx context.Context, name string) (place vo.Status, err error)
		Save(ctx context.Context, status vo.Status) (updated vo.Status, err error)
		Delete(ctx context.Context, id string) (err error)
	}

	Tag interface {
		GetByName(ctx context.Context, name string) (tag vo.Tag, err error)
		Save(ctx context.Context, tag vo.Tag) (updated vo.Tag, err error)
		Delete(ctx context.Context, id string) (err error)
	}
)
