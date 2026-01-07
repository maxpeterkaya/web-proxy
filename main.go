package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	npmVersion   = "npm"
	startCommand = ""
	staticFolder = "dist"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg("Error loading .env file")
	}

	InitConfig()
}

func main() {
	// Set log time to be in unix
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	// --- web-proxy CLI flags
	var shouldRunApp bool
	flag.BoolVar(&shouldRunApp, "app", false, "Run web-app through web-proxy")

	var shouldLogApp bool
	flag.BoolVar(&shouldLogApp, "log", true, "Log web-app")

	var staticWeb bool
	flag.BoolVar(&staticWeb, "static", false, "Use static website build")
	flag.StringVar(&staticFolder, "static-dir", "dist", "Folder containing static build")

	flag.Parse()

	// Initialize echo
	e := echo.New()

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
				uri, _ := url.Parse(v.URI)
				uri.RawQuery = ""
				route = uri.String()
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

	var CMD *exec.Cmd

	if shouldRunApp {
		ExtractPackage()

		CMD = exec.Command(strings.Split(startCommand, " ")[0], strings.Split(startCommand, " ")[1:]...)

		if shouldLogApp {
			CMD.Stdout = os.Stdout
		}

		CMD.Stderr = os.Stderr

		err := CMD.Start()
		if err != nil {
			log.Error().Err(err).Msg("Error starting web-app")
			return
		}
	}

	if staticWeb {
		e.Static("/", staticFolder)
	} else {
		target, _ := url.Parse(fmt.Sprintf("http://localhost:%d", Config.ProxyPort))

		e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: target,
			},
		})))
	}

	NpmVersionExtractor()

	log.Info().Str("proxy-version", version).Str("proxy-commit", commit).Str("proxy-date", date).Msg("proxy log")
	log.Info().Str("npm-version", npmVersion).Msg("npm log")

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

		if shouldRunApp {
			err = CMD.Process.Signal(os.Interrupt)
			if err != nil {
				log.Error().Err(err).Msg("error interrupting web-app")

				err = CMD.Process.Signal(os.Kill)
				if err != nil {
					log.Fatal().Err(err).Msg("error killing web-app")
				}
			}
		}
	}
}
