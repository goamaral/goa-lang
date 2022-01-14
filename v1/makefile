all: build clean

build: parser/parser.y
	goyacc -o parser/parser.go -v parser.output parser/parser.y

clean:
	rm -f parser.output

test:
	./bin/test