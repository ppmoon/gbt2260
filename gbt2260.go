package gbt2260

import (
	"fmt"
)

type BGT2260 struct{}

//通过全局的方式创建trie树
var trie = New()

func NewGBT2260() *BGT2260 {
	gbt2260Table := GetGbt2260Table()
	for _, cell := range gbt2260Table {
		createTrieTree(cell[0], cell[1], trie)
	}
	return &BGT2260{}
}

//向树中插入数据
func createTrieTree(code string, name string, trie *Trie) {
	//检查传递参数
	if code == "" || len(code) == 0 {
		return
	}
	//过滤下数据构造插入lCode
	var lCode = stringParse(code)
	//创建trie树
	trieRoot := trie
	trieRoot.Add(lCode, name)
}

//将传入的字符串解析成字符串数组
func stringParse(str string) []string {
	var lCode []string
	for i := 0; i < len(str)/2; i++ {
		if str[2*i:2*(i+1)] != "00" {
			lCode = append(lCode, str[2*i:2*(i+1)])
		}
	}
	return lCode
}

//从树里面读取数据
func (b *BGT2260) SearchGBT2260(code string) []string {
	var lCode = stringParse(code)
	var newCode = []string{}
	node := trie.Root()
	for i := range lCode {
		r := lCode[i]
		if n, ok := node.children[r]; ok {
			newCode = append(newCode, n.value)
			node = n
		} else {
			fmt.Printf("对不起，您输入的地域码不在列表当中")
		}
	}
	return newCode
}

// 获取所有省份
func (b *BGT2260) GetAllProvince() map[string]string{
	var provinceList = make(map[string]string)
	node := trie.Root()
	for k,v:=range node.Children(){
		provinceList[k + "0000"] = v.value
	}
	return provinceList
}

// 获取省份下的城市
func (b *BGT2260) GetCitysByProvince(code string) map[string]string{
	var cityList = make(map[string]string)
	var lCode = stringParse(code)
	node := trie.Root()
	for k,v:=range node.Children(){
		if lCode[0] == k{
			for kk,vv:=range v.Children(){
				cityList[k+kk + "00"] = vv.value
			}
		}
	}
	return cityList
}

// 获取城市下的区
func (b *BGT2260) GetAreaByCity(code string) map[string]string{
	var areaList = make(map[string]string)
	var lCode = stringParse(code)
	node := trie.Root()
	for k,v:=range node.Children(){
		if lCode[0] == k{
			for kk,_:=range v.Children(){
				if kk == lCode[1] {
					for kkk,vvv:=range v.Children(){
						areaList[k+kk+kkk] = vvv.value
					}
				}
			}
		}
	}
	return areaList
}
