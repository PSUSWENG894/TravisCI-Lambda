package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
    "log"
)

type MyEvent struct {
	Name string `json: "name"`
}

func AlexaTravisCISkill(ctx context.Context, name MyEvent) (string, error) {
    log.Print("Handling event " + name.Name)
    return fmt.Sprintf("Hello %s!", name.Name), nil
}

func saveCurrentUser(user string) string {
	return "You asked to save current user: " + user
}

func main() {
    fmt.Printf("Hello, world!\n")
    lambda.Start(AlexaTravisCISkill)
}