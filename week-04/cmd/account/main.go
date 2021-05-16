package main

import (
	"log"
	"context"
	"os/signal"
	"github.com/gin-gonic/gin"
	"week04/cmd/account/router"
)

var engine *gin.Engine = gin.New()

func init() {
	router.InitRouter(engine)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return engine.Run(":10002")
	})
	c := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	signal.Notify(c, sigs...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				cancel()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}
