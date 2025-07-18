.PHONY: clean help run build start

help:
	./build/Help

clean:
	rm -rf $(TARGET_PATH)/*

FrontEnd:

BackEnd: FrontEnd
	cd $(SOURCE_PATH) && go build -o $(TARGET_PATH)/scheduler

build: BackEnd


run:
	SCHEDULER_DB_PATH="$(DATABASE)" $(TARGET_PATH)/scheduler --config $(CONFIG_FILE_PATH) --app.templatedir $(TEMPLATE_DIR) --app.memberfile $(CSV_PATH) --app.frontenddir $(FRONTEND_PATH) $(ARGS)

start: build run
