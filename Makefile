build:
	esc -o="staticassets.go" golang py27
	go build -o cjam github.com/fenrirunbound/cjam
