package main

import (
	"database/sql"
	"github/heimaolst/simplebank/api"
	db "github/heimaolst/simplebank/db/sqlc"
	"github/heimaolst/simplebank/util"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)

	server := api.NewServer(store)

	server.Start(config.ServerAddress)
}
