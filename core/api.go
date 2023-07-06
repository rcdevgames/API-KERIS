package core

import (
	"QAPI/library"
	routes "QAPI/router"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func StartRest(port string) {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())
	routes.Init(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			library.Log.Err(err).Msg("Error Listen:")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	library.Log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		library.Log.Err(err).Msg("Server Shutdown:")
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		library.Log.Info().Msg("timeout of 5 seconds.")
	}
	library.Log.Info().Msg("Server exiting")
}
