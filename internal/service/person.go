package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	v1 "github.com/jgilliescommure/go-person/api/person"
	"github.com/jgilliescommure/go-person/internal/biz"
)

type PersonService struct {
	v1.UnimplementedPersonServer

	uc  *biz.PersonUsecase
	log *log.Helper
}

var db *sql.DB

func NewPersonService(uc *biz.PersonUsecase, logger log.Logger) *PersonService {
	return &PersonService{uc: uc, log: log.NewHelper(logger)}
}

func (s *PersonService) CreatePerson(ctx context.Context, req *v1.CreatePersonRequest) (*v1.CreatePersonReply, error) {
	CreateClient()
	result, err := db.Exec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", req.FirstName, req.LastName, req.Email)
	if err != nil {
		s.log.Error("CreatePerson: " + err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		s.log.Error("CreatePerson: " + err.Error())
	}
	s.log.Info("New ID: ", id)

	return &v1.CreatePersonReply{Id: int32(id)}, nil
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

func CreateClient() {
	cfg := mysql.Config{
		User:   "root",     //os.Getenv("DBUSER"),
		Passwd: "mysql123", //os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "health_records",
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	}
	fmt.Println("Connected!")
}
