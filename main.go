package main

import (
	"context"
	"time"

	"github.com/LucasNT/dwl-tag-viewer/internal/adapters/dwlmsgtags"
	"github.com/LucasNT/dwl-tag-viewer/internal/adapters/eww"
)

func main() {
	for {
		aux := dwlmsgtags.DwlMsgGetter{}
		a, err := aux.GetTags(context.Background())
		if err != nil {
			panic(err)
		}

		aux2 := eww.EwwTaskBar{}

		aux2.Output(context.Background(), a)

		d, _ := time.ParseDuration("300ms")

		time.Sleep(d)
	}

}
