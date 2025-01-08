package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"face-be-service/common/constants"
	"face-be-service/common/third_party"
	"face-be-service/handler"
	"face-be-service/handler/middlewares"
	"face-be-service/internal/domains/employee"
	"face-be-service/internal/infrastructure/repository"
)

func initDatabase() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(viper.GetString("DB_SOURCE")), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("error getting DB instance: %v", err)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(60 * 60)

	err = sqlDB.Ping()
	if err != nil {
		logrus.Errorf("error pinging database: %v", err)
		return nil, err
	}

	logrus.Info("Successfully connected to the database")

	return db, nil
}

func initMinio(uri, accessKey, secretKey, region string) (*minio.Client, error) {
	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}

	client, err := minio.New(
		uri,
		&minio.Options{
			Creds:     credentials.NewStaticV4(accessKey, secretKey, ""),
			Region:    region,
			Transport: transport,
			Secure:    false,
		})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	viper.AddConfigPath("build")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	db, err := initDatabase()
	if err != nil {
		logrus.Fatal("failed to open database:", err)
		return
	}

	// minio
	minioClient, err := initMinio(viper.GetString(constants.MinioURI), viper.GetString(constants.MinioAccessKey),
		viper.GetString(constants.MinioSecretKey), viper.GetString(constants.MinioRegion))
	if err != nil {
		logrus.Fatal("failed to open minio:", err)
		return
	}

	// repository
	minioRepo := repository.NewMinioRepository(minioClient, 30*time.Second)
	employeeRepo := repository.NewEmployeeRepository(db)

	err = minioRepo.InitBucket(viper.GetString(constants.MinioBucket))
	if err != nil {
		logrus.Errorf(constants.FormatTaskErr, "minioRepo.InitBucket", err)
		return
	}

	// init third party
	third_party.NewFaceExtractionClient()

	// service
	employeeService := employee.NewEmployeeService(employeeRepo, minioRepo)

	h := handler.NewHandler(employeeService)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middlewares.Recover())

	r.Static("/static", "./static")

	r.GET("/checkin", func(c *gin.Context) {
		c.File("./static/checkin.html")
	})

	r.GET("/enroll", func(c *gin.Context) {
		c.File("./static/enroll.html")
	})

	// employee APIs
	employeeAPI := r.Group("api/v1/employees")
	{
		// auth
		employeeAPI.Use(middlewares.APIKeyAuthentication())

		employeeAPI.POST("insert", h.Insert)
		employeeAPI.POST("search", h.Search)
	}

	err = r.RunTLS(":"+viper.GetString("PORT"), "server.crt", "server.key")
	if err != nil {
		return
	}
}
