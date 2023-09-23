package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"github.com/joho/godotenv"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
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

func WriteFileAndUpload(r *http.Request, uErr chan error, file multipart.File, podcast_id string, file_name string) {
	client := BunnyClient()

	defer file.Close()

	tempFile, err := os.CreateTemp("temp-files", fmt.Sprintf("*.%s", file_name))

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

	iTemp := tempFile.Name()
	oFile := ConvertToMp3(file_name)

	fmt.Println(iTemp, oFile)

	err = ffmpeg_go.Input(iTemp).
		Output("./temp-files/formatted/" + oFile).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		fmt.Println(err)
		uErr <- err
		return
	}

	godotenv.Load(".env")

	newFile, err := os.Open(tempFile.Name())

	if err != nil {
		log.Fatal(err)
		uErr <- err
		return
	}
	//  TODO: GET FFMPEG AND UPLOAD TO WORK WITH TRANSCODE
	fmt.Println(tempFile.Name())

	resp, err := client.Upload(context.Background(), "/"+podcast_id, ConvertToMp3(file_name), "", newFile)

	if err != nil {
		log.Fatal(err)
		uErr <- err
		return
	}
	defer newFile.Close()

	fmt.Printf("Successfully Uploaded File to Bunny: %d\n", resp.Status)

	removeErr := os.Remove(tempFile.Name())
	removeErr = os.Remove("./temp-files/formatted/" + oFile)

	if removeErr != nil {
		log.Fatal(removeErr)
		uErr <- err
		return
	}

}

func ConvertToMp3(filename string) string {
	re := regexp.MustCompile(`\.[^./]+$`)
	mp3Filename := re.ReplaceAllString(filename, ".mp3")
	return mp3Filename
}
