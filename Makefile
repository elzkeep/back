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
	rm -rf web
	cd ../app && make release-web version=$(version) && cp -rf build/web ../back/

docker: dockerbuild
	docker build --no-cache -t netb.co.kr:5000/zkeep:$(tag) .

dockerrun:
	docker run -d --name="zkeep" -p 9303:9303 netb.co.kr:5000/zkeep

push: docker
	docker push netb.co.kr:5000/zkeep:$(tag)

deploy: push
	docker-compose --context dev pull
	docker-compose --context dev up -d

localdeploy: 
	git pull
	cd ../app && make release && cp -rf build/web ../back/	
	docker build --no-cache -t netb.co.kr:5000/zkeep:$(tag) .
	docker push netb.co.kr:5000/zkeep:$(tag)
	docker-compose --context dev pull
	docker-compose --context dev up -d

clean:
	rm -f bin/main

dummy:
