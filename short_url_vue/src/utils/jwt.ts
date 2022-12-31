function JWT(){
    this.analysisClaims =  function (val: string) {
            let strings = val.split("."); //截取token，获取载体\
            let info = JSON.parse(
              decodeURIComponent(
                window.atob(strings[1].replace(/-/g, "+").replace(/_/g, "/"))
              )
            ); //解析，需要吧‘_’,'-'进行转换否则会无法解析
            return info
          }
}

export default JWT;