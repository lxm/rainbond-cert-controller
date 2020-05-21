LINUXBUILDER=CGOENABLE=0 go
all: cert-controller cert-checker
cert-controller:
	${LINUXBUILDER} build -o dist/cert-controller cmd/certcontroller/main.go
cert-checker:
	${LINUXBUILDER} build -o dist/cert-checker cmd/certchecker/main.go
clean:
	${LINUXBUILDER} clean && rm -rf dist/*
