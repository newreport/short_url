FROM golang:alpine AS gobuild
WORKDIR $GOPATH/src/short_url_go
ENV GOPROXY https://goproxy.cn
ENV CGO_CFLAGS -g -O2 -Wno-return-local-addr
COPY short_url_go/ .
RUN apk add --update gcc musl-dev && \
go build -o app .

FROM node:latest AS vuebuild
WORKDIR /short_url_vue
COPY short_url_vue/ .
RUN rm -f package-lock.json && \
rm -rf yarn.lock && \
rm -rf node_modules/ && \
npm config set registry https://registry.npm.taobao.org && \
npm config set disturl https://npm.taobao.org/dist && \
yarn config set registry http://registry.npm.taobao.org/ && \
yarn config set registry https://registry.npmjs.org/ && \
yarn config set chromedriver_cdnurl "https://npm.taobao.org/mirrors/chromedriver" && \
yarn cache clean && \
npm install -g npm && \
npm i node-sass -D --verbose &&\
yarn install && \
yarn run build 

FROM nginx:alpine
COPY --from=gobuild /go/src/short_url_go/app /app/go/
COPY --from=gobuild /go/src/data /app/data
COPY --from=vuebuild /short_url_vue/dist /app/vue/


# FROM nginx:alpine
# RUN mkdir /go && echo "test" > /go/1.text && \
# wget https://golang.google.cn/dl/go1.19.3.linux-amd64.tar.gz && \
# tar -C /usr/local/ -xzf go1.19.3.linux-amd64.tar.gz && \
# rm -rf go1.19.3.linux-amd64.tar.gz && \
# mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
# echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile && \
# source /etc/profile