BINARY_NAME_LINUX=turing
BINARY_NAME=turing.exe
DATETIME=$(shell date +%Y%m%d%H%M)
OS := $(shell uname)
ifeq ($(OS), Linux)
	BINARY_NAME=$(BINARY_NAME_LINUX)
endif
.PHONY:build
buildserver:
	go mod tidy
	go build -o ${BINARY_NAME} -ldflags "-X turingdance.com/turing/internal/conf.Version=${DATETIME} -X turingdance.com/turing/internal/config.Commit=${CI_COMMIT_SHA}"

buildsite:
	cd ../web/ && npm install && npm run build

build: buildsite buildserver


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
	rm ${BINARY_NAME}