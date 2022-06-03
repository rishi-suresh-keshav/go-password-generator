###
### build
###
.PHONY: build-gpg-cli
build-gpg-cli:
	go build -o gpg cmd/go-password-generator/main.go

###
### clean
###
.PHONY: clean-gpg-cli
clean-gpg-cli:
	rm -rf gpg


###
### build
###
.PHONY: build-gpg-server
build-gpg-server:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-password-generator server/main.go

###
### clean
###
.PHONY: run-gpg-server
run-gpg-server:
	./go-password-generator

###
### clean
###
.PHONY: clean-gpg-server
clean-gpg-server:
	rm -rf go-password-generator