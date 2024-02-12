package seeder

import (
	articleModel "blog/internal/modules/article/models"
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPass, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("Hash password error")
	}
	user := userModel.User{Name: "Hossein", Email: "info@hossein.cloud", Password: string(hashedPass)}
	db.Create(&user)
	log.Printf("User created successfully with email %s \n", user.Email)

	for i := 0; i < 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Title %d", i), Content: fmt.Sprintf("Content %d", i), UserID: 1}
		db.Create(&article)

		log.Printf("Article created successfully with title %s \n", article.Title)
	}

	log.Println("Seeder Done.")
}
