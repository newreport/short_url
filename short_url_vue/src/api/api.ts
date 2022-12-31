import { request } from "./axios";

/**
 * @description -封装User类型的接口方法
 */
export class UserService {
  /**
   * @description 用户登录
   * @param {string} name - 用户名
   * @param {string} pwd - 密码
   * @return {HttpResponse} result
   */
  static async login(params) {
    // 登录
    return request("/users/login", params, "post");
  }
  /**
   * @description 刷新accessToken
   * @param {string} refreshToken refreshToken
   * @returns {HttpResponse} result
   */
  static async refreshToken(refreshToken) {
    // 刷新token
    return request("/users/tocken/account", refreshToken, "post");
  }
  /**
   * @description 根据id删除用户
   * @param {number}  uid 用户id
   * @return {HttpResponse} result
   */
  static async deleteUser(uid) {
    return request("/users/" + uid, uid, "delete");
  }

  /**
   * @description 分页查询
   * @param {int} offset  偏移量
   * @param {int} limit   返回行数
   * @param {string}  sort  排序
   * @param {string}  name  账号名
   * @param {string}  nickname  昵称
   * @param {string}  phone 手机号
   * @return  {HttpResponse}  result
   */
  static async getUsersPage(params) {
    return request("/users", params, "get");
  }

  /**
   * @description 新增一个用户
   * @param {Blob} author 头像
   * @param {boolean} autoInsertSpace 盘古之白
   * @param {string} domain 域名
   * @param {string} group 分组
   * @param {string} i18n 国际化
   * @param {string} name 用户名
   * @param {string} nickname 昵称
   * @param {string} phone 联系方式
   * @param {string} pwd 密码
   * @param {string} remarks 备注
   * @param {int} role 权限
   * @param {int} urlLength 默认长度
   */
  static async addUser(params){
    return request("/users", params, "post");
  }
  /**
   * @description 修改用户信息
   * @param
   */
  static async updateUser() {}

  static async updateUserPasswd() {}
}

export class ShortService {
  /**
   * @description 分页查询
   * @param   {int} offset  //偏移量
   * @param   {int} limit   //返回行数
   * @param   {string}  sort  //排序
   * @param   {string}    sorrce_url  //源链接
   * @param   {string}    target_url  //目标链接
   * @param   {string}    group   //分组
   * @param   {string}    is_enable   //是否启用
   * @param   {string}    exp //过期时间
   * @param   {string}    crt //创建时间
   * @param   {string}    upt //修改时间
   * @param   {string}    del //删除时间
   * @returns {HttpResponse} result
   */
  static async getShortsPage(params) {
    return request("/shorts", params, "get");
  }
  /**
   * @description 新增一个短链接
   * @param   {string}    sourceURL  //源链接
   * @param   {bool}  automactic  //自动生成
   * @param   {int}   length  //默认长度
   * @param   {string}    targetURL  //目标url(长度大于0时为自定义链接)
   * @param   {string}    shortGroup //分组
   * @param   {bool}  isEnable   //是否启用
   * @param   {string}    remarks //备注
   * @param   {string}  exp //过期时间
   * @returns {HttpResponse} result
   */
  static async createShort(data) {
    return request("/shorts", data, "post");
  }

  /**
   * @description 修改一个短链接(全部信息)
   * @param   {string}    id  //短链接id
   * @param   {bool}  automactic  //自动生成
   * @param   {int}   length  //默认长度
   * @param   {string}    targetURL  //目标url(长度大于0时为自定义链接)
   * @param   {string}    shortGroup //分组
   * @param   {bool}  isEnable   //是否启用
   * @param   {string}    remarks //备注
   * @param   {string}  exp //过期时间
   * @returns {HttpResponse} result
   */
  static async updateShort(data) {
    return request("/shorts/" + data.id, data, "put");
  }

  /**
   * @description   修改一个短链接是否启用
   * @param {string}    id //短链接id
   * @param {bool}  isEnable    //是否启用
   * @returns {HttpResponse} result
   * @returns
   */
  static async updateShortIsEnable(data) {
    return request("/shorts/" + data.id + "/is_enable", data, "patch");
  }
  /**
   * @description 删除一个短链接
   * @param   {string}    id  //短链接id
   * @returns {HttpResponse} result
   */
  static async deleteShort(data) {
    return request("/shorts/" + data.id, data, "delete");
  }
}
