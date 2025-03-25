package gettags

import (
	"context"

	"github.com/LucasNT/dwl-tag-viewer/internal/entities"
)

type GetTags interface {
	GetTag(ctx context.Context) (entities.DwlTagsInformation, error)
}
