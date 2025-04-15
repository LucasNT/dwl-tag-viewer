package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/LucasNT/dwl-tag-viewer/internal/adapters/dwlmsgtags"
	"github.com/LucasNT/dwl-tag-viewer/internal/adapters/eww"
)

func main() {
	parentContext, stopFunction := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stopFunction()
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	for {
		aux := dwlmsgtags.DwlMsgGetter{}
		a, err := aux.GetTags(parentContext)
		if err != nil {
			panic(err)
		}

		aux2, _ := eww.CreateEwwTaskBar(parentContext, os.Args[1])

		aux2.Output(parentContext, a)

		d, _ := time.ParseDuration("300ms")

		time.Sleep(d)
	}

}
