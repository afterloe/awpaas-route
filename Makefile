# create by afterloe <lm6289511@gmail.com>
# on 11-29-2018 22:38

# create by afterloe <lm6289511@gmail.com>
# # on 11-29-2018 22:38

.PHONY: build,compile,structure
SHELL := /bin/bash
WORKDIR = $(shell pwd)
VERSION = $(shell more package.json | grep version | awk -F '"' 'NR==1{print$$4}
')
NAME = $(shell more package.json | grep name | awk -F '"' 'NR==1{print$$4}')

all: compile structure

.ONESHELL:
compile: package.json
        docker run -it \
        -v $(src):/go/src \
        -v $(WORKDIR):/app \
        --rm \
        awpaas/builder:1.0.0 \
        go build -v

structure: app
        docker build -t awpaas/$(NAME):$(VERSION) .