build:
	go build -ldflags="-s -w" -o ./bin/github-activity main.go

run:
	./bin/github-activity $(USER)
