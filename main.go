package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version   = "dev"
	commit    = "none"
	buildTime = "unknown"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}

	InitConfig()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	e := echo.New()

	target, _ := url.Parse(fmt.Sprintf("http://localhost:%d", Config.ProxyPort))
	proxy := httputil.NewSingleHostReverseProxy(target)

	e.GET("/*", echo.WrapHandler(proxy))

	// --- Middleware
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogError:         true,
		LogHost:          true,
		LogLatency:       true,
		LogMethod:        true,
		LogContentLength: true,
		LogProtocol:      true,
		LogReferer:       true,
		LogUserAgent:     true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogResponseSize:  true,
		LogURIPath:       true,
		LogRoutePath:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ext := hasExtension(v.URI)
			route := ""

			if !ext {
				route = v.URI
			}

			if v.Error != nil {
				log.Error().
					Err(v.Error).
					Str("URI", v.URI).
					Int("status", v.Status).
					Str("method", v.Method).
					Str("remote_ip", v.RemoteIP).
					Str("host", v.Host).
					Str("uri", v.URI).
					Str("protocol", v.Protocol).
					Str("referer", v.Referer).
					Str("user_agent", v.UserAgent).
					Str("id", v.RequestID).
					Int("latency", int(v.Latency.Nanoseconds())).
					Str("latency_human", v.Latency.String()).
					Int("bytes_in", int(c.Request().ContentLength)).
					Int("bytes_out", int(v.ResponseSize)).
					Str("route", route).
					Msg("error")
			} else {
				log.Info().
					Str("URI", v.URI).
					Int("status", v.Status).
					Str("method", v.Method).
					Str("remote_ip", v.RemoteIP).
					Str("host", v.Host).
					Str("uri", v.URI).
					Str("protocol", v.Protocol).
					Str("referer", v.Referer).
					Str("user_agent", v.UserAgent).
					Str("id", v.RequestID).
					Int("latency", int(v.Latency.Nanoseconds())).
					Str("latency_human", v.Latency.String()).
					Int("bytes_in", int(c.Request().ContentLength)).
					Int("bytes_out", int(v.ResponseSize)).
					Str("route", route).
					Msg("request")
			}

			return nil
		},
	}))

	log.Info().Str("version", version).Str("commit", commit).Str("date", buildTime).Msg("")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", Config.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msg("error starting server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("error shutting down server")
	} else {
		log.Info().Msg("shutting down server")
	}
}
