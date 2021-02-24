package seed

import (
	"log"

	"github.com/iamsyahidi/forum-backend-golang/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
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

var posts = []models.Post{
	models.Post{
		Title:   "Title 3",
		Content: "Hello world 3",
	},
	models.Post{
		Title:   "Title 4",
		Content: "Hello world 4",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
