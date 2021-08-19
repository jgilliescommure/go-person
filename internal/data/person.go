package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"github.com/jgilliescommure/go-person/internal/biz"
)

var db *sql.DB

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

func CreateClient() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
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

func (r *personRepo) CreatePerson(ctx context.Context, p *biz.Person) error {
	CreateClient()
	result, err := db.Exec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", p.FirstName, p.LastName, p.Email)
	if err != nil {
		fmt.Println("CreatePerson: " + err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("CreatePerson: " + err.Error())
		return err
	}
	fmt.Println("New ID: " + strconv.FormatInt(int64(id), 10))
	return nil
}

func (r *personRepo) UpdatePerson(ctx context.Context, g *biz.Person) error {
	return nil
}
