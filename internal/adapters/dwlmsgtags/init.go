package dwlmsgtags

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/LucasNT/dwl-tag-viewer/internal/entities"
)

type DwlMsgGetter struct{}

func (d DwlMsgGetter) GetTags(ctx context.Context) (entities.DwlTagsInformation, error) {
	var ret entities.DwlTagsInformation
	cmdContext, cmdCancel := context.WithCancel(ctx)
	defer cmdCancel()
	cmd := exec.CommandContext(cmdContext, "dwlmsg", "-g", "-t")
	output, err := cmd.Output()
	select {
	case <-ctx.Done():
		return ret, ctx.Err()
	default:
	}
	if err != nil {
		return ret, fmt.Errorf("Failed to execute dwlmsg %w", err)
	}
	i := len(output)
	outputString := string(output[0:i])
	j := 0
	var pos uint8 = 0
	var tag entities.DwlTag
	var sliceString string

	for z := 0; z < i; z++ {
		if outputString[z] != ' ' && outputString[z] != '\n' {
			continue
		}
		sliceString = outputString[j:z]
		switch pos {
		case 0:
			tag.Monitor = sliceString
			break
		case 1:
			if sliceString == "tags" {
				for outputString[z] != '\n' {
					z++
				}
				pos = 0
				j = z + 1
				continue
			}
			break
		case 2:
			aux, err := strconv.Atoi(sliceString)
			if err != nil {
				return entities.DwlTagsInformation{}, fmt.Errorf("Failed to convert tag value to int %w", err)
			}
			tag.Tag = uint8(aux)
			break
		case 3:
			tag.Viewed = '1' == sliceString[0]
			break
		case 4:
			aux, err := strconv.Atoi(sliceString)
			if err != nil {
				return entities.DwlTagsInformation{}, fmt.Errorf("Failed to convert number of elements value to int %w", err)
			}
			tag.NElement = uint8(aux)
			break
		case 5:
			tag.Focus = '1' == sliceString[0]
			pos = 255
			ret.Tags = append(ret.Tags, tag)
			break
		}
		pos++
		j = z + 1
		select {
		case <-ctx.Done():
			return entities.DwlTagsInformation{}, ctx.Err()
		default:
		}
	}
	return ret, nil
}
