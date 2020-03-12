run:
ifdef url
	curl $(url) > .env && go run main.go
else
	@echo "masukan URL .env Anda \nContoh : \n - make run url=http://bla/bla/bla.env \n\n" && go run main.go
endif

build:
	go build *.go


autostart:
ifdef url
	curl $(url) > .env && reflex -r '\.go' -s -- sh -c 'go run main.go'
else
	@echo "masukan URL .env Anda \nContoh : \n - make run url=http://bla/bla/bla.env"
endif

init:
	rm -f glide.yaml && glide init