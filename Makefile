build:
	go build -ldflags="-s -w" -o ./bin/github-user-activity main.go

run:
	./bin/github-user-activity $(USER)
