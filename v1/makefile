all: compile

compile: clean parser.y
	goyacc -o parser.go -v parser.output parser.y

clean:
	rm parser.output parser.go