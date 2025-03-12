.PHONY: install build clean tidy

tidy:
	go mod tidy

clean: tidy
	rm -f ~/.daiv/plugins/daiv-worklog.so
	rm -f out/daiv-worklog.so

build: clean
	go build -o out/daiv-worklog.so --buildmode=plugin main.go

install: build
	cp out/daiv-worklog.so ~/.daiv/plugins/daiv-worklog.so
