package main

import (
	"context"
	"fmt"
)

const (
	key1 = "wakuwaku"
	key2 = "bank"
)

func fn1(ctx context.Context) {
	ctx = context.WithValue(ctx, key1, "fn1で値をセット")
	fmt.Printf("[fn1]\tctx:%v\twakuwaku:%v\tbank:%v\n", &ctx, ctx.Value(key1), ctx.Value(key2))
	fn2(ctx)
}

func fn2(ctx context.Context) {
	ctx = context.WithValue(ctx, key2, "fn2で値をセット")
	fmt.Printf("[fn2]\tctx:%v\twakuwaku:%v\tbank:%v\n", &ctx, ctx.Value(key1), ctx.Value(key2))
}

func fn3(ctx context.Context) {
	ctx = context.WithValue(ctx, key1, "fn3で値をセット")
	fmt.Printf("[fn3]\tctx:%v\twakuwaku:%v\tbank:%v\n", &ctx, ctx.Value(key1), ctx.Value(key2))
	fn4(ctx)
}

func fn4(ctx context.Context) {
	ctx = context.WithValue(ctx, key1, "fn4で値をセット")
	fmt.Printf("[fn4]\tctx:%v\twakuwaku:%v\tbank:%v\n", &ctx, ctx.Value(key1), ctx.Value(key2))
}

func sample_2() {
	ctx := context.Background()
	fmt.Printf("[main]\tctx:%v\twakuwaku:%v\tbank:%v\n", &ctx, ctx.Value(key1), ctx.Value(key2))
	fn1(ctx)
	fn3(ctx)
}
