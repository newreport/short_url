export default {
  // 方法1
  testMethod1: function () {
    console.log("对，我只是一个测试公共方法");
  },
  // 方法2
  testMethod2: function (val) {
    console.log("巧了，我也是测试公共方法，我是", val);
  },
  calcJWT: function (val: string) {
    let strings = val.split("."); //截取token，获取载体
    let userinfo = JSON.parse(
      decodeURIComponent(
        window.atob(strings[1].replace(/-/g, "+").replace(/_/g, "/"))
      )
    ); //解析，需要吧‘_’,'-'进行转换否则会无法解析
    return userinfo
  },
};
