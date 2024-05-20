run:
	@echo Running GameOfLifeAPP . . . . . . 
	go run cmd/gameoflife/main.go

build:
	@echo Building DistanceApp . . . . . .
	go build cmd/gameoflife/main.go
	@echo Executing the build
	./main

test:
	@echo Running all the tests
	go test ./...
	