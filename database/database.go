package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type News struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:255;not null"`
	Link      string `gorm:"type:text;not null"`
	Keyword   string `gorm:"size:100;not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ .env 파일 로드 실패")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// DB 생성용 루트 연결
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	rootDB, err := gorm.Open(mysql.Open(dsnRoot), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ MySQL 루트 연결 실패:", err)
	}

	rootDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))

	// 크롤러 DB 연결
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ DB 연결 실패:", err)
	}
	DB = db

	err = DB.AutoMigrate(&News{})
	if err != nil {
		log.Fatal("❌ 테이블 생성 실패:", err)
	}

	fmt.Println("✅ DB 연결 및 테이블 준비 완료")
}

// 크롤링한 데이터를 MySQL에 저장하는 함수
func SaveNews(title, link, keyword string) {
	news := News{Title: title, Link: link, Keyword: keyword}
	result := DB.Create(&news)

	if result.Error != nil {
		log.Println("❌ 데이터 저장 실패:", result.Error)
	}
	fmt.Println("📝 데이터 저장 완료:", title)
}
