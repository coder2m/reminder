.PHONY: build-image
build-image: build-app-server

.PHONY: build-app-server
build-app-server:
	docker build -f build/app/Dockerfile -t reminder-app:v1 .

.PHONY: run-app
run-app:
	go run cmd/app/main.go run -c=config/confin.toml

.PHONY: build-app
build-app:
	go build -o  cmd/app/app cmd/app/main.go