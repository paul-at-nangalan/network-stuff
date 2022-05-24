package models

import (
	"database/sql"
	"github.com/paul-at-nangalan/errorhandler/handlers"
	"log"
)
import "github.com/paul-at-nangalan/db-util/migrator"

type ApikeysModel struct{
	get *sql.Stmt
	set *sql.Stmt
}

func NewApikesModel(db *sql.DB)*ApikeysModel{
	cols := map[string]string{
		"name": "text",
		"apikey": "text",
		"updated": "timestamp",
	}
	primary := []string{"name"}
	index := []string{"name"}
	mig := migrator.NewMigrator(db, migrator.DBTYPE_POSTGRES)
	mig.Migrate("create-api-key-table", "api_keys", cols, index, primary)

	inssql := `INSERT INTO api_keys (name, apikey, updated)
 				VALUES($1, $2, NOW())
				ON CONFLICT (name) DO UPDATE SET apikey=excluded.apikey`
	insstmt, err := db.Prepare(inssql)
	handlers.PanicOnError(err)

	getsql := `SELECT apikey FROM api_keys WHERE name=$1`
	getstmt, err := db.Prepare(getsql)
	handlers.PanicOnError(err)

	return &ApikeysModel{
		set: insstmt,
		get: getstmt,
	}
}

func (p *ApikeysModel)Set(name, key string){
	_, err := p.set.Exec(name, key)
	handlers.PanicOnError(err)
}

func (p *ApikeysModel)Get(name string) string{
	res, err := p.get.Query(name)
	handlers.PanicOnError(err)
	defer res.Close()

	if !res.Next(){
		log.Panicln("Invalid api key")
	}
	apikey := sql.NullString{}
	err = res.Scan(&apikey)
	handlers.PanicOnError(err)
	if !apikey.Valid{
		log.Panicln("Null api key")
	}
	return apikey.String
}