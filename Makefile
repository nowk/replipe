default: install

build-go:
	 CGO_ENABLED=0 GOOS=linux go build \
		-ldflags "-s" \
		-a -installsuffix cgo \
		-o replipe .


install: build-go
	sudo cp $(shell pwd)/replipe /usr/local/bin/replipe
