// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package generator

import (
	"be-interview-task/internal/contracts"
	"be-interview-task/internal/modules/yara/repository"
	"be-interview-task/internal/modules/yara/usecase"
	"be-interview-task/internal/utils/database/manager"
	"github.com/google/wire"
	"time"
)

// Injectors from wire.go:

func InitializeYaraUsecase(timeout time.Duration, conn manager.DBTx) contracts.YaraUsecase {
	yaraRepository := repository.NewYaraRepository(conn)
	yaraUsecase := usecase.NewYaraUsecase(timeout, yaraRepository)
	return yaraUsecase
}

// wire.go:

var YaraUsecaseSet = wire.NewSet(repository.NewYaraRepository, usecase.NewYaraUsecase)
