package eww

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/LucasNT/dwl-tag-viewer/internal/entities"
)

type EwwTaskBar struct {
	output *os.File
}

func CreateEwwTaskBar(ctx context.Context, output *os.File) (EwwTaskBar, error) {
	ret := EwwTaskBar{}
	ret.output = output
	return ret, nil
}

func (e EwwTaskBar) Output(ctx context.Context, tags entities.DwlTagsInformation) error {
	fmt.Print("(box :orientation \"h\" :space-evenly false :halign \"start\" ")
	ctxChild, ctxChildCancel := context.WithCancel(ctx)
	defer ctxChildCancel()
	for _, v := range tags.Tags {
		aux, _ := e.item(ctxChild, v)
		fmt.Print(aux)
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}
	fmt.Print(")\n")
	return nil
}

func (e EwwTaskBar) item(ctx context.Context, tag entities.DwlTag) (string, error) {
	var builder strings.Builder
	builder.WriteString("(button :width 20 :vexpand false :onclick \"dwlmsg -s -t ")
	builder.WriteString(strconv.Itoa(int(tag.Tag)))
	builder.WriteString("\" ")

	builder.WriteString(":class \"block")

	if tag.Viewed {
		builder.WriteString(" block-active")
	}

	builder.WriteString("\" ")
	tag_string := e.convertTagToString(ctx, tag)

	if tag.NElement > 0 {
		builder.WriteString("(overlay  \"")
		builder.WriteString(tag_string)
		builder.WriteString("\" (box :class \"block-empty")
		if tag.Focus {
			builder.WriteString(" block-empty-active")
		}
		builder.WriteString("\" :valign \"start\" :halign \"start\")))")
	} else {
		builder.WriteString("\"")
		builder.WriteString(tag_string)
		builder.WriteString("\")")
	}
	return builder.String(), nil
}

func (e EwwTaskBar) convertTagToString(ctx context.Context, tag entities.DwlTag) string {
	var ret string = "f"
	if tag.Monitor == "eDP-1" {
		switch tag.Tag {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			ret = strconv.Itoa(int(tag.Tag + 1))
		case 4:
			ret = "q"
		case 5:
			ret = "w"
		case 6:
			ret = "e"
		case 7:
			ret = "r"
		}
	}
	return ret
}
