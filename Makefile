.PHONY: all clean inject

all: clean build inject

build:
	buf generate

inject:
	protoc-go-inject-tag -remove_tag_comment -input=./gen/go/unitedb/v1/*.pb.go

clean:
	rm -rf ./gen/*
