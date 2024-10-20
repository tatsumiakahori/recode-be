package main

import (
	// "fmt"
	// "log"
	// "os"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func main() {
	// 環境変数から接続情報を取得して、DSNを構築
	// dbHost := os.Getenv("DB_HOST")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName)

	// GORMでデータベースに接続
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database: ", err)
	// }

	// db変数を使ってテーブルを作成
	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate database: ", err)
	// }

	// Ginエンジンのインスタンスを作成
	router := gin.Default()

	// とりあえずのHello World!
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// usersを取得するエンドポイント
	// router.GET("/users", func(c *gin.Context) {
	// 	var users []User
	// 	db.Find(&users)
	// 	c.JSON(200, users)
	// })

	// 8080ポートでサーバーを起動
	router.Run(":8080")
}
