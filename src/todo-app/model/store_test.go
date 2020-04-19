package model

import (
	"testing"

	"github.com/crixo/k8s-basic-demo/config"
)

func TestInMemoryRepository(t *testing.T) {
	sut := NewInMemoryRepositry()
	email := "test@acme.org"
	msg := Message{Email: email}
	sut.Store(msg)
	msgs := sut.GetList()

	var got, want int

	got = len(msgs)
	want = 1
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	msgs[0].Email = "test2@acme.org"
	msgs = sut.GetList()
	msgs[0].Email = email
}

func TestMysqlRepository(t *testing.T) {

	config := config.Config{
		DbUser:     "root",
		DbPassword: "root",
		DbName:     "k8sbasicdemo",
		DbHost:     "localhost",
	}

	CreateMysqlDbIfNotExists(config.SqlConnString(), config.DbName)

	db := CreateDb(config.GormConnString())
	defer db.Close()
	RunMigration(db)

	sut := NewMysqlRepository(db)
	email := "test@acme.org"
	msg := &Message{Email: email}
	currentMsgs := sut.GetList()
	sut.Store(msg)
	msgs := sut.GetList()

	var got, want int
	got = len(msgs)
	want = len(currentMsgs) + 1
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
