package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "some key1", "some value1")
	ctx = context.WithValue(ctx, "some key2", "some value2")

	getContextValues(ctx)
}

func getContextValues(ctx context.Context) {
	key := "some key1"
	fmt.Println(key+":", ctx.Value(key))
	key = "some key2"
	fmt.Println(key+":", ctx.Value(key))
}
