package main

import (
	"context"
	"fmt"
)

func main() {
	ProcessRequest("patika", "patika123")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "patikaID", userID)
	ctx = context.WithValue(ctx, "patikaToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"context içerisinde gelen veriler : %v (%v)",
		ctx.Value("patikaID"),
		ctx.Value("patikaToken"),
	)
}
