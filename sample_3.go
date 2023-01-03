package main

import (
	"context"
	"fmt"
	"time"
)

func fn5(ctx context.Context) {
	log("start fn5")
	defer log("done fn5")
	for i := 1; i <= 4; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			log("loop fn5")
			time.Sleep(1 * time.Second)
		}
	}
}

func fn6(ctx context.Context) {
	log("start fn6")
	defer log("done fn6")
	for i := 1; i <= 4; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			log("loop fn6")
		}
	}
}

func log(timing string) {
	fmt.Printf("%s second:%v\n", timing, time.Now().Second())
}

func sample_3() {
	log("start main")
	defer log("done main")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go fn5(ctx)
	go fn6(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}
