package common

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run	优雅启停
func Run(r *gin.Engine, addr string, srvName string, stop func()) {

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	//	保证下面的优雅启停
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("%s failed to start, cause by: %v \n", srvName, err)
		}
	}()
	//	标记通道
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s... \n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if stop != nil {
		stop()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown, cause by : %v \n", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Printf("%s stop success... \n", srvName)
}
