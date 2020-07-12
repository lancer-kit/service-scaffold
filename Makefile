build:
	./build.sh ./app

build_debug:
	go build -gcflags "all=-N -l" -o ./app .

run_debug: build_debug
	dlv --listen=:40000 --headless=true --api-version=2 exec ./app serve

build_docker:
	docker build -t $$image --build-arg CONFIG=$$config .

build_debug_docker:
	docker build -f debug.docker -t $$image --build-arg CONFIG=$$config .

lint:
	golangci-lint run --out-format tab

test:
	go test ./...

generate:
	go generate -v ./...
