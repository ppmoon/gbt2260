package gbt2260

import (
	"fmt"
	"sync"
)

type BGT2260 struct {
	trie *Trie
}

var gbt2260 *BGT2260
var once sync.Once

func NewGBT2260() *BGT2260 {
	once.Do(func() {
		gbt2260Table := GetGbt2260Table()
		t := NewTrie()
		for _, cell := range gbt2260Table {
			//检查传递参数
			if cell[0] == "" || len(cell[0]) == 0 {
				continue
			}
			//过滤下数据构造插入lCode
			var lCode = stringParse(cell[0])
			//创建trie树
			trieRoot := t
			trieRoot.Add(lCode, cell[1])
		}
		gbt2260 = &BGT2260{
			trie: t,
		}
	})
	return gbt2260
}

// 将传入的字符串解析成字符串数组
func stringParse(str string) []string {
	var lCode []string
	for i := 0; i < len(str)/2; i++ {
		if str[2*i:2*(i+1)] != "00" {
			lCode = append(lCode, str[2*i:2*(i+1)])
		}
	}
	return lCode
}

// 从树里面读取数据
func (b *BGT2260) SearchGBT2260(code string) []string {
	var lCode = stringParse(code)
	var newCode []string
	node := b.trie.Root()
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
func (b *BGT2260) GetAllProvince() map[string]string {
	var provinceList = make(map[string]string)
	node := b.trie.Root()
	for k, v := range node.Children() {
		provinceList[k+"0000"] = v.value
	}
	return provinceList
}

// 获取省份下的城市
func (b *BGT2260) GetCityByProvince(code string) map[string]string {
	var cityList = make(map[string]string)
	var lCode = stringParse(code)
	node := b.trie.Root()
	for pk, pv := range node.Children() {
		if lCode[0] == pk {
			for ck, cv := range pv.Children() {
				cityList[pk+ck+"00"] = cv.value
			}
		}
	}
	return cityList
}

// 获取城市下的区
func (b *BGT2260) GetAreaByCity(code string) map[string]string {
	var areaMap = make(map[string]string)
	var lCode = stringParse(code)
	node := b.trie.Root()
	for pk, pv := range node.Children() {
		if lCode[0] == pk {
			for ck, cv := range pv.Children() {
				if ck == lCode[1] {
					for ak, av := range cv.Children() {
						areaMap[pk+ck+ak] = av.value
					}
				}
			}
		}
	}
	return areaMap
}
