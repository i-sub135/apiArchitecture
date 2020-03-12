run:
	go run main.go


build:
	go build *.go


autostart:
	reflex -r '\.go' -s -- sh -c 'go run main.go'

init:
	rm -f glide.yaml && glide init