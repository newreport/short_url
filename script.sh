﻿
bee api golang # 创建api
cd golang
go mod tidy
bee generate routers 
bee run -gendoc=true -downdoc=true #使用swagger运行

npm i   # 更新node包
yarn    # 更新node包
yarn dev    # 预览测试
vit vite --port=8888    # 预览测试



# https://cloud.tencent.com/developer/article/1574630 #linux打開文件過多