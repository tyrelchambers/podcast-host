package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"github.com/joho/godotenv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbadmin"
	password = "JUYVgv3vutcr4hjJd"
	dbname   = "db"
)

var BUNNY_URL_BASE = "https://podcast-files.b-cdn.net"

var DbUrl = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

func DbClient() *sql.DB {

	db, err := sql.Open("postgres", DbUrl)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func BunnyClient() *bunnystorage.Client {
	readOnlyKey := GoDotEnvVariable("BUNNYNET_READ_API_KEY")

	readWriteKey := GoDotEnvVariable("BUNNYNET_WRITE_API_KEY")

	cfg := &bunnystorage.Config{
		Application: &bunnystorage.Application{
			Name:    "Podcast Host",
			Version: "1.0.0",
			Contact: "tychambers3@gmail.com",
		},
		StorageZone: "podcast-files",
		Key:         readWriteKey,
		ReadOnlyKey: readOnlyKey,
		Endpoint:    bunnystorage.EndpointNewYork,
		Timeout:     5 * time.Minute,
	}

	client, err := bunnystorage.NewClient(cfg)

	// Create a new Client instance with the given Config.
	if err != nil {
		log.Fatal(err)
	}

	return client

}

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal(fmt.Sprintf("Key with name: %s, not found", key))
	}

	return env
}

func WriteFileAndUpload(r *http.Request, uErr chan error) (uploadPath string) {
	client := BunnyClient()

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		uErr <- err
		return
	}

	defer file.Close()

	tempFile, err := os.CreateTemp("temp-files", "upload-*.mp3")
	if err != nil {
		fmt.Println(err)
		uErr <- err
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	godotenv.Load(".env")

	newFile, err := os.Open(tempFile.Name())

	if err != nil {
		log.Fatal(err)
		uErr <- err
		return
	}

	// some-ID will be ID of podcast
	uploadPathUrl := fmt.Sprintf("some-id/%s", "audio")

	resp, err := client.Upload(context.Background(), "/", uploadPathUrl, "", newFile)

	if err != nil {
		log.Fatal(err)
		uErr <- err
		return
	}
	defer newFile.Close()

	fullPath := fmt.Sprintf("%s%s", BUNNY_URL_BASE, uploadPathUrl)

	fmt.Printf("Successfully Uploaded File to Bunny: %d\n", resp.Status)

	removeErr := os.Remove(tempFile.Name())

	if removeErr != nil {
		log.Fatal(removeErr)
		uErr <- err
		return
	}

	return fullPath
}
