package main

import (
	"context"
	"time"
)

func fn7(ctx context.Context) {
	log("start fn7")
	defer log("done fn7")
	for i := 1; i <= 4; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			log("loop fn7")
			time.Sleep(1 * time.Second)
		}
	}
}

func fn8(ctx context.Context) {
	log("start fn8")
	defer log("done fn8")
	for i := 1; i <= 4; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			log("loop fn8")
		}
	}
}

func sample_4() {
	log("start main")
	defer log("done main")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go fn7(ctx)
	go fn8(ctx)
	time.Sleep(5 * time.Second)
}
