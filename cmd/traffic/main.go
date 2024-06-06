package main

import (
	"context"
	"fmt"
	"go-concurrency/internal/traffic"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	light := traffic.InitLight()
	s := light.Run(ctx, 1*time.Second)
	lane := traffic.InitLane[traffic.Car]()

	for i := 1; i < 1000; i++ {
		c := traffic.Car{
			ID: uint(i),
		}
		lane.Manage(c, s)
		fmt.Println((<-s).String(), lane)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}
