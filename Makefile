# Nome do projeto
PROJECT_NAME := $(shell basename $(PWD))

# Diretório principal
CMD_DIR := ./cmd/

# Variáveis de compilação
BUILD_DIR := ./build
BINARY_NAME := $(PROJECT_NAME)
BINARY_PATH := $(BUILD_DIR)/$(BINARY_NAME)

# Variáveis de teste
TEST_FLAGS := -v -race
COVER_FILE := coverage.out

# Modo vendor (usar -mod=vendor quando a pasta vendor existir)
VENDOR_EXISTS := $(shell [ -d ./vendor ] && echo 1)
MOD_FLAGS := $(if $(VENDOR_EXISTS),-mod=vendor,)

# Variáveis Docker
APP_NAME   := json-post-box
DOCKER_TAG := latest
PORT       := 8030
PORT_HOST       := 8030

# Cores para saída
GREEN := \033[0;32m
NC := \033[0m

# Opção help padrão
.DEFAULT_GOAL := help

.PHONY: build
build: ## Compila o projeto
	@echo "$(GREEN)Compilando aplicação...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build $(MOD_FLAGS) -o $(BINARY_PATH) $(CMD_DIR)

.PHONY: run
run: clean build ## Executa o projeto
	@echo "$(GREEN)Executando aplicação...$(NC)"
	@go run $(MOD_FLAGS) $(CMD_DIR)

.PHONY: cover
cover: ## Executa os testes com cobertura
	@echo "$(GREEN)Executando testes com cobertura...$(NC)"
	@go test $(MOD_FLAGS) $(TEST_FLAGS) -coverprofile=$(COVER_FILE) ./...
	@go tool cover -html=$(COVER_FILE)

.PHONY: fmt
fmt: ## Formata o código
	@echo "$(GREEN)Formatando código...$(NC)"
	@go fmt ./...

.PHONY: vet
vet: ## Analisa código em busca de erros
	@echo "$(GREEN)Analisando código...$(NC)"
	@go vet $(MOD_FLAGS) ./...

.PHONY: lint
lint: ## Executa golangci-lint
	@echo "$(GREEN)Executando linter...$(NC)"
	@golangci-lint run $(MOD_FLAGS)

.PHONY: clean
clean: ## Remove arquivos gerados
	@echo "$(GREEN)Limpando arquivos gerados...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f $(COVER_FILE)

.PHONY: vendor
vendor: ## Cria e atualiza a pasta vendor
	@echo "$(GREEN)Atualizando dependências e criando vendor...$(NC)"
	@go mod tidy
	@go mod vendor

.PHONY: update
update: ## Atualiza dependências
	@echo "$(GREEN)Atualizando dependências...$(NC)"
	@go get -u ./...
	@go mod tidy
	@go mod vendor

.PHONY: docker-build
docker-build: ## Constrói a imagem Docker
	@echo "$(GREEN)Construindo imagem Docker...$(NC)"
	@docker build -t $(APP_NAME):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: docker-rm ## Executa o container Docker
	@echo "$(GREEN)Iniciando container na porta $(PORT)...$(NC)"
	docker run -p $(PORT_HOST):$(PORT) --name $(APP_NAME) $(APP_NAME):$(DOCKER_TAG)

.PHONY: docker-run-detach
docker-run-detach: docker-rm ## Executa o container Docker em modo detached (background)
	@echo "$(GREEN)Iniciando container em modo detached na porta $(PORT)...$(NC)"
	docker run -d -p $(PORT_HOST):$(PORT) --name $(APP_NAME) $(APP_NAME):$(DOCKER_TAG)

.PHONY: docker-stop
docker-stop: ## Para o container Docker
	@echo "$(GREEN)Parando container...$(NC)"
	@docker stop $(APP_NAME) || true

.PHONY: docker-rm
docker-rm: ## Remove o container Docker
	@echo "$(GREEN)Removendo container...$(NC)"
	@docker rm $(APP_NAME) || true

.PHONY: docker-clean
docker-clean: docker-stop docker-rm ## Limpa containers e imagens Docker
	@echo "$(GREEN)Removendo imagem...$(NC)"
	@docker rmi $(APP_NAME):$(DOCKER_TAG) || true

.PHONY: docker-rebuild
docker-rebuild: docker-clean docker-build ## Reconstroi a imagem Docker do zero

.PHONY: docker-up
docker-up: docker-build docker-run ## Build + Run (fluxo completo)

.PHONY: docker-logs
docker-logs: ## Mostra logs do container
	@docker logs -f $(APP_NAME)

.PHONY: docker-shell
docker-shell: ## Acessa o shell do container
	@docker exec -it $(APP_NAME) /bin/sh

.PHONY: help
help: ## Mostra esta ajuda
	@echo "Por favor, use 'make <alvo>' onde <alvo> é um dos seguintes:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Development tools section
.PHONY: install-tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	go install gotest.tools/gotestsum@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

# Testing targets
.PHONY: test
test: ## Run tests with formatted output
	@echo "$(GREEN)Executando testes...$(NC)"
	gotestsum --format testname -- -v -short ./...

.PHONY: test-ci
test-ci: ## CI-friendly test output
	gotestsum --format dots --junitfile test-results.xml -- -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	gotestsum --format testname -- -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

# Linting and validation
.PHONY: lint
lint: ## Run linters
	golangci-lint run ./...
	staticcheck ./...

.PHONY: vuln-check
vuln-check: ## Check for vulnerabilities
	govulncheck ./...