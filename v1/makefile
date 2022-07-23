all: build clean

build: parser_yacc/parser.y
	goyacc -o parser_yacc/parser.go -v parser.output parser_yacc/parser.y

clean:
	rm -f parser.output

test:
	./bin/test