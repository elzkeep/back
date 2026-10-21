FROM golang:1.19 AS gobuilder

WORKDIR /build
COPY ./ /build
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main main.go

FROM       alpine
MAINTAINER missbsd@gmail.com

COPY --from=gobuilder /build/main /usr/local/main/main
COPY ./.env.yml /usr/local/main/.env.yml
COPY ./certs /usr/local/main/certs
COPY web /usr/local/main/web
# 엑셀 양식·지로 OCR 폰트·점검 보고서 PDF 폰트 (Dockerfile_admin 과 같게)
ADD ./doc /usr/local/main/doc
RUN  mkdir -p /usr/local/main/fonts /usr/local/main/webdata/temp
COPY ./fonts/noto.ttf ./fonts/OCR-B1.ttf ./fonts/NotoSansKR-Regular.ttf ./fonts/NotoSansKR-Bold.ttf /usr/local/main/fonts/
#COPY ./fcm.json /usr/local/main/fcm.json
#RUN  mkdir -p /usr/local/main/webdata

WORKDIR /usr/local/main
CMD    ./main --mode production
