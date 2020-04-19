package config

import "fmt"

type Config struct {
	DbUser        string `short:"u" long:"db-user" description:"The mysql db user" default:"root" env:"DB_USER"`
	DbPassword    string `short:"p" long:"db-password" description:"The mysql db password" default:"root" env:"DB_PWD"`
	DbName        string `short:"n" long:"db-name" description:"The mysql db name" default:"k8sbasicdemo" env:"DB_NAME"`
	DbHost        string `short:"d" long:"db-host" description:"The mysql db host" default:"localhost" env:"DB_HOST"`
	MigrationOnly bool   `short:"m" long:"migration-only" description:"run only the db migration" required:"false" env:"MIGRATION_ONLY"`

	CUnknown string `short:"c" long:"db-c" description:"without this parameter deploying with skaffold crashes with unknown flag c"`
}

func (c Config) SqlConnString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/", c.DbUser, c.DbPassword, c.DbHost)
}

func (c Config) GormConnString() string {
	return fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", c.DbUser, c.DbPassword, c.DbHost, c.DbName)
}
