package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(port int, r http.Handler, socketPath string) error {
	server := &http.Server{Addr: fmt.Sprintf(":%v", port), Handler: r}
	var socket net.Listener
	if socketPath != "" {
		var err error
		socket, err = net.Listen("unix", socketPath)
		if err != nil {
			return err
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	go func() {
		<-sig
		log.Println("Stopping the server...")
		if socketPath != "" {
			_ = os.Remove(socketPath)
		}

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	if socketPath != "" {
		go func() {
			fmt.Printf("Server started, listening on socket: %v\n", socketPath)
			if err := server.Serve(socket); err != nil {
				fmt.Println("cannot listen on socket")
			}
		}()
	}

	fmt.Printf("Server started, listening on port: %v\n", port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	return nil
}
