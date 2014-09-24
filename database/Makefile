include ../includes.mk

build: check-docker
	docker build -t builtdock/database:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/database:$(BUILD_TAG) $(REGISTRY)/builtdock/database:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/database:$(BUILD_TAG)
	deisctl config database set image=$(REGISTRY)/builtdock/database:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/database:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/database:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/database | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/database | xargs docker rmi -f

install: check-deisctl
	deisctl install database-data
	deisctl scale database=1

uninstall: check-deisctl
	deisctl scale database=0
	deisctl uninstall database-data

start: check-deisctl
	deisctl start database

stop: check-deisctl
	deisctl stop database

restart: stop start

run: install start

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
