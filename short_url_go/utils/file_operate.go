package utils

import "os"

//压缩文件zip
//http://c.biancheng.net/view/4583.html#:~:text=Go%E8%AF%AD%E8%A8%80%E7%9A%84%E6%A0%87%E5%87%86%E5%BA%93%E6%8F%90%E4%BE%9B%E4%BA%86%E5%AF%B9%E5%87%A0%E7%A7%8D%E5%8E%8B%E7%BC%A9%E6%A0%BC%E5%BC%8F%E7%9A%84%E6%94%AF%E6%8C%81%EF%BC%8C%E5%85%B6%E4%B8%AD%E5%8C%85%E6%8B%AC%20gzip%EF%BC%8C%E5%9B%A0%E6%AD%A4%20Go%20%E7%A8%8B%E5%BA%8F%E5%8F%AF%E4%BB%A5%E6%97%A0%E7%BC%9D%E5%9C%B0%E8%AF%BB%E5%86%99.gz%20%E6%89%A9%E5%B1%95%E5%90%8D%E7%9A%84%20gzip%20%E5%8E%8B%E7%BC%A9%E6%96%87%E4%BB%B6%E6%88%96%E9%9D%9E.gz,%E6%89%A9%E5%B1%95%E5%90%8D%E7%9A%84%E9%9D%9E%E5%8E%8B%E7%BC%A9%E6%96%87%E4%BB%B6%E3%80%82%20%E6%AD%A4%E5%A4%96%E6%A0%87%E5%87%86%E5%BA%93%E4%B9%9F%E6%8F%90%E4%BE%9B%E4%BA%86%E8%AF%BB%E5%92%8C%E5%86%99.zip%20%E6%96%87%E4%BB%B6%E3%80%81tar%20%E5%8C%85%E6%96%87%E4%BB%B6%EF%BC%88.tar%20%E5%92%8C.tar.gz%EF%BC%89%EF%BC%8C%E4%BB%A5%E5%8F%8A%E8%AF%BB.bz2%20%E6%96%87%E4%BB%B6%EF%BC%88%E5%8D%B3.tar.bz2%20%E6%96%87%E4%BB%B6%EF%BC%89%E7%9A%84%E5%8A%9F%E8%83%BD%E3%80%82

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
