
build:
	go build -o github main.go

clean:
	-rm github

.PHONY:
	build clean
