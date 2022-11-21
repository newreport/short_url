﻿import axios, { AxiosResponse } from "axios";
import { showMessage } from "./status"; // 引入状态码文件
import { ElMessage } from "element-plus"; // 引入el 提示框，这个项目里用什么组件库这里引什么
import myStore from "@store/index"; // 引入useStore 方法
import qs from "qs";

const store = myStore;

//https://cloud.tencent.com/developer/article/1916167
// 设置接口超时时间
// axios.defaults.timeout = 6000;

// 请求地址，这里是动态赋值的的环境变量，下一篇会细讲，这里跳过
// @ts-ignore
axios.defaults.baseURL = import.meta.env.VITE_API_DOMAIN;

// http request 拦截器
axios.interceptors.request.use(
  (config) => {
    // 配置请求头
    let acessToken = store.state.token.acessToken;
    if (acessToken && config.url != "/users/login"&&config.url != "/users/tocken/account") {
      acessToken = "Bearer "+acessToken;
    }
    config.headers = {
      "Content-Type":
        config.url != "/users/tocken/account"
          ? "application/json;"
          : "text/plain; charset=utf-8", // 传参方式json charset=UTF-8
        "Authorization":acessToken
    };
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// http response 拦截器
axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    const { response } = error;
    console.log("errorResponse:");
    console.log(error);
    if (response) {
      showMessage(response.status); // 传入响应码，匹配响应码对应信息
      if(response.status==401&&response.config.url!="/users/login"&&response.config.url!="/users/tocken/account"){
       let promise = axios({
          method: "post",
          url: "/users/tocken/account",
          data: store.state.token.refreshToken,
        });
        promise.then(result=>{
          if (result?.status == 200) {
            console.log("刷新了accesstoken：")
            console.log(result.data)
            store.commit('accessToken', result.data)
            return new Promise(
              (resolve: (value: AxiosResponse<any, any>) => void, reject) => {
                axios(error.config).then((res) => {
                  resolve(res);
                })
                .catch((err) => {
                  reject(err);
                });
              })
            
          }
        })
      }
      return Promise.reject(response.data);
    } else {
      ElMessage.warning("网络连接异常,请稍后再试!");
    }
  }
);

// 封装 GET POST 请求并导出
export function request(url = "", params = {}, type = "POST") {
  return new Promise(
    (resolve: (value: AxiosResponse<any, any>) => void, reject) => {
      let promise;
      if (type.toUpperCase() == "GET") {
        promise = axios({
          method: "get",
          url: url,
          params: params,
        });
      } else if (type.toUpperCase() == "POST") {
        promise = axios({
          method: "post",
          url: url,
          data: params,
        });
      } else if (type.toUpperCase() == "PUT") {
      }
      //处理返回
      promise
        .then((res) => {
          resolve(res);
        })
        .catch((err) => {
          reject(err);
        });
    }
  );
}
