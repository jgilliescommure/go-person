package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/jgilliescommure/go-person/api/person"
	"github.com/jgilliescommure/go-person/internal/biz"
)

type PersonService struct {
	v1.UnimplementedPersonServer

	uc  *biz.PersonUsecase
	log *log.Helper
}

func NewPersonService(uc *biz.PersonUsecase, logger log.Logger) *PersonService {
	return &PersonService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PersonService) CreatePerson(ctx context.Context, req *v1.CreatePersonRequest) (*v1.CreatePersonReply, error) {
	return &v1.CreatePersonReply{Id: 11}, nil
}
func (s *PersonService) UpdatePerson(ctx context.Context, req *v1.UpdatePersonRequest) (*v1.PersonReply, error) {
	return &v1.PersonReply{}, nil
}
func (s *PersonService) GetPerson(ctx context.Context, req *v1.GetPersonRequest) (*v1.PersonReply, error) {
	return &v1.PersonReply{}, nil
}
func (s *PersonService) ListPersons(ctx context.Context, req *v1.ListPersonRequest) (*v1.ListPersonReply, error) {
	return &v1.ListPersonReply{}, nil
}
