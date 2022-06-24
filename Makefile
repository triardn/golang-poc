SHELL         = /bin/bash
APP_NAME      = golang-poc
VERSION       = $(shell git describe --always --tags)
GIT_COMMIT    = $(shell git rev-parse HEAD)
GIT_DIRTY     = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE    = $(shell date '+%Y-%m-%d-%H:%M:%S')
SQUAD         = "backend"
BUSINESS      = "campaign-management"

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                             Compile the project.'
	@echo '    make package                           Build, tag, and push Docker image.'
	@echo '    make deploy                            Deploy to Kubernetes via Helmfile.'
	@echo '    make rollback RELEASE= REVISION=       Rollback via Helm. If REVISION is omitted, it will roll back to the previous release.'
	@echo '    make run ARGS=                         Run with supplied arguments.'
	@echo '    make test                              Run unit-tests on a this project.'
	@echo '    make coverage                          Show your unit-test code coverage.'
	@echo '    make clean                             Clean the binary directory tree.'
	@echo

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -X github.com/kitabisa/${APP_NAME}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/kitabisa/${APP_NAME}/version.Version=${VERSION} -X github.com/kitabisa/${APP_NAME}/version.Environment=${ENV} -X github.com/kitabisa/${APP_NAME}/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

.PHONY: package
package:
	@echo "Build, tag, and push Docker image ${APP_NAME} ${VERSION} ${GIT_COMMIT}"
	docker buildx build \
		--build-arg VERSION=${VERSION},GIT_COMMIT=${GIT_COMMIT}${GIT_DIRTY} \
		--cache-from type=local,src=/tmp/.buildx-cache \
		--cache-to type=local,dest=/tmp/.buildx-cache \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:${GIT_COMMIT} \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:${VERSION} \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:latest \
		--push .

.PHONY: deploy
deploy:
	@echo "Deploying ${APP_NAME} ${VERSION}"
	export APP_NAME=${APP_NAME} && \
	export VERSION=${VERSION} && \
	export SQUAD=${SQUAD} && \
	export BUSINESS=${BUSINESS} && \
	helmfile apply

.PHONY: helm-history-length
helm-history-length:
	@helm history \
		--namespace ${APP_NAME} \
		--output yaml \
		${APP_NAME}-server-${ENV} | yq r - --length

.PHONY: helm-oldest-revision
helm-oldest-revision:
	@helm history \
		--namespace ${APP_NAME} \
		--output yaml \
		${APP_NAME}-server-${ENV} | yq r - "[0]".revision

.PHONY: helm-image-tag
helm-image-tag:
	@helm get values \
		--namespace ${APP_NAME} \
		--revision ${REVISION} \
		--output yaml \
		${APP_NAME}-server-${ENV} | yq r - image.tag

.PHONY: prune
prune:
	@echo "Removing Docker image ${DOCKER_REPOSITORY}/${APP_NAME}:${IMAGE_TAG}"
	gcloud container images delete \
		--force-delete-tags \
		--quiet \
		${DOCKER_REPOSITORY}/${APP_NAME}:${IMAGE_TAG}

.PHONY: rollback
rollback:
	@echo "Rollback ${RELEASE} ${REVISION}"
	helm rollback \
		--namespace ${APP_NAME} \
		${RELEASE} ${REVISION}

.PHONY: run
run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: test
test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test -race ./...

.PHONY: coverage
coverage:
	go test -race -coverprofile=coverage-reports.out ./...
	go tool cover -html=coverage-reports.out

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}

.PHONY: get-app-name
get-app-name:
	@echo ${APP_NAME}

.PHONY: get-business-unit
get-business-unit:
	@echo ${BUSINESS}

.PHONY: run-compose
run-compose:
	docker-compose -f docker-compose-dev.yaml up
