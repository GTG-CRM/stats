package stats

import (
	"fmt"
	"strings"

	"github.com/VictoriaMetrics/metrics"
)

var _app = "gtg"

func WithAppName(app string) {
	_app = app
}

func Incr(subsystem, name string, tags ...string) {
	c := getCounter(subsystem, name, tags...)
	c.Add(1)
}

func Set(subsystem, name string, value uint64, tags ...string) {
	c := getCounter(subsystem, name, tags...)
	c.Set(value)
}

func GaugeSet(subsystem, name string, value float64, tags ...string) {
	g := getGauge(subsystem, name, tags...)
	g.Set(value)
}

func GaugeAdd(subsystem, name string, value float64, tags ...string) {
	g := getGauge(subsystem, name, tags...)
	g.Add(value)
}

func GaugeGet(subsystem, name string, tags ...string) float64 {
	g := getGauge(subsystem, name, tags...)
	return g.Get()
}

func Histogram(subsystem, name string, value float64, tags ...string) {
	h := getHistogram(subsystem, name, tags...)
	h.Update(value)
}

func getCounter(subsystem, name string, tags ...string) *metrics.Counter {
	s := metrics.GetDefaultSet()
	return s.GetOrCreateCounter(nm(subsystem, name, tags...))
}

func getGauge(subsystem, name string, tags ...string) *metrics.Gauge {
	s := metrics.GetDefaultSet()
	return s.GetOrCreateGauge(nm(subsystem, name, tags...), nil)
}

func getHistogram(subsystem, name string, tags ...string) *metrics.Histogram {
	s := metrics.GetDefaultSet()
	return s.GetOrCreateHistogram(nm(subsystem, name, tags...))
}

func nm(subsystem, name string, tags ...string) string {
	if len(tags) == 0 {
		return fmt.Sprintf("%s_%s_%s", _app, subsystem, name)
	}

	t := strings.Join(tags, ",")
	return fmt.Sprintf("%s_%s_%s{%s}", _app, subsystem, name, t)
}
