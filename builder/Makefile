include ../includes.mk

build: check-docker
	docker build -t builtdock/builder:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/builder:$(BUILD_TAG) $(REGISTRY)/builtdock/builder:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/builder:$(BUILD_TAG)
	deisctl config builder set image=$(REGISTRY)/builtdock/builder:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/builder:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/builder:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/builder | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/builder | xargs docker rmi -f

install: check-deisctl
	deisctl install builder-data
	deisctl scale builder=1

uninstall: check-deisctl
	deisctl scale builder=0
	deisctl uninstall builder-data

start: check-deisctl
	deisctl start builder

stop: check-deisctl
	deisctl stop builder

restart: stop start

run: install start

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
