all:
	CGO_ENABLED=0 go build -o dwl-tag-viewer main.go
