package metrics

import (
	"fmt"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/rs/zerolog/log"
)

// TelegrafOption defines options for telegraf
type TelegrafOption struct {
	IsEnabled bool
	Host      string
	Port      int
}

// IMetric ...
type IMetric interface {
	IncrErr(tag []string)
	IncrSuccess(tag []string)

	IncrErrRedis()
	IncrErrDB()
}

// Metric ...
type Metric struct {
	st       *statsd.Client
	isEnable bool
}

// table metrics
var (
	TableErr     = "ERROR"
	TableSuccess = "SUCCESS"
)

// tags
var (
	TErrRedis   = []string{"type:redis"}
	TErrDB      = []string{"type:db"}
	TErrService = []string{"type:service"}
)

// NewMetric initializes telegraf instance
func NewMetric(option TelegrafOption, namespace string) *Metric {
	telegrafMetrics := &Metric{
		isEnable: option.IsEnabled,
	}

	if option.IsEnabled {
		telegrafURL := fmt.Sprintf("%s:%d", option.Host, option.Port)
		c, err := statsd.New(telegrafURL)
		if err != nil {
			panic(err)
		}

		if namespace != "" {
			c.Namespace = fmt.Sprintf("%s_", namespace)
		}
		telegrafMetrics.st = c
	}

	return telegrafMetrics
}

// IncrErr general error metric
func (m *Metric) IncrErr(tag []string) {
	if !m.isEnable {
		return
	}

	err := m.st.Incr(TableErr, tag, 1)
	if err != nil {
		log.Err(err).Send()
	}
}

// IncrSuccess general success metric
func (m *Metric) IncrSuccess(tag []string) {
	if !m.isEnable {
		return
	}

	err := m.st.Incr(TableSuccess, tag, 1)
	if err != nil {
		log.Err(err).Send()
	}
}

// IncrErrRedis for metric error redis
func (m *Metric) IncrErrRedis() {
	if !m.isEnable {
		return
	}

	err := m.st.Incr(TableErr, TErrRedis, 1)
	if err != nil {
		log.Err(err).Send()
	}
}

// IncrErrDB for metric error database
func (m *Metric) IncrErrDB() {
	if !m.isEnable {
		return
	}

	err := m.st.Incr(TableErr, TErrDB, 1)
	if err != nil {
		log.Err(err).Send()
	}
}
