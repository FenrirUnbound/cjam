build:
	esc -o="staticassets.go" golang
	go build -o cjam github.com/fenrirunbound/cjam
