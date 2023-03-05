package server

import (
	"github.com/spf13/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_yaraHandler "be-interview-task/internal/modules/yara/delivery"
	"be-interview-task/internal/generator"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func (s *Server) MapHandlers(f *fiber.App) error {

	// logger middleware
	s.app.Use(logger.New())

	// cors middleware
	s.app.Use(cors.New())

	// Swagger UI
	s.app.Static("/swagger", "./docs/swagger-ui")

	s.app.Static("/swagger-docs", "./docs/swagger")

	_yaraHandler.NewYaraHandler(f, generator.InitializeYaraUsecase(s.contextTimeout, s.db))

	return nil
}
