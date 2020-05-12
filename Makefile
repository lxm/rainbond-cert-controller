LINUXBUILDER=CGOENABLE=0 go
all: certs-controller
certs-controller:
	${LINUXBUILDER} build -o dist/certs-controller cmd/certcontroller/main.go
clean:
	${LINUXBUILDER} clean && rm -rf dist/*
