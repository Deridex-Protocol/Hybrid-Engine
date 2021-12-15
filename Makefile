api:
	go run ./cmd/api/main.go

websocket:
	go run ./cmd/websocket/main.go

watcher:
	go run ./cmd/watcher/main.go

engine:
	go run ./cmd/engine/main.go

launcher:
	go run ./cmd/launcher/main.go

maker:
	go run ./cmd/maker/main.go

docs:
	swag init -o services/api/docs/ -d services/api/ -g service.go --parseDependency=true

contract:
	abigen --abi=./contracts/Deridex.abi --pkg=contract --type=Contract --out=./common/contract/contract.go
	abigen --abi=./contracts/DeridexTrade.abi --pkg=contract --type=Trade --out=./common/contract/trade.go

redeploy:
	docker-compose down
	docker-compose build
	docker-compose up -d
	yes Y | docker system prune

.PHONY: test api ws watcher engine launcher
