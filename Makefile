tag=latest

all: server

server: dummy
	buildtool-model ./ 
	buildtool-model ./ dart
	buildtool-router_fiber ./ > ./router/router.go
	go build -o bin/main main.go

model: dummy
	buildtool-model ./ 
	buildtool-model ./ dart
	buildtool-router_fiber ./ > ./router/router.go
	go build -o bin/main main.go

router: dummy
	buildtool-router_fiber ./ > ./router/router.go
	go build -o bin/main main.go

build: dummy
	go build -o bin/main main.go

run:
	buildtool-watch ./	

test: dummy
	#go test -v ./...
	go test

linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/main.linux main.go

dockerbuild:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o bin/main.linux main.go

docker: 
	docker build --no-cache -t netb.co.kr:5000/aoi:$(tag) .

dockerrun:
	docker run -d --name="aoi" -p 9301:9301 netb.co.kr:5000/aoi

push: docker
	docker push netb.co.kr:5000/aoi:$(tag)

deploy: push
	docker-compose --context dev pull
	docker-compose --context dev up -d

localdeploy: 
	git pull
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o bin/main.linux main.go
	cd ../front && npm run release && cp -rf dist ../back/
	docker build --no-cache -t netb.co.kr:5000/aoi:$(tag) .
	docker push netb.co.kr:5000/aoi:$(tag)
	docker-compose --context dev pull
	docker-compose --context dev up -d

clean:
	rm -f bin/main

dummy:
