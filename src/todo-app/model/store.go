package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Repository interface {
	Store(message *Message) error
	GetList() []Message
}

func NewInMemoryRepositry() *InMemoryRepository {
	return &InMemoryRepository{
		messages: []Message{},
	}
}

type InMemoryRepository struct {
	messages []Message
}

func (im *InMemoryRepository) Store(message Message) error {
	im.messages = append(im.messages, message)
	return nil
}

func (im InMemoryRepository) GetList() []Message {
	return im.messages
}

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (mr MysqlRepository) Store(message *Message) error {
	mr.db.Create(message)
	return nil
}

func (mr MysqlRepository) GetList() []Message {
	msgs := []Message{}
	mr.db.Find(&msgs)
	return msgs
}

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(&Message{})
}

func CreateMysqlDbIfNotExists(connString, dbName string) {
	db, err := sql.Open("mysql", connString) //"admin:admin@tcp(127.0.0.1:3306)/"
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}
}

func CreateDb(connString string) *gorm.DB {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	return db
}
