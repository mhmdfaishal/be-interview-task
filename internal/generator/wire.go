//go:build wireinject
// +build wireinject

package generator

import (
	"be-interview-task/internal/contracts"
	"be-interview-task/internal/utils/database/manager"
	"time"
	
	"github.com/google/wire"
	_yaraUsecase "be-interview-task/internal/modules/yara/usecase"
	_yaraRepository "be-interview-task/internal/modules/yara/repository"
)

var YaraUsecaseSet = wire.NewSet(
	_yaraRepository.NewYaraRepository,
	_yaraUsecase.NewYaraUsecase,
)

func InitializeYaraUsecase(timeout time.Duration, conn manager.DBTx) contracts.YaraUsecase {
	panic(wire.Build(YaraUsecaseSet))
}
