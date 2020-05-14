LINUXBUILDER=CGOENABLE=0 go
all: certs-controller
certs-controller:
	${LINUXBUILDER} build -o dist/certs-controller cmd/certcontroller/main.go
cert-checker:
	${LINUXBUILDER} build -o dist/cert-checker cmd/certchecker/main.go
clean:
	${LINUXBUILDER} clean && rm -rf dist/*
