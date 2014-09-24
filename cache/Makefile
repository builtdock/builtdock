include ../includes.mk

build: check-docker
	docker build -t builtdock/cache:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/cache:$(BUILD_TAG) $(REGISTRY)/builtdock/cache:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/cache:$(BUILD_TAG)
	deisctl config cache set image=$(REGISTRY)/builtdock/cache:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/cache:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/cache:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/cache | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/cache | xargs docker rmi -f

install: check-deisctl
	deisctl scale cache=1

uninstall: check-deisctl
	deisctl scale cache=0

start: check-deisctl
	deisctl start cache

stop: check-deisctl
	deisctl stop cache

restart: stop start

run: install start

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
