compile:
	go build -o bin/ClosedDoors cmd/closeddoors/main.go
clean:
	rm -rf bin
testing:
	docker-compose -f build/docker/docker-compose.yaml up test_service --attach test_service --no-log-prefix --abort-on-container-exit --build
	docker-compose -f build/docker/docker-compose.yaml down
db-up:
	docker-compose -f build/docker/docker-compose.yaml up db -d
db-down:
	docker-compose -f build/docker/docker-compose.yaml down