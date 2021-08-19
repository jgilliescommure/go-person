package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jgilliescommure/go-person/internal/biz"
)

type personRepo struct {
	data *Data
	log  *log.Helper
}

// NewPersonRepo .
func NewPersonRepo(data *Data, logger log.Logger) biz.PersonRepo {
	return &personRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *personRepo) CreatePerson(ctx context.Context, g *biz.Person) error {
	return nil
}

func (r *personRepo) UpdatePerson(ctx context.Context, g *biz.Person) error {
	return nil
}
