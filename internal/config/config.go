package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type CacheConfig struct {
	LifeWindow  time.Duration `envconfig:"CACHE_LIFEWINDOW" default:"10m"`
	CleanWindow time.Duration `envconfig:"CACHE_CLEANWINDOW" default:"15m"`

	//Max useable ram in MB
	HardMaxCacheSize int `envconfig:"CACHE_MAX_SIZE" default:"512"`
}

type GRPCConfig struct {
	DataCollectionGRPCHost string `envconfig:"DATA_COLLECTION_GRPC_HOST" required:"true"`
	MapGRPCHost            string `envconfig:"MAP_GRPC_HOST" required:"true"`
	UserTrackingGRPCHost   string `envconfig:"USER_TRACKING_GRPC_HOST" required:"true"`
	UserManagerGRPCHost    string `envconfig:"USER_MANAGER_GRPC_HOST" requried:"true"`
}

type MinioConfig struct {
	Bucket string `envconfig:"MINIO_BUCKET" default:"document-service"`
}

func LoadConfig() {
	if currentEnvironment, ok := os.LookupEnv("ENV"); ok {
		_, b, _, _ := runtime.Caller(0)
		basePath := filepath.Dir(b)
		envPath := fmt.Sprintf("%v/../../.env.%v", basePath, currentEnvironment)
		fmt.Println(envPath)
		err := godotenv.Load(envPath)
		if err != nil {
			logrus.Errorf("error load file .env.%s: %v", currentEnvironment, err.Error())
			if err := godotenv.Load(); err != nil {
				panic(err)
			}
		}
		fmt.Printf("using %s env, configuration from file path : %s \n", currentEnvironment, envPath)
	}
}
