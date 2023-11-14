BINARY_NAME=longarms
PKG_VERSION=1.0.0

build:
	go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

# Runs go clean and deletes any packages generated from this file.
clean:
	go clean
	rm -f ./longarms-${PKG_VERSION}-amd64.deb
	rm -f ./longarms-${PKG_VERSION}-amd64.tar.gz

# runs build and uses fpm to create a .deb package from the resulting executable
deb: build
	fpm -s dir -t deb -p longarms-${PKG_VERSION}-amd64.deb --name longarms --license MIT --version ${PKG_VERSION} --architecture amd64 --description "A simple tool to recognize multitouch taps." --url "https://github.com/mxash/longarms" --maintainer "Ash Wilburn <mxashwilburn@icloud.com>" longarms=/usr/bin/longarms

# runs build and uses fpm to create a .tar.gz archive of the resulting executable.
tar: build
	fpm -s dir -t tar -p longarms-${PKG_VERSION}-amd64.tar.gz --name longarms longarms