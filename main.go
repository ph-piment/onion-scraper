package main

import (
	"context"
	"time"
)

func main() {
	e := InitializeEvent()

	e.Import(context.Background(), time.Now())
}
