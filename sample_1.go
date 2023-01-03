package main

import (
	"context"
	"fmt"
)

func sample_1() {
	ctx1 := context.Background()
	ctx2 := context.TODO()

	fmt.Printf("%+v\n", ctx1) // context.Background
	fmt.Printf("%+v\n", ctx2) // context.TODO
}
