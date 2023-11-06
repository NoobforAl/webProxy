GO := go
SH := sh
GO_RUN := $(GO) run
GO_BUILD := $(GO) build
PATH_MAIN := ./main.go
BUILD_OUT := ./webPer
BUILD_FILE := ./build_release.sh

run:
	$(GO_RUN) $(PATH_MAIN)

build:
	$(GO_BUILD) -o $(BUILD_OUT) $(PATH_MAIN)

release:
	$(SH) $(BUILD_FILE)