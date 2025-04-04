package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type News struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:255;not null"`
	Link      string    `gorm:"type:text;not null,uniqueIndex"`
	Keyword   string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Gorm이 자동으로 시간 설정
}

var DB *gorm.DB

func ConnectDB() {
	//  환경 변수를 보고 개발환경일 때만 .env 로드
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println(" .env 파일을 찾을 수 없습니다. 무시하고 계속 진행합니다.")
		} else {
			log.Println(" .env 파일 로드 성공")
		}
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
		log.Fatal(" MySQL 루트 연결 실패:", err)
	}

	rootDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))

	// 크롤러 DB 연결
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(" DB 연결 실패:", err)
	}
	DB = db

	err = DB.AutoMigrate(&News{})
	if err != nil {
		log.Fatal(" 테이블 생성 실패:", err)
	}

	fmt.Println(" DB 연결 및 테이블 준비 완료")
}

// 크롤링한 데이터를 MySQL에 저장하는 함수
func SaveNews(title, link, keyword string) {
	news := News{Link: link}
	result := DB.Where("link = ?", link).First(&news)

	if result.RowsAffected == 0 {
		news.Title = title
		news.Keyword = keyword
		if err := DB.Create(&news).Error; err != nil {
			log.Println(" 데이터 저장 실패:", err)
		} else {
			fmt.Println(" 데이터 저장 완료:", title)
		}
	} else {
		fmt.Println(" 이미 존재하는 뉴스입니다. 중복 저장하지 않음:", title)
	}
}
