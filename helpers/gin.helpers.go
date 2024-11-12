package helpers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartGinServer(server *gin.Engine) error {
	srv := &http.Server{
		Addr:           ":3000",
		Handler:        server,
		IdleTimeout:    10 * time.Second,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := srv.ListenAndServe()

	return err
}
