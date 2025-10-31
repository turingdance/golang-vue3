BINARY_NAME=etrain.exe
DATETIME=$(shell date +%Y%m%d%H%M)
.PHONY:build
buildserver:
	go mod tidy
	go build -o ${BINARY_NAME} -ldflags "turingdance.com/turing/internal/conf/conf.Version=${DATETIME} -X turingdance.com/turing/internal/conf/config.Commit=${CI_COMMIT_SHA}"

buildsite:
	cd ../console/ && npm run build

build: buildsite buildserver

menu:
	go run main.go menu -t cms系统 -n $(app) -o ./internal/app/$(app)/rest

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