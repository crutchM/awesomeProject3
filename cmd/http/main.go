package main

import (
	"awesomeProject3/internal/common/config"
	"awesomeProject3/internal/common/logger"
	"awesomeProject3/internal/repository"
	"awesomeProject3/pkg/logrus"
	"awesomeProject3/pkg/pg"
)

func main() {
	config := config.NewConfig(".")

	logrus := logrus.NewLogger()

	logger := logger.New(logrus)

	db := pg.NewDatabase(config.PgURL)

	rm := repository.NewRepoManager(db)
}
