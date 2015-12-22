.PHONY: all build package install clean

release_files := ask_nest index.js nest.yml

all: build

build: ask_nest

ask_nest: ask_nest.go
	go test
	go build -ldflags "-X github.com/otherinbox/gobuild.BuildDate '$$(date)' -X github.com/otherinbox/gobuild.Version '$$(git rev-parse HEAD)'"

ask_nest.zip: $(release_files)
	zip ask_nest.zip -xi $(release_files)

package: ask_nest.zip

install: ask_nest.zip
	aws lambda update-function-code --function-name ask_nest --zip-file fileb://ask_nest.zip --publish

clean:
	rm ask_nest

cleandist: clean
	rm ask_nest.zip

