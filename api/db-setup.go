package main

import (
	"api/helpers"
	"database/sql"
	"fmt"
)

func migrationWrapper(tableName string, sql string, db *sql.DB) {
	fmt.Printf("CREATING: %s table\n", tableName)
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
	fmt.Printf("SUCCESS: %s table\n", tableName)
}

func SetupDb() {
	db := helpers.DbClient()
	migrationWrapper("Users", `CREATE TABLE IF NOT EXISTS Users (id VARCHAR PRIMARY KEY, password VARCHAR NOT NULL, email VARCHAR UNIQUE NOT NULL)`, db)
	migrationWrapper("Podcasts", `CREATE TABLE IF NOT EXISTS Podcasts (id VARCHAR PRIMARY KEY, title VARCHAR NOT NULL, description VARCHAR NOT NULL, url VARCHAR NOT NULL, image VARCHAR, user_id VARCHAR NOT NULL, FOREIGN KEY (user_id) REFERENCES Users(id) )`, db)
	migrationWrapper("Episodes", `CREATE TABLE IF NOT EXISTS Episodes (id VARCHAR PRIMARY KEY, title VARCHAR NOT NULL, description TEXT, url VARCHAR NOT NULL, image VARCHAR, keywords VARCHAR, publishDate BIGINT, author VARCHAR, episodeNumber SMALLINT, user_id VARCHAR NOT NULL, FOREIGN KEY (user_id) REFERENCES Users(id) )`, db)
	migrationWrapper("Sessions", `CREATE TABLE IF NOT EXISTS Sessions (id VARCHAR PRIMARY KEY, email VARCHAR NOT NULL, user_id VARCHAR NOT NULL, session_token VARCHAR UNIQUE NOT NULL, expires_at DATE NOT NULL)`, db)

}
