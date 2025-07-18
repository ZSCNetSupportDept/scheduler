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
	$(TARGET_PATH)/scheduler --config $(CONFIG_FILE_PATH) --csv-path $(CSV_PATH)

