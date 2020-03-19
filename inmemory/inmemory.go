package inmemory

import (
	"fmt"
	"time"

	"github.com/mgw2007/golang-metric/metric"
)

// inMemory Impelemt manage metric save using in memory
type inMemory struct {
	metrics map[string]*metric.Metric
}

// AddMetric for add
func (im *inMemory) AddMetric(key string, t time.Time) {
	if targetMetric, ok := im.metrics[key]; ok {
		// validate time
		duration := t.Sub(targetMetric.UpdateTime).Hours()
		if duration > 1 {
			targetMetric.Count = 1
		} else {
			targetMetric.Count++
		}
	} else {
		im.metrics[key] = &metric.Metric{
			Count:      1,
			UpdateTime: t,
		}
	}
}

//GetMetricCount for
func (im *inMemory) GetMetricCount(key string, t time.Time) (int, error) {
	if targetMetric, ok := im.metrics[key]; ok {
		// check if time ended
		duration := t.Sub(targetMetric.UpdateTime).Hours()
		if duration > 1 {
			targetMetric.Count = 0
		}
		return targetMetric.Count, nil
	}
	return 0, fmt.Errorf("Metic not exist")
}

// NewMetric for
func NewMetric() metric.HandleMetrics {
	return &inMemory{
		metrics: map[string]*metric.Metric{},
	}
}
