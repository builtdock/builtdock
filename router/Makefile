include ../includes.mk

build: check-docker
	docker build -t builtdock/router:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/router:$(BUILD_TAG) $(REGISTRY)/builtdock/router:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/router:$(BUILD_TAG)
	deisctl config router set image=$(REGISTRY)/builtdock/router:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/router:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/router:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/router | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/router | xargs docker rmi -f

install: check-deisctl
	deisctl scale router=1

uninstall: check-deisctl
	deisctl scale router=0

start: check-deisctl
	deisctl start router

stop: check-deisctl
	deisctl stop router

restart: stop start

run: install start

test: test-unit test-functional

test-unit:
	@echo no unit tests

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
