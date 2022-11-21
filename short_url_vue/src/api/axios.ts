import axios, { AxiosResponse } from "axios";
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
    config.headers = {
      //'Content-Type':'application/x-www-form-urlencoded',   // 传参方式表单
      "Content-Type": "application/json;", // 传参方式json charset=UTF-8
      // 'token':'80c483d59ca86ad0393cf8a98416e2a1'              // 这里自定义配置，这里传的是token
    };
    return config;
  },
  (error) => {
    console.log("请求error:");
    console.log(error);
    return Promise.reject(error);
  }
);

// http response 拦截器
axios.interceptors.response.use(
  (response) => {
    console.log("响应response:");
    console.log(response);
    return response;
  },
  (error) => {
    console.log("响应error:");
    console.log(error);
    console.log("响应errorResponse:");
    console.log(error.response);
    const { response } = error;

    if (response) {
      showMessage(response.status); // 传入响应码，匹配响应码对应信息
      return Promise.reject(response.data);
    } else {
      console.log("进入异常");
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
          console.log("promiseRES：");
          console.log(res);
          resolve(res);
        })
        .catch((err) => {
          console.log("promiseERR：");
          console.log(err);
          reject(err);
        });
    }
  );
}
