BUILD_DIR = build
SRC_DIR = src
SYSTEMD_DIR = systemd

APP_BUILD = $(BUILD_DIR)/wdal-app
WEB_BUILD = $(BUILD_DIR)/wdal-web

ETC_DIR = /etc/wdal
CONF_LOC = $(ETC_DIR)/conf.json

MODULE_JSON_DIR = $(SRC_DIR)/desktop/modules/json/

all: wdal-app wdal-web

json:
	mkdir $(MODULE_JSON_DIR)/build && cd $(MODULE_JSON_DIR)/build && cmake .. && cmake --build . --target install

wdal-app:
	g++ `pkg-config --cflags gtk+-3.0 webkit2gtk-4.0` -o build/wdal-app  $(SRC_DIR)/desktop/app.cpp `pkg-config --libs gtk+-3.0 webkit2gtk-4.0`

wdal-web:
	cd $(SRC_DIR)/site/ && go build -o ../../$(BUILD_DIR)/wdal-web

clean:
	rm -f $(BUILD_DIR)/*
	$(MAKE) -c $(MODULE_JSON_DIR) clean

install:
	cp -f $(APP_BUILD) /usr/bin
	cp -f $(WEB_BUILD) /usr/bin
	mkdir -p $(ETC_DIR)
	cp -n ./conf.ex.json $(CONF_LOC)
	cp -n $(SYSTEMD_DIR)/wdal-app.service /etc/systemd/system
	cp -n $(SYSTEMD_DIR)/wdal-web.service /etc/systemd/system