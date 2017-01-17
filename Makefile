VERSION = 0.8.13
ORGANIZATION=commands
REPOSITORY=the-enforcer
PORT=8813
base_dir=`pwd`
gopath="$(base_dir)/src:$(GOPATH)"

all: build
  
build: Dockerfile go-build
	docker build -t ${ORGANIZATION}/${REPOSITORY}:${VERSION} .

go-build:
	@echo "linux x86_64"
	@echo "GOPATH=$(gopath)"
	
	cd src;GOPATH=$(gopath) GOOS=linux GOARCH=amd64 go build -o ../bin/enforcer .

push:
	#docker push ${ORGANIZATION}/${REPOSITORY}:${VERSION}
	git add Dockerfile
	git push origin ${VERSION}

shell:
	docker run --name ${REPOSITORY} -it --rm -p ${PORT}:${PORT} ${ORGANIZATION}/${REPOSITORY}:${VERSION} /bin/sh

run:
	docker run --name ${REPOSITORY} --rm ${ORGANIZATION}/${REPOSITORY}:${VERSION}
