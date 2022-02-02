package cmd

import (
	"fmt"
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository"
	"github.com/aririfani/auth-app/fetch-service/internal/app/service"
	"github.com/patrickmn/go-cache"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aririfani/auth-app/fetch-service/config"
	"github.com/aririfani/auth-app/fetch-service/internal/app/handler"
	"github.com/aririfani/auth-app/fetch-service/internal/app/server"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "auth app",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute executes the root command.
func Execute() (err error) {
	if err = rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}

	return
}

func start() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	cfg := config.NewConfig()
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	withCache := cache.New(5*time.Minute, 10*time.Minute)
	signal.Notify(quit, os.Interrupt)
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	srv := service.NewService(repository.NewRepo(cfg, httpClient, withCache), cfg)

	s := server.NewServer(
		net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")),
		handler.NewHandler(cfg, srv),
		time.Duration(cfg.GetInt("server.read_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.write_timeout"))*time.Second,
		time.Duration(cfg.GetInt("server.idle_timeout"))*time.Second,
	)

	httpServer := s.GetHTTPServer()
	go s.GracefullShutdown(httpServer, logger, quit, done)

	logger.Println("=> http server started on", net.JoinHostPort(cfg.GetString("server.host"), cfg.GetString("server.port")))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", cfg.GetString("server.port"), err)
	}

	<-done

	logger.Println("Server stopped")
}
