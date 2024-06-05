package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tsileo/defender"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	baseLog "log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
	"tinderleaked/modules/auth/authDelivery"
	"tinderleaked/modules/auth/authRepository"
	"tinderleaked/modules/auth/authUsecase"
	"tinderleaked/modules/packages/packagesDelivery"
	"tinderleaked/modules/packages/packagesRepository"
	"tinderleaked/modules/packages/packagesUsecase"
	"tinderleaked/modules/users/usersDelivery"
	"tinderleaked/modules/users/usersRepository"
	"tinderleaked/modules/users/usersUsecase"
)

// Header
const (
	HeaderSort          = "X-SORT"
	HeaderSortDirection = "X-SORT-DIRECTION"
	HeaderLimit         = "X-LIMIT"
	HeaderOffset        = "X-OFFSET"
	HeaderKeyword       = "X-KEYWORD"
)

var rxURL = regexp.MustCompile(`^/regexp\d*`)

var (
	dbHost, dbPort, dbUser, dbPass, dbName string
	d                                      *defender.Defender
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Error().Msg("failed read configuration file !!!!!!")
		return
	}
}

func main() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")

	val := url.Values{}
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s&%s", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName), val.Encode())

	newLogger := gormLogger.New(
		baseLog.New(os.Stdout, "\r\n", baseLog.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold: time.Second,     // Slow SQL threshold
			LogLevel:      gormLogger.Info, // Log level
			Colorful:      true,            // Disable color
		},
	)

	db, errGorm := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if errGorm != nil {
		log.Error().Msg("Failed Connect to database using gorm =" + errGorm.Error())
		return
	}

	//init defender brute force login
	d = defender.New(5, 5*time.Minute, 10*time.Minute)

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		log.Error().Msg("port cant empty")
		return
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "userid", "leadid",
			HeaderKeyword, HeaderOffset, HeaderLimit, HeaderSort, HeaderSortDirection},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           720 * time.Hour,
	}))

	r.Use(logger.SetLogger(
		logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		}),
		logger.WithUTC(true),
		logger.WithSkipPath([]string{"/skip"}),
		logger.WithSkipPathRegexps(rxURL),
	))

	r.Use(gin.Recovery())

	go func() {
		//auth
		authRepository := authRepository.NewAuthRepository(db)
		authUsecase := authUsecase.NewAuthUsecase(authRepository)
		authDelivery.NewAuthHTTPHandler(r, d, authUsecase)

		//packages
		packagesRepository := packagesRepository.NewPackagesRepository(db)
		packagesUsecase := packagesUsecase.NewPackagesUsecase(packagesRepository)
		packagesDelivery.NewPackagesHTTPHandler(r, packagesUsecase)

		//users
		usersRepository := usersRepository.NewUsersRepository(db)
		usersUsecase := usersUsecase.NewUserUsecase(usersRepository, packagesRepository)
		usersDelivery.NewUsersHTTPHandler(r, usersUsecase)

	}()

	log.Info().Msg("Service Running version 1.0.0 at port : " + port + " last update " + time.Now().Format(time.RFC3339))

	if errHTTP := http.ListenAndServe(":"+port, r); errHTTP != nil {
		log.Error().Msg(errHTTP.Error())
	}
}
