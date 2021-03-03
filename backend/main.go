package main

import (
	"log"
	"time"

	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/envconfig"
	"github.com/short-d/app/fw/logger"
	"github.com/cross-team/clublink/backend/app"
	"github.com/cross-team/clublink/backend/cmd"
	"github.com/cross-team/clublink/backend/dep"
)

func main() {
	log.Println("starting")
	env := dep.InjectEnv()
	log.Println("auto load env")
	env.AutoLoadDotEnvFile()
	log.Println("env config")
	envConfig := envconfig.NewEnvConfig(env)

	config := struct {
		Runtime              string        `env:"ENV" default:"development"`
		DBHost               string        `env:"DB_HOST" default:"localhost"`
		DBPort               int           `env:"DB_PORT" default:"5432"`
		DBUser               string        `env:"DB_USER" default:"postgres"`
		DBPassword           string        `env:"DB_PASSWORD" default:"password"`
		DBName               string        `env:"DB_NAME" default:"short"`
		ReCaptchaSecret      string        `env:"RECAPTCHA_SECRET" default:""`
		GithubClientID       string        `env:"GITHUB_CLIENT_ID" default:""`
		GithubClientSecret   string        `env:"GITHUB_CLIENT_SECRET" default:""`
		FacebookClientID     string        `env:"FACEBOOK_CLIENT_ID" default:""`
		FacebookClientSecret string        `env:"FACEBOOK_CLIENT_SECRET" default:""`
		FacebookRedirectURI  string        `env:"FACEBOOK_REDIRECT_URI" default:""`
		GoogleClientID       string        `env:"GOOGLE_CLIENT_ID" default:""`
		GoogleClientSecret   string        `env:"GOOGLE_CLIENT_SECRET" default:""`
		GoogleRedirectURI    string        `env:"GOOGLE_REDIRECT_URI" default:""`
		JWTSecret            string        `env:"JWT_SECRET" default:""`
		WebFrontendURL       string        `env:"WEB_FRONTEND_URL" default:""`
		KeyGenBufferSize     int           `env:"KEY_GEN_BUFFER_SIZE" default:"50"`
		KgsHostname          string        `env:"KEY_GEN_HOSTNAME" default:"localhost"`
		KgsPort              int           `env:"KEY_GEN_PORT" default:"8080"`
		GraphQLAPIPort       int           `env:"GRAPHQL_API_PORT" default:"8080"`
		HTTPAPIPort          int           `env:"HTTP_API_PORT" default:"80"`
		GRPCAPIPort          int           `env:"GRPC_API_PORT" default:"8081"`
		EnableEncryption     bool          `env:"ENABLE_ENCRYPTION" default:"false"`
		CertFilePath         string        `env:"CERT_FILE_PATH" default:"/etc/certs/tls.crt"`
		KeyFilePath          string        `env:"KEY_FILE_PATH" default:"/etc/certs/tls.key"`
		AuthTokenLifeTime    time.Duration `env:"AUTH_TOKEN_LIFETIME" default:"1w"`
		SearchTimeout        time.Duration `env:"SEARCH_TIMEOUT" default:"1s"`
		SwaggerUIDir         string        `env:"SWAGGER_UI_DIR" default:"app/adapter/routing/public"`
		OpenAPISpecPath      string        `env:"OPEN_API_SPEC_PATH" default:"app/adapter/routing/api.yml"`
		GraphQLSchemaPath    string        `env:"GRAPHQL_SCHEMA_PATH" default:"app/adapter/gqlapi/schema.graphql"`
		GraphiQLDefaultQuery string        `env:"GRAPH_I_QL_DEFAULT_QUERY" default:""`
		DataDogAPIKey        string        `env:"DATA_DOG_API_KEY" default:""`
		SegmentAPIKey        string        `env:"SEGMENT_API_KEY" default:""`
		IPStackAPIKey        string        `env:"IP_STACK_API_KEY" default:""`
		GoogleAPIKey         string        `env:"GOOGLE_API_KEY" default:""`
	}{}

	log.Println(config)
	err := envConfig.ParseConfigFromEnv(&config)
	log.Println("error: ")
	log.Println(err)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	cmdFactory := dep.InjectCommandFactory()
	log.Println("connecting db")
	dbConnector := dep.InjectDBConnector()
	log.Println("migrating")
	dbMigrationTool := dep.InjectDBMigrationTool()

	dbConfig := db.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DbName:   config.DBName,
	}

	log.Println("dbConfig")
	log.Println(dbConfig)

	serviceConfig := app.ServiceConfig{
		Runtime:              config.Runtime,
		LogPrefix:            "Short",
		LogLevel:             logger.LogInfo,
		RecaptchaSecret:      config.ReCaptchaSecret,
		GithubClientID:       config.GithubClientID,
		GithubClientSecret:   config.GithubClientSecret,
		FacebookClientID:     config.FacebookClientID,
		FacebookClientSecret: config.FacebookClientSecret,
		FacebookRedirectURI:  config.FacebookRedirectURI,
		GoogleClientID:       config.GoogleClientID,
		GoogleClientSecret:   config.GoogleClientSecret,
		GoogleRedirectURI:    config.GoogleRedirectURI,
		JwtSecret:            config.JWTSecret,
		WebFrontendURL:       config.WebFrontendURL,
		GraphQLAPIPort:       config.GraphQLAPIPort,
		HTTPAPIPort:          config.HTTPAPIPort,
		GRPCAPIPort:          config.GRPCAPIPort,
		EnableEncryption:     config.EnableEncryption,
		CertFilePath:         config.CertFilePath,
		KeyFilePath:          config.KeyFilePath,
		KeyGenBufferSize:     config.KeyGenBufferSize,
		KgsHostname:          config.KgsHostname,
		KgsPort:              config.KgsPort,
		AuthTokenLifetime:    config.AuthTokenLifeTime,
		SearchTimeout:        config.SearchTimeout,
		SwaggerUIDir:         config.SwaggerUIDir,
		OpenAPISpecPath:      config.OpenAPISpecPath,
		GraphQLSchemaPath:    config.GraphQLSchemaPath,
		GraphiQLDefaultQuery: config.GraphiQLDefaultQuery,
		DataDogAPIKey:        config.DataDogAPIKey,
		SegmentAPIKey:        config.SegmentAPIKey,
		IPStackAPIKey:        config.IPStackAPIKey,
		GoogleAPIKey:         config.GoogleAPIKey,
	}

	rootCmd := cmd.NewRootCmd(
		dbConfig,
		serviceConfig,
		cmdFactory,
		dbConnector,
		dbMigrationTool,
	)
	cmd.Execute(rootCmd)
}
