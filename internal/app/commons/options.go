package commons

import (
	"github.com/go-redis/redis/v8"
	"github.com/kitabisa/golang-poc/config"
	"github.com/kitabisa/golang-poc/internal/app/appcontext"
	"github.com/kitabisa/golang-poc/internal/app/metrics"
	"gopkg.in/gorp.v3"
)

// Options common option for all object that needed
type Options struct {
	AppCtx      *appcontext.AppContext
	Config      config.Provider
	CacheClient *redis.Client
	Metric      metrics.IMetric
	DbPostgre   *gorp.DbMap
	
}
