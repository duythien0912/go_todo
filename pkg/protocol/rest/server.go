package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"google.golang.org/grpc"

	"github.com/duythien0912/go_todo/pkg/api/v1"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("../../api/swagger/v1"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	gwmux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	mux.Handle("/", gwmux)

	fs = http.FileServer(http.Dir("../../web/swagger/dist"))

	prefix := "/docs/"

	mux.Handle(prefix, http.StripPrefix(prefix, fs))

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}
	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}

// func runSer() {
// 	fs := http.FileServer(http.Dir("../../api/swagger/v1"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))
// 	http.ListenAndServe(":"+"8081", nil)
// }
