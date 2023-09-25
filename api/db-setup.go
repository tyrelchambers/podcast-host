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
	migrationWrapper("Podcasts", `CREATE TABLE IF NOT EXISTS Podcasts (id VARCHAR PRIMARY KEY, title VARCHAR NOT NULL, description TEXT NOT NULL, thumbnail VARCHAR, explicit_content BOOLEAN, primary_category VARCHAR, secondary_category VARCHAR, author VARCHAR NOT NULL, copyright VARCHAR, keywords VARCHAR, website VARCHAR, language VARCHAR, timezone VARCHAR, show_owner VARCHAR NOT NULL, owner_email VARCHAR, display_email_in_rss BOOLEAN, user_id VARCHAR NOT NULL, FOREIGN KEY (user_id) REFERENCES Users(id) )`, db)
	migrationWrapper("Episodes", `CREATE TABLE IF NOT EXISTS Episodes (id VARCHAR PRIMARY KEY, title VARCHAR NOT NULL, description TEXT, url VARCHAR NOT NULL, image VARCHAR, keywords VARCHAR, publish_date BIGINT, author VARCHAR, episode_number SMALLINT, podcast_id VARCHAR NOT NULL, draft BOOLEAN, FOREIGN KEY (podcast_id) REFERENCES Podcasts(id) )`, db)
	migrationWrapper("Migrations", `CREATE TABLE IF NOT EXISTS Migrations (id VARCHAR PRIMARY KEY, name VARCHAR NOT NULL)`, db)
}
