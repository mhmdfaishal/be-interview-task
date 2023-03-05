package server

import (
	"be-interview-task/internal/boot"
	"be-interview-task/internal/utils/database/manager"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	app            *fiber.App
	db             *gorm.DB
	manger         *manager.Manager
	contextTimeout time.Duration
}

func NewServer() *Server {
	contextTimeout := viper.GetInt("CONTEXT_TIMEOUT")
	return &Server{
		app:            fiber.New(),
		db:             boot.MainDBConn,
		manger:         manager.NewManager(boot.MainDBConn),
		contextTimeout: time.Duration(contextTimeout) * time.Second,
	}
}

func (s *Server) Start() {

	if err := s.MapHandlers(s.app); err != nil {
		log.Panic(err)
	}
	server_port := viper.GetString("SERVER_PORT")
	go func() {
		fmt.Printf("Server running on port: %s\n", server_port)
		if err := s.app.Listen(fmt.Sprintf(":%s", server_port)); err != nil {
			log.Panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("shutting down...")
	_ = s.app.Shutdown()
}
