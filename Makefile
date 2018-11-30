# create by afterloe <lm6289511@gmail.com>
# on 11-29-2018 22:38

.PHONY: build,compile,structure
SHELL := /bin/bash
WORKDIR = $(shell pwd)
BUILD_IMG = awpaas/builder
VERSION = $(shell more package.json | grep version | awk -F '"' 'NR==1{print$$4}')
NAME = $(shell more package.json | grep name | awk -F '"' 'NR==1{print$$4}')
BUILD_ENV = $(shell docker image ls | grep ${BUILD_IMG} | wc -l)

all: compile structure

.ONESHELL:
compile: $(src) package.json
	docker run -it -v $(src):/go/src -v $(WORKDIR):/app --rm ${BUILD_IMG}:1.0.0 go build -v
structure: app
	docker build -t awpaas/$(NAME):$(VERSION) .