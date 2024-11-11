package delivery

import (
	"log"
	"test-ordent/config"
	"test-ordent/delivery/controller"
	"test-ordent/middleware"
	"test-ordent/repository"
	"test-ordent/usecase"
	"test-ordent/utils/common"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authMiddlware middleware.AuthMiddleware
	userUc        usecase.UserUsecase
	bookUc        usecase.BookUsecase
	engine        *gin.Engine
	host          string
}

func (s *Server) setupControllers() {
	controller.NewUserController(s.userUc, s.engine).Route()
	controller.NewBookController(s.bookUc, s.engine, s.authMiddlware).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("Server Error : ", err.Error())
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal("Config : ", err.Error())
	}

	db, err := config.NewDbConnection(cfg)

	if err != nil {
		log.Fatal("DB Connect : ", err.Error())
	}

	// Token Service
	jwtToken := common.NewJwtToken(cfg.TokenConfig)

	//Middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtToken)

	// User
	userRepo := repository.NewUserRepository(db.Conn())
	bookRepo := repository.NewBookRepository(db.Conn())
	userUc := usecase.NewUserUsecase(userRepo, jwtToken)
	bookUc := usecase.NewBookUsecase(bookRepo)

	// Gin Engine
	engine := gin.Default()

	return &Server{
		authMiddlware: authMiddleware,
		userUc:        userUc,
		bookUc:        bookUc,
		engine:        engine,
		host:          ":8085",
	}
}
