build:
	go build

deploy:
	go build -o scheduler
	cp scheduler /opt/zscww_scheduler/scheduler
	cp -r FrontEnd /opt/zscww_scheduler/FrontEnd
	chmod +x /opt/zscww_scheduler/scheduler

