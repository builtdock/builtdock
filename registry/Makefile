include ../includes.mk

build: check-docker
	docker build -t builtdock/registry:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/registry:$(BUILD_TAG) $(REGISTRY)/builtdock/registry:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/registry:$(BUILD_TAG)
	deisctl config registry set image=$(REGISTRY)/builtdock/registry:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/registry:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/registry:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/registry | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/registry | xargs docker rmi -f

install: check-deisctl
	deisctl install registry-data
	deisctl scale registry=1

uninstall: check-deisctl
	deisctl scale registry=0
	deisctl uninstall registry-data

start: check-deisctl
	deisctl start registry

stop: check-deisctl
	deisctl stop registry

restart: stop start

run: install start

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
