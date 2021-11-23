package usecase

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DBConfigure interface {
	createDomainDB(domainName string) error
}

type DBConfiguration struct {
	log zap.SugaredLogger

}

func (dbc DBConfiguration) createDomainDB(domainName string) error  {
	// TODO update the credentials
	connInfo := "user=postgres password=yourpassword host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		dbc.log.Error("error on opening connection to database %v", err)
		return err
	}
	_, err = db.Exec("create database " + domainName)
	if err != nil {
		dbc.log.Error("error on creating database database %v", err)
		return err
	}
	return nil
}






