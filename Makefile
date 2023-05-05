# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

OS = $(shell uname)

DOCKER_BUILD_EXTRA_ARGS ?=
# Export HOST_NETWORK=1 if you want to build the docker images with host network (useful when using some VPNs)
ifeq (${HOST_NETWORK}, 1)
	DOCKER_BUILD_EXTRA_ARGS += --network host
endif

# Project variables
PACKAGE = github.com/banzaicloud/bank-vaults
BINARY_NAME ?= bank-vaults
DOCKER_REGISTRY ?= ghcr.io/banzaicloud
DOCKER_IMAGE = ${DOCKER_REGISTRY}/bank-vaults
WEBHOOK_DOCKER_IMAGE = ${DOCKER_REGISTRY}/vault-secrets-webhook
OPERATOR_DOCKER_IMAGE = ${DOCKER_REGISTRY}/vault-operator
VAULT_ENV_DOCKER_IMAGE = ${DOCKER_REGISTRY}/vault-env

# Build variables
BUILD_DIR ?= build
BUILD_PACKAGE = ${PACKAGE}/cmd/...
VERSION ?= $(shell echo `git symbolic-ref -q --short HEAD || git describe --tags --exact-match` | tr '[/]' '-')
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS += -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE}
export CGO_ENABLED ?= 1
export GOOS = $(shell go env GOOS)
ifeq (${VERBOSE}, 1)
	GOARGS += -v
endif

# Docker variables
DOCKER_TAG ?= ${VERSION}

# Dependency versions
GOTESTSUM_VERSION = 0.4.0
GOLANGCI_VERSION = 1.52.2
LICENSEI_VERSION = 0.8.0
CODE_GENERATOR_VERSION = 0.27.1
CONTROLLER_GEN_VERSION = v0.11.4

GOLANG_VERSION = 1.19.2

## include "generic" targets
include main-targets.mk

.PHONY: up
up: ## Set up the development environment

.PHONY: down
down: clean ## Destroy the development environment


.PHONY: reset
reset: down up ## Reset the development environment


.PHONY: build-release
build-release: LDFLAGS += -w
build-release: build ## Build a binary without debug information

.PHONY: build-debug
build-debug: GOARGS += -gcflags "all=-N -l"
build-debug: BINARY_NAME_SUFFIX += debug
build-debug: build ## Build a binary with remote debugging capabilities

.PHONY: image-vault-env
image-vault-env: ## Build an OCI vault-env image
	buildah bud -t ${VAULT_ENV_DOCKER_IMAGE}:${DOCKER_TAG} -f Dockerfile.vault-env .
ifeq (${IMAGE_LATEST}, 1)
	buildah tag ${VAULT_ENV_DOCKER_IMAGE}:${DOCKER_TAG} ${VAULT_ENV_DOCKER_IMAGE}:latest
endif

.PHONY: docker-push
docker-push: ## Push a Docker image
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
ifeq (${DOCKER_LATEST}, 1)
	docker push ${DOCKER_IMAGE}:latest
endif

.PHONY: docker-operator
docker-operator: ## Build a Docker image for the Operator
	docker build ${DOCKER_BUILD_EXTRA_ARGS} -t ${OPERATOR_DOCKER_IMAGE}:${DOCKER_TAG} -f Dockerfile.operator .
ifeq (${DOCKER_LATEST}, 1)
	docker tag ${OPERATOR_DOCKER_IMAGE}:${DOCKER_TAG} ${OPERATOR_DOCKER_IMAGE}:latest
endif

.PHONY: docker-operator-push
docker-operator-push: ## Push a Docker image for the Operator
	docker push ${OPERATOR_DOCKER_IMAGE}:${DOCKER_TAG}
ifeq (${DOCKER_LATEST}, 1)
	docker push ${OPERATOR_DOCKER_IMAGE}:latest
endif


.PHONY: test-%
test-%: ## Run a specific test suite
	@${MAKE} VERBOSE=0 GOTAGS=$* test


.PHONY: operator-up
operator-up:
	kubectl replace -f deploy/crd.yaml || kubectl create -f deploy/crd.yaml
	kubectl apply -f deploy/rbac.yaml
	OPERATOR_NAME=vault-dev go run cmd/manager/main.go -verbose

.PHONY: operator-down
operator-down:
	kubectl delete -f deploy/crd.yaml
	kubectl delete -f deploy/rbac.yaml
