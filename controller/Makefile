include ../includes.mk

.PHONY: all test logs

all: build run

build: check-docker
	docker build -t builtdock/controller:$(BUILD_TAG) .

push: check-docker check-registry check-deisctl
	docker tag builtdock/controller:$(BUILD_TAG) $(REGISTRY)/builtdock/controller:$(BUILD_TAG)
	docker push $(REGISTRY)/builtdock/controller:$(BUILD_TAG)
	deisctl config controller set image=$(REGISTRY)/builtdock/controller:$(BUILD_TAG)

clean: check-docker check-registry
	docker rmi builtdock/controller:$(BUILD_TAG)
	docker rmi $(REGISTRY)/builtdock/controller:$(BUILD_TAG)

full-clean: check-docker check-registry
	docker images -q builtdock/controller | xargs docker rmi -f
	docker images -q $(REGISTRY)/builtdock/controller | xargs docker rmi -f

install: check-deisctl
	deisctl scale controller=1

uninstall: check-deisctl
	deisctl scale controller=0

start: check-deisctl
	deisctl start controller

stop: check-deisctl
	deisctl stop controller

restart: stop start

run: install start

runserver:
	python manage.py runserver

db:
	python manage.py syncdb --migrate --noinput

coverage:
	coverage run manage.py test --noinput api
	coverage html

test: test-unit test-functional

setup-venv:
	@if [ ! -d venv ]; then virtualenv venv; fi
	venv/bin/pip install -q -r requirements.txt -r dev_requirements.txt

test-style: setup-venv
	venv/bin/flake8

test-unit: setup-venv test-style
	venv/bin/python manage.py test --noinput api

test-functional:
	@docker history builtdock/test-etcd >/dev/null 2>&1 || docker pull builtdock/test-etcd:latest
	@docker history builtdock/test-postgresql >/dev/null 2>&1 || docker pull builtdock/test-postgresql:latest
	GOPATH=$(CURDIR)/../tests/_vendor:$(GOPATH) go test -v ./tests/...
