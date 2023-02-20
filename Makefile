.PHONY: build clean

GORUN = env GO111MODULE=on go

build:
	$(GORUN) build -o report-port cmd/report-port.go 
	@echo "Done building."
	@echo "Run \"./report-port\" to launch report-port."

clean:
	env GO111MODULE=on go clean -cache
	rm -f ./report-port

