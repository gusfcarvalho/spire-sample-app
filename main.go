package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func main() {
	socket := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if socket == "" {
		socket = "unix:///run/spire/sockets/public/api.sock"
	}
	fmt.Println("SPIFFE Endpoint Socket:", socket)

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gracefully shutting down.")
			return
		default:
			resp, err := workloadapi.FetchX509Context(ctx)
			if err != nil {
				fmt.Printf("error fetching SVID: %v", err)
			} else {
				fmt.Printf("\nFetched %d SVIDS:\n", len(resp.SVIDs))
				for _, s := range resp.SVIDs {
					fmt.Println("- ", s.ID.String())
				}
			}
			time.Sleep(5 * time.Second)
		}
	}
}
