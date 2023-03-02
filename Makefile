up:
	docker-compose up -d

down:
	docker-compose down

run:
	go run ./cmd/playlistService/main.go

store_mock:
	mockgen -source=internal/repository/repo_interface.go -destination=internal/repository/mocks/store_mock.go -package=mocks

usecase_mock:
	mockgen -source=internal/usecases/interface.go -destination=internal/usecases/mocks/usecases_mock.go -package=mocks

test:
	go test ./... -v

client:
	go run ./cmd/client/main.go