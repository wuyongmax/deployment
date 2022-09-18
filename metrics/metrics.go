package metrics

import (
        "fmt"
        "time"

        "github.com/prometheus/client_golang/prometheus"
)

func Register() {
        err := prometheus.Register(functionLatency)
        if err != nil {
                fmt.Println(err)
        }
}
//记录当前秒
func NewExectionTimer(histo *prometheus.HistogramVec) *ExecutionTimer {
        now := time.Now()
        return &ExecutionTimer{
                histo: histo,
                start: now,
                last:  now,
        }
}

const (
        MetricNamespace = "default"
)

func NewTimer() *ExecutionTimer {
        return NewExectionTimer(functionLatency)
}

var (
        functionLatency = CreateExecutionTimeMetric(MetricNamespace,
                "Time spent.")
)
//记录执行函数总时长
func (t *ExecutionTimer) ObserveTotal() {
        (*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}
//创建注册直方图
func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
        return prometheus.NewHistogramVec(
                prometheus.HistogramOpts{
                        Namespace: namespace,
                        Name:      "execution_latency_seconds",
                        Help:      help,
                        Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
                }, []string{"step"},
        )
        // fmt.Println("cc")
}

type ExecutionTimer struct {
        histo *prometheus.HistogramVec
        start time.Time
        last  time.Time
}

