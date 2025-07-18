.PHONY: clean help

help:
	./build/Help

clean:
	rm -rf $(TARGET_PATH)/*

scheduler:
	cd $(SOURCE_PATH) && go build -o $(TARGET_PATH)/scheduler

all: scheduler
