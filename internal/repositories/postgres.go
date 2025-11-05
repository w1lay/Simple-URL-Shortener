package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
)

func NewPostgresqlConnection() *sqlx.DB {
	db, err := sqlx.Open("postgres", settings.Settings.Postgres.URI())
	if err != nil {
		logrus.Fatalf("Can`t connect to DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatalf("Can`t connect to DB: %v", err)
	}

	var dbName string

	err = db.Get(&dbName, "SELECT current_database();")
	if err != nil {
		logrus.Fatalf("Can`t connect to PostgreSQL DB: %v", err)
	}

	logrus.Infof("Connected to PostgreSQL database: %s", dbName)

	return db
}
