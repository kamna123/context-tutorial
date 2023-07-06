package main

import (
	"context"
	"fmt"
	"time"
)

func Authenticate(ctx context.Context, token string) bool {
	validToken := "secret_token"
	fmt.Println("-------request ID ------", ctx.Value("requestID"))
	ctx = context.WithValue(ctx, "authenticated", false)
	if token == validToken {
		ctx = context.WithValue(ctx, "authenticated", true)
	}
	return ctx.Value("authenticated").(bool)
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", "12345")
	go func(ctx context.Context) {
		isAuthenticated := Authenticate(ctx, "secret_token")
		fmt.Println("Authenticated:", isAuthenticated)
	}(ctx)
	ctx = context.WithValue(ctx, "requestID", "12346")
	go func(ctx context.Context) {
		isAuthenticated := Authenticate(ctx, "secret_token")
		fmt.Println("Authenticated:", isAuthenticated)
	}(ctx)

	// ...
	// Perform other concurrent operations
	// ...
	time.Sleep(1 * time.Second)
}
