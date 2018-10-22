package metrics

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	updateMetricsFunctions = make(map[string]func(context.Context, string))
)

// RegisterUpdateMetricsFunctions allows you to register a function
// that will update prometheus metrics
func RegisterUpdateMetricsFunctions(name string, f func(context.Context, string)) {
	updateMetricsFunctions[name] = f
}

// UpdateMetrics
func UpdateMetrics(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	t := time.Now()

	// Process which updates metrics
	for {
		log.Debugf("Start metrics update: %s", t)

		// Loop over all update functions metrics
		for updateMetricsFuncName, updateMetricsFunc := range updateMetricsFunctions {
			// We detach the update process so if it takes more than the refresh
			// time it does not get blocked
			go func(ctx context.Context, updateMetricsFuncName string, updateMetricsFunc func(context.Context, string), t time.Time) {
				id := hashTime(t)
				fields := log.Fields{
					"_id":  id,
					"name": updateMetricsFuncName,
				}

				log.WithFields(fields).Debug("Start update metrics function")
				updateMetricsFunc(ctx, id)
				log.WithFields(fields).Debug("End update metrics function")
			}(ctx, updateMetricsFuncName, updateMetricsFunc, t)
		}

		t = <-ticker.C
	}
}

func hashTime(t time.Time) string {
	h := md5.New()
	io.WriteString(h, t.String())
	s := fmt.Sprintf("%x", h.Sum(nil))

	return s[0:16]
}
