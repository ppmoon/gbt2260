# 中国行政区划代码
中华人民共和国县以上行政区划代码的golang包

## 安装

````bash
go get -u github.com/ppmoon/gbt2260
````
或者你使用go module(vgo)
````bash
go get -u github.com/ppmoon/gbt2260
````
## 如何使用？
输入一个行政区域代码比如 130104 返回一个省市县的数组[河北省 石家庄市 桥西区]
如果是直辖市比如北京 110000 返回就是[北京市]

````golang
// 引入包
import "github.com/ppmoon/gbt2260"
// 打印查看
region := gbt2260.NewGBT2260()
localCode := region.SearchGBT2260("130104")
fmt.Println(localCode)
````
````bash
[河北省 石家庄市 桥西区]
````
````golang
// 获取所有省份
allProvince := region.GetAllProvince()
fmt.Println(allProvince)
// 通过省份获得省份下的城市
allCity := gbt.GetCityByProvince("130000")
fmt.Println(allCity)
// 通过城市获得城市的区县
allArea := gbt.GetAreaByCity("130100")
fmt.Println(allArea)
````

## 数据来源

目前数据使用2022年1月的最新数据

[中华人民共和国民政部](https://www.mca.gov.cn/mzsj/xzqh/2022/202201xzqh.html)

民政部啥时候发布新数据还是挺没准的，看之前的数据是一年一发，现在也找不到什么规律。
不然考虑每次载入程序的时候都去民政部官方网站更新下数据。

## 其他

1.gbt2260的数据存储到了一个trie 字典树当中。这样的数据结构查询的时候更快。（虽然3000多条数据完全没感觉~~~）

2.数据现在可以通过 cmd/gengbt2260.go 生成了，调整好文件go run cmd/gengbt2260.go就可以了，后面会再优化下。

3.希望这个扩展包可以帮助大家解决行政区域的业务需求。我自己是用来解析身份证信息来使用的。如果觉得对你有帮助欢迎star谢谢。

大家使用的过程当中如果发现任何问题或者想吐槽我的代码可以提交issue。
数据需要更新的时候也麻烦大家提交issue，当然未来也希望这个包可以自动从民政部的网站上更新中华人民共和国县以上行政区划代码的数据。
