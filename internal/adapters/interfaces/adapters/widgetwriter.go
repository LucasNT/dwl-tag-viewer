package adapters

import (
	"context"

	"github.com/LucasNT/dwl-tag-viewer/internal/entities"
)

type WidgetWriter interface {
	Output(ctx context.Context, tags entities.DwlTagsInformation) error
}
