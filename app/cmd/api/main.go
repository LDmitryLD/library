package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"projects/LDmitryLD/library/app/config"
	"projects/LDmitryLD/library/app/internal/db"
	"projects/LDmitryLD/library/app/internal/inrfrastructure/router"
	"projects/LDmitryLD/library/app/internal/modules"
	"projects/LDmitryLD/library/app/internal/storages"
	"syscall"
	"time"
)

func main() {
	confDB := config.NewAppConf().DB

	_, sqlAdapter, err := db.NewSqlDB(confDB)
	if err != nil {
		log.Fatal("ошибка при подключении к бд: ", err)
	}

	storages := storages.NewStorages(sqlAdapter)

	services := modules.NewServices(storages)

	controllers := modules.NewControllers(services)

	r := router.NewRouter(*controllers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("Starting server")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}
