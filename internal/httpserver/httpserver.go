package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/alifnh/bjb-auction-backend/internal/handler/httphandler"
	"github.com/alifnh/bjb-auction-backend/internal/httpserver/middleware"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/database"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/encryptutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/jwtutils"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/logger"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/randutils"
	"github.com/alifnh/bjb-auction-backend/internal/repository"
	"github.com/alifnh/bjb-auction-backend/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func initServer(cfg *config.Config) *http.Server {
	db, err := database.InitDB(cfg)
	if err != nil {
		logger.Log.Fatal("error initializing database: ", err.Error())
	}

	// database
	postgresWrapper := database.NewPostgresWrapper(db)
	transactor := database.NewTransactor(db)

	// utils
	jwtUtil := jwtutils.NewJwtUtil(cfg.Jwt)
	passwordEncryptor := encryptutils.NewBcryptPasswordEncryptor(cfg.App.BCryptCost)
	randutil := randutils.NewStdLibRandomUtil()

	// repositories
	authRepository := repository.NewAuthRepository(postgresWrapper)

	// usecases
	authUsecase := usecase.NewAuthUsecase(authRepository, transactor, passwordEncryptor, jwtUtil, cfg, randutil)

	// handlers
	appHandler := httphandler.NewAppHandler()
	authHandler := httphandler.NewAuthHandler(authUsecase)

	// to remove the Gin's warning
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.ContextWithFallback = true // enables .Done(), .Err(), and .Value()

	registerValidators()

	// init middlewares
	// authMiddleware := middleware.NewAuthMiddleware(jwtUtil)
	// authzMiddleware := middleware.NewAuthorizationMiddleware()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	// registering middlewares
	middlewares := []gin.HandlerFunc{
		middleware.ErrorHandler(),
		middleware.Logger(),
		gin.Recovery(),
		cors.New(corsConfig),
	}
	r.Use(middlewares...)

	// registering routes
	r.NoRoute(appHandler.RouteNotFound)
	r.GET("/", appHandler.Index)

	fmt.Println("in")
	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	return srv
}

func StartGinHttpServer(cfg *config.Config) {
	srv := initServer(cfg)

	// graceful shutdown
	go func() {
		logger.Log.Info("running server on port :", cfg.HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatal("error while server listen and serve: ", err)
			}
		}
		logger.Log.Info("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(cfg.HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	logger.Log.Info("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("error shutting down server: ", err)
	}

	logger.Log.Info("http server is shutting down gracefully")
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
			if valuer, ok := field.Interface().(decimal.Decimal); ok {
				return valuer.String()
			}
			return nil
		}, decimal.Decimal{})

		v.RegisterValidation("dgte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.GreaterThanOrEqual(baseValue)
		})

		v.RegisterValidation("dlte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.LessThanOrEqual(baseValue)
		})
	}
}
