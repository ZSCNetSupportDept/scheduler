build:
	go build

deploy:
	go build -o scheduler
	cp ./scheduler /opt/zscww_scheduler/scheduler
	cp -r FrontEnd /opt/zscww_scheduler/FrontEnd
	chmod +x /opt/zscww_scheduler/scheduler
	echo "部署脚本执行完成，记得放入member.csv数据文件"

default: help

clean:
	rm ./scheduler
help:
	echo "请查看MakeFile文件查看构建选项"
