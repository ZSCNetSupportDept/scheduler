.PHONY: clean help run build

help:
	./build/Help

clean:
	rm -rf $(TARGET_PATH)/*

FrontEnd:

BackEnd: FrontEnd
	cd $(SOURCE_PATH) && go build -o $(TARGET_PATH)/scheduler

build: BackEnd


run:
	TEMPLATE=$(TEMPLATE_DIR) FRONTEND=$(FRONTEND_PATH) CSV_PATH=$(CSV_PATH) $(TARGET_PATH)/scheduler --config $(CONFIG_FILE_PATH)

