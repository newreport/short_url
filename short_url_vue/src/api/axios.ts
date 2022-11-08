﻿import axios, { AxiosResponse } from 'axios';
import { showMessage } from "./status";   // 引入状态码文件
import { ElMessage } from 'element-plus'  // 引入el 提示框，这个项目里用什么组件库这里引什么
import qs from 'qs'

//https://cloud.tencent.com/developer/article/1916167
// 设置接口超时时间
// axios.defaults.timeout = 6000;

// 请求地址，这里是动态赋值的的环境变量，下一篇会细讲，这里跳过
// @ts-ignore
axios.defaults.baseURL = import.meta.env.VITE_API_DOMAIN;   

// http request 拦截器
axios.interceptors.request.use(
  config => {
  // 配置请求头
    config.headers = {
      //'Content-Type':'application/x-www-form-urlencoded',   // 传参方式表单
      'Content-Type':'application/json;',        // 传参方式json charset=UTF-8
      // 'token':'80c483d59ca86ad0393cf8a98416e2a1'              // 这里自定义配置，这里传的是token
    };
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// http response 拦截器
axios.interceptors.response.use(
  response => {
    // console.log("response拦截器：")
    // console.log(response)
    console.log("进入拦截器response拦截器：:")
    console.log(response)
    return response;
  },
  error => {
    const {response} = error;
    // console.log("error拦截器：")
    // console.log(error)
    if (response) {
      // 请求已发出，但是不在2xx的范围
      console.log("进入拦截器error:")
      console.log(error)
      showMessage(response.status);           // 传入响应码，匹配响应码对应信息
      return Promise.reject(response.data);
    } else {
      ElMessage.warning('网络连接异常,请稍后再试!');
    }
  }
);

// 封装 GET POST 请求并导出
export function request(url='',params={},type='POST'){
//设置 url params type 的默认值
return new Promise((resolve:(value:AxiosResponse<any,any>)=>void,reject)=>{
  console.log(type)
  console.log(url)
  let promise
  if( type.toUpperCase()==='GET' ){
  console.log('请求了get')
    promise = axios({
      method:'get',
      url:url,
      params:params
    })
  }else if( type.toUpperCase()=== 'POST' ){
    console.log(params)
    console.log('请求了post'+url+params)
    promise = axios({
      method:'post',
      url:url,
      data:params
    })
  }
  //处理返回
  promise.then(res=>{
    console.log("promise")
    console.log(promise)
    resolve(res)
  }).catch(err=>{
    reject(err)
  })
})
}