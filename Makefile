# create by afterloe <lm6289511@gmail.com>
# on 11-21-2017 14:30

.PHONY: remote,package,remove,buildJar,mkdirTmp,buildTar,clear,move

SHELL := /bin/bash
DOCKER_FILES := src/main/resources/application.yml docker-entrypoint.sh Dockerfile 

GROUPNAME = ascs
PROJECTNAME = $(shell more build.gradle | grep 'def artifactId' | awk -F '"' '{print $$2}')
VERSION = $(shell more build.gradle | grep 'version' | awk -F '"' 'NR==1{print $$2}')
DOCKERIMAGESNAME = $(GROUPNAME)/$(PROJECTNAME):$(VERSION)
TAR = $(PROJECTNAME)-$(VERSION).tar.gz

# 轮询指令
all: package

# 构建指令
package: remove buildJar mkdirTmp buildTar clear move

# remove - 删除上次生成的jar
.ONESHELL:
remove:
	rm -rf build/libs
	rm -rf $(TAR)
	rm -rf .docker

# buildJar - 构建jar包
buildJar:
	gradle build

# mkdirTmp - 构建docker file 相关包
.ONESHELL:
mkdirTmp: $(DOCKER_FILES) build/libs
	mkdir -p .docker 
	cp -R $(DOCKER_FILES) .docker
	cp $(shell find build/libs -name *.jar) .docker

.ONESHELL:
buildTar: .docker
	tar -czvf $(TAR) -C $< $(shell ls $<)

clear: .docker
	rm -rf $<

move: $(TAR) $(OUT_DIR)
	mv $< $(OUT_DIR)

# docker 导出
remote:
	docker save -o $(IMAGETARNAME) $(DOCKERIMAGESNAME) 
