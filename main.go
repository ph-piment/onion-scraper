/*
package main
import (
	"context"
	"time"
)

func main() {
	e := InitializeEvent()

	e.Import(context.Background(), time.Now())
}
*/

package main

import "github.com/ph-piment/onion-scraper/cmd"

func main() {
	cmd.Execute()
}
