package app

import (
	"VacancyService/internal/config"
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	_ "net/http/pprof"
)

func newPrometheusService(log zerolog.Logger, conf *config.Config) {
	mux := http.NewServeMux()
	log.Info().Msg("METRICS: Starting metrics server")

	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

	mux.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf(":%d", conf.PrometheusConf.PromPort)

	go func() {
		err := http.ListenAndServe(addr, mux)
		if err != nil {
			log.Fatal().Err(err).Msg("METRICS: Failed to start metrics server")
		}
	}()
}
