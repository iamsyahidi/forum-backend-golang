package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/iamsyahidi/forum-backend-golang/api/controllers"
	"github.com/iamsyahidi/forum-backend-golang/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()
	os.Exit(m.Run())
}

func Database() {
	var err error

	TEST_DB_DRIVER := os.Getenv("TEST_DB_DRIVER")

	if TEST_DB_DRIVER == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_NAME"))
		server.DB, err = gorm.Open(TEST_DB_DRIVER, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database \n", TEST_DB_DRIVER)
			log.Fatal("This is the error : ", err)
		} else {
			fmt.Printf("We are connected to the %s database \n", TEST_DB_DRIVER)
		}
	}

	if TEST_DB_DRIVER == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
		server.DB, err = gorm.Open(TEST_DB_DRIVER, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TEST_DB_DRIVER)
			log.Fatal("This is the error : ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TEST_DB_DRIVER)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()
	user := models.User{
		Nickname: "iams",
		Email:    "iams@gmail.com",
		Password: "p455word",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table : %v", err)
	}

	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Nickname: "Ilham Syahidi",
			Email:    "iamsyahidi@gmail.com",
			Password: "p455word",
		},
		models.User{
			Nickname: "Dummy",
			Email:    "dummy@gmail.com",
			Password: "p455word",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func refreshUserAndPostTable() error {

	err := server.DB.DropTableIfExists(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	err = server.DB.AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		return err
	}

	log.Printf("Successfully refreshed tables")
	return nil
}

func seedOneUserAndOnePost() (models.Post, error) {

	err := refreshUserAndPostTable()
	if err != nil {
		return models.Post{}, err
	}

	user := models.User{
		Nickname: "Dummy iam",
		Email:    "dummyiam@gmail.com",
		Password: "p455word",
	}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.Post{}, err
	}

	post := models.Post{
		Title:    "This is the title dummy iam",
		Content:  "This is the content dummy iam",
		AuthorID: user.ID,
	}
	err = server.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func seedUsersAndPosts() ([]models.User, []models.Post, error) {

	var err error
	if err != nil {
		return []models.User{}, []models.Post{}, err
	}

	var users = []models.User{
		models.User{
			Nickname: "Ucil",
			Email:    "ucil@gmail.com",
			Password: "p455word",
		},
		models.User{
			Nickname: "Panci",
			Email:    "panci@gmail.com",
			Password: "p455word",
		},
	}

	var posts = []models.Post{
		models.Post{
			Title:   "Title 1",
			Content: "Hello world 1",
		},
		models.Post{
			Title:   "Title 2",
			Content: "Hello world 2",
		},
	}

	for i, _ := range users {
		err = server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table : %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = server.DB.Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table : %v", err)
		}
	}

	return users, posts, nil
}
