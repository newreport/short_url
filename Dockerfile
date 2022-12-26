# 从 go 镜像进行编译
FROM golang:alpine AS gobuild
WORKDIR $GOPATH/src/short_url_go
# 设置代理
ENV GOPROXY https://goproxy.cn
# 解决 sqlite3 不兼容 gcc 10+
ENV CGO_CFLAGS -g -O2 -Wno-return-local-addr
# 复制源代码
COPY short_url_go/ .
# 更新 gcc 并构建
RUN apk add --update gcc musl-dev && \
go build -o app .

# 从 node 镜像 构建
FROM node:latest AS vuebuild
WORKDIR /short_url_vue
# 复制源代码
COPY short_url_vue/ .
# 删除本地的 lock 文件和 node_modules 文件
RUN rm -f package-lock.json && \
rm -rf yarn.lock && \
rm -rf node_modules/ && \
# 设置代理
npm config set registry https://registry.npm.taobao.org && \
npm config set disturl https://npm.taobao.org/dist && \
yarn config set registry http://registry.npm.taobao.org/ && \
yarn config set registry https://registry.npmjs.org/ && \
yarn config set chromedriver_cdnurl "https://npm.taobao.org/mirrors/chromedriver" && \
# 清理
yarn cache clean && \
# 更新 npm
npm install -g npm && \
# 安装 node-sass
npm i node-sass -D --verbose &&\
yarn install && \
# 构建
yarn run build && \
# echo 'www.example.com' > dist/CNAME

# 从 nginx 构建
FROM nginx:alpine
# 复制 go 二进制程序
COPY --from=gobuild /go/src/short_url_go/app /app/go/
# 复制 data 配置文件和源
COPY data/ /back
# 复制 vue 静态文件
COPY --from=vuebuild /short_url_vue/dist /app/vue/
# 数据文件配置
RUN mkdir -p /app/data && \
rm -f /back/go/db/main.db && \
cp /back /app/data && \
nohub /app/go/app & 

# ss
# FROM nginx:alpine
# RUN mkdir /go && echo "test" > /go/1.text && \
# wget https://golang.google.cn/dl/go1.19.3.linux-amd64.tar.gz && \
# tar -C /usr/local/ -xzf go1.19.3.linux-amd64.tar.gz && \
# rm -rf go1.19.3.linux-amd64.tar.gz && \
# mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
# echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile && \
# source /etc/profile