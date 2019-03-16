PKGS := ./internal/... ./cmd/...

test: dep
	go test $(PKGS)

lint:
	golint $(PKGS)

build: dep
	go build -o bin/srmessenging

install: dep
	go install

clean:
	go clean -testcache

dep:
	dep ensure

run-kafka:
	cd scripts && docker-compose up -d

run: run-kafka
	go run main.go run
