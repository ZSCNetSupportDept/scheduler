PWD = $(shell pwd)
SOURCE_PATH = $(PWD)/src

# **构建配置项**

 # 如果构建，将构建目标放到哪个目录下
TARGET_PATH = $(PWD)/build/target
 # 如果安装，将程序安装到哪里
INSTALL_PATH = /opt/scheduler

# 是否将前端嵌入到最终二进制文件里？
EMBED_FRONTEND = 1

# **运行配置项**

# 如果运行，使用的配置文件在哪里？
CONFIG_FILE_PATH = $(PWD)/ignore/secret.yaml
# 如果运行，使用的成员信息文件在哪里？
CSV_PATH = $(PWD)/ignore/aa.csv

include build/Makefile
