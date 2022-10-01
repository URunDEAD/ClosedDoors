compile:
	go build -o bin/ClosedDoors cmd/closeddoors/main.go
clean:
	rm -rf bin
testing:
	docker-compose -f build/docker/docker-compose.yaml up --attach service --no-log-prefix --abort-on-container-exit --build
