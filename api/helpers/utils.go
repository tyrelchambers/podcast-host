package helpers

import (
	"api/constants"
	"api/model"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func InitDb() {
	var err error

	db, err = gorm.Open(postgres.Open(constants.DbUrl), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.AutoMigrate(&model.User{}, &model.Podcast{}, &model.Episode{})

	fmt.Println("Successfully connected to the database.")
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

func WriteFileAndUpload(c echo.Context, file multipart.File, podcast_id string, file_name string) {
	client := BunnyClient()

	defer file.Close()

	tempFile, err := os.CreateTemp("temp-files", fmt.Sprintf("*.%s", file_name))

	if err != nil {
		fmt.Println(err)
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

	err = ffmpeg_go.Input(iTemp).
		Output("./temp-files/formatted/" + oFile).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		fmt.Println(err)
		return
	}

	godotenv.Load(".env")

	newFile, err := os.Open(tempFile.Name())

	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := client.Upload(context.Background(), "/"+podcast_id, ConvertToMp3(file_name), "", newFile)

	if err != nil {
		log.Fatal(err)
		return
	}
	defer newFile.Close()

	fmt.Printf("Successfully Uploaded File to Bunny: %d\n", resp.Status)

	removeErr := os.Remove(tempFile.Name())
	removeErr = os.Remove("./temp-files/formatted/" + oFile)

	if removeErr != nil {
		log.Fatal(removeErr)
		return
	}

}

func ConvertToMp3(filename string) string {
	re := regexp.MustCompile(`\.[^./]+$`)
	mp3Filename := re.ReplaceAllString(filename, ".mp3")
	return mp3Filename
}

func Contains(s []string, val string) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

func ConvertToUnix(date string) uint64 {
	i, err := strconv.ParseUint(date, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ConvertToDto(model interface{}, dto interface{}) {
	modelValue := reflect.ValueOf(model)
	dtoValue := reflect.ValueOf(dto).Elem()

	for i := 0; i < dtoValue.NumField(); i++ {
		dtoField := dtoValue.Type().Field(i)
		modelName := dtoField.Name

		if modelName != "" {
			modelField := modelValue.FieldByName(modelName)

			if modelField.IsValid() {
				dtoValue.Field(i).Set(modelField)
			}
		}
	}
}
