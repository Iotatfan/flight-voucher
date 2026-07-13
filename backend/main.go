package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"iotatfan.com/airline-voucher/config"
	"iotatfan.com/airline-voucher/internal/voucher"
)

const (
	serverReadHeaderTimeout = 5 * time.Second
	serverReadTimeout       = 15 * time.Second
	serverWriteTimeout      = 30 * time.Second
	serverIdleTimeout       = 60 * time.Second
	serverShutdownTimeout   = 10 * time.Second
)

func handleRequests() {
	db := config.InitDB()

	g := gin.Default()
	voucher.Register(g, db)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.GetConfig().Server.Port),
		Handler:           g,
		ReadHeaderTimeout: serverReadHeaderTimeout,
		ReadTimeout:       serverReadTimeout,
		WriteTimeout:      serverWriteTimeout,
		IdleTimeout:       serverIdleTimeout,
	}

	go func() {
		// service connections
		log.Printf("listening at port %d", config.GetConfig().Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}

	}()

	// Wait for interrupt signal to gracefully shutdown the server.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	if err := ctx.Err(); err == context.DeadlineExceeded {
		log.Printf("graceful shutdown timed out after %s", serverShutdownTimeout)
	}
	log.Println("Server exiting")
}

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	handleRequests()
}
