
bee api golang # 创建api
cd golang
go mod tidy
 #使用swagger运行
bee generate routers & bee run -gendoc=true -downdoc=true 
nohup bee run -gendoc=true -downdoc=true  &
npm i   # 更新node包
yarn    # 更新node包
yarn dev    # 预览测试
vit vite --port=8888    # 预览测试

lsof -i:8080    @ #端口占用



# https://cloud.tencent.com/developer/article/1574630 #linux打開文件過多