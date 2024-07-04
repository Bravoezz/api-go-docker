package config

import "fmt"

type PgSqlConfig struct {
	User     string
	DbName   string
	Password string
	Host     string
}

func (pg PgSqlConfig) Format() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=verify-full password=%s host=%s",
		pg.User,
		pg.DbName,
		pg.Password,
		pg.Host,
	)
}
