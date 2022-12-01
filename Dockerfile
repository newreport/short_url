FROM nginx:alpine
RUN mkdir /go && echo "test" > /go/1.text && \
wget https://golang.google.cn/dl/go1.19.3.linux-amd64.tar.gz && \
tar -C / -xzf go1.19.3.linux-amd64.tar.gz && \
rm -rf go1.19.3.linux-amd64.tar.gz && \
mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 && \
echo 'export PATH=$PATH:/go/bin' >> /etc/profile && \
source /etc/profile