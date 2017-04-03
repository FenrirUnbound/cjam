build:
	esc -o="staticassets.go" golang py27 node6
	go build -o cjam github.com/fenrirunbound/cjam
