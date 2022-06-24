package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/go-redis/redis/v8"
	"github.com/kitabisa/golang-poc/config"
	"github.com/kitabisa/golang-poc/internal/app/appcontext"
	"github.com/kitabisa/golang-poc/internal/app/commons"
	"github.com/kitabisa/golang-poc/internal/app/driver"
	"github.com/kitabisa/golang-poc/internal/app/metrics"
	"github.com/kitabisa/golang-poc/internal/app/repository"
	"github.com/kitabisa/golang-poc/internal/app/server"
	"github.com/kitabisa/golang-poc/internal/app/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v3"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-poc",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.Config()
	app := appcontext.NewAppContext(cfg)

	config.ConfigureLogLevel(app.GetAppOption().LogLevel)

	var err error
	var dbPostgre *gorp.DbMap
	if app.GetPostgreOption().IsEnable {
		dbPostgre, err = app.GetDBInstance(appcontext.DBDialectPostgres)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Postgre | %v", err)
			return
		}
	}

	var cacheClient *redis.Client
	if app.GetCacheOption().IsEnable {
		cacheClient = driver.NewCache(app.GetCacheOption())
		defer cacheClient.Close()
	}

	opt := commons.Options{
		AppCtx:      app,
		CacheClient: cacheClient,
		Config:      cfg,
		DbPostgre:   dbPostgre,
		Metric:      metrics.NewMetric(app.GetTelegrafOption(), app.GetAppOption().Name),
	}

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	service := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption)

	repo := repository.Repository{
		Cache: cacheRepo,
	}
	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)

	svc := service.Services{
		HealthCheck: hc,
	}
	return &svc
}
