import { request } from './axios'

/**
 * @description -封装User类型的接口方法
 */
export class UserService {       // 模块一
    /**
     * @description 用户登录
     * @param {string} name - 用户名
     * @param {string} pwd - 密码
     * @return {HttpResponse} result
     */
    static async login(params) {   // 登录
        return request('/users/login',params, 'post')
    }
    /**
     * @description 刷新accessToken
     * @param {string} refreshToken refreshToken
     * @returns {HttpResponse} result
     */
    static async refreshToken(refreshToken) {   // 刷新token
        return request('/users/tocken/account',refreshToken,'post')
    }
    static async getall() {   // 接口二
        return request('/users/all',{},'get')
    }
}

export class ShortService{
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
     * @returns 
     */
    static async getShortsPage(params) {
        return request('/shorts',params,'get')
    }
    /**
     * @description 新增一个短链接
     * @param   {string}    shortGroup //分组
     * @param   {string}    sourceURL  //源链接
     * @param   {bool}  isEnable   //是否启用
     * @param   {string}    targetURL  //目标url
     * @param   {int}   length  //默认长度
     * @param   {string}    remarks //备注
     */
    static async createShort(data){
        return  request('/shorts/default_length/'+data.length,data,'post')
    }

    /**
     * @description 删除一个短链接
     * @param   {string}    sid  //短链接id
     */
    static async deleteShort(data){
        return  request('/shorts/'+data.sid,data,'delete')
    }

}

export class landRelevant {     // 模块二
    /**
     * @description 获取地列表
     * @return {HttpResponse} result
     */
    static async landList(params) {
        return request('/land_list_info',params, 'get')
    }
}
