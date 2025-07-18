# 配置项
PWD = $(shell pwd)

SOURCE_PATH = $(PWD)/src
TARGET_PATH = $(PWD)/build/target
INSTALL_PATH = /opt/scheduler

include build/Makefile
