build:
	go build

deploy:
	go build -o scheduler
	cp ./scheduler /opt/scheduler/scheduler
	cp -r FrontEnd /opt/scheduler/FrontEnd
	cp -r templates /opt/scheduler/templates
	cp -r tools /opt/scheduler/tools
	ln -s /opt/scheduler/tools/ZSCWW-Scheduler.service /etc/systemd/system/ZSCWW-Scheduler.service
	echo "部署脚本执行完成，记得放入以及在tools/start.sh中修改配置文件,在新环境运行时记得带上--init-db参数"

default: help

clean:
	rm ./scheduler
help:
	echo "请查看MakeFile文件查看构建选项"
