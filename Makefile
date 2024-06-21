.PHONY: run build start clean

dev:
	@go run ./cmd/main.go

build:
	@echo "Construyendo aplicacion"
	@go build -v -o ./build/app_srv_tic.exe ./cmd/main.go

clean:
	@echo "Limpiando build"
	@rm -f ./build/app_srv_tic.exe

start: clean build
	@echo "Ejecutando servidor"
	./build/app_srv_tic.exe
