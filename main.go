package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/calendar/v3"
)

const FS_NAME = "Whenfsgo"

func main() {
	arguments := os.Args
	fmt.Println(arguments)

	ctx := context.Background()
	calendar, err := calendar.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calendar)

	os.Exit(0)
}
