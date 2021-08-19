package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Person struct {
	Id        int32
	FirstName string
	LastName  string
	Email     string
}

type PersonRepo interface {
	CreatePerson(context.Context, *Person) error
	UpdatePerson(context.Context, *Person) error
}

type PersonUsecase struct {
	repo PersonRepo
	log  *log.Helper
}

func NewPersonUsecase(repo PersonRepo, logger log.Logger) *PersonUsecase {
	return &PersonUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (pc *PersonUsecase) Create(ctx context.Context, p *Person) error {
	return pc.repo.CreatePerson(ctx, p)
}

func (pc *PersonUsecase) Update(ctx context.Context, p *Person) error {
	return pc.repo.UpdatePerson(ctx, p)
}
