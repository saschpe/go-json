include $(GOROOT)/src/Make.inc

TARG=json
GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
	parser.go\
	json.go\

include $(GOROOT)/src/Make.cmd

format:
	${GOFMT} -w json.go
	#${GOFMT} -w examples/hello.go
