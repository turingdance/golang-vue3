APPNAME=$(app)
BINARY_NAME=$(app)
SYSTEM = LINUX
ifeq ($(app),)
APPNAME := smartform
BINARY_NAME := smartform
endif
ifeq ($(OS),Windows_NT)
    SYSTEM := windows
    # Windows 可执行文件后缀（.exe）
    BINARY_NAME := $(BINARY_NAME).exe
endif
DATETIME=$(shell date +%Y%m%d%H%M)

.PHONY:build
server:
	@echo "build $(SYSTEM) as $(BINARY_NAME) DATATIME@$(DATETIME)"
	go mod tidy
	go build -ldflags "-w -s -X turingdance.com/turing/internal/conf.AppName=$(APPNAME) -X turingdance.com/turing/internal/conf.Version=$(DATETIME) -X turingdance.com/turing/internal/config.Commit=$(CI_COMMIT_SHA)" -o $(BINARY_NAME)

web:
	cd ../web/ && npm run build

build: web server


app:
	codectl reverse
	codectl router  -a winlion -s ./internal/app/$(app)/rest
	codectl runsql --dir ../sql/$(app)/	
reverse:
	codectl reverse
	
router:
	codectl router  -a winlion -s ./internal/app/$(app)/rest

.PHONY:clean
clean:
	go clean
	rm $(BINARY_NAME)