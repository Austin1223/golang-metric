package metric

import "time"

// Metric for
type Metric struct {
	Count      int
	UpdateTime time.Time
}

// HandleMetrics for
type HandleMetrics interface {
	AddMetric(key string, t time.Time)
	GetMetricCount(key string, t time.Time) (int, error)
}

// PostMetricResponse for
type PostMetricResponse struct {
}

// GetMetricCountResponse for
type GetMetricCountResponse struct {
	Value int
}
