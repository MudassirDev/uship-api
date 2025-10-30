run:
	./out

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o out .

dev:
	air
