package gbt2260

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
)
type BGT2260 struct {}
//通过全局的方式创建trie树
var trie = New()

func NewGBT2260() *BGT2260 {
	//读取文件
	file,err := os.Open("GBT2260-201802.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break //文件读完了就结束
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		code := record[0]
		name := record[1]
		createTrieTree(code,name,trie)
	}
	return &BGT2260{}
}
//向树中插入数据
func createTrieTree(code string,name string,trie *Trie) {
	//检查传递参数
	if code == "" || len(code) == 0 {
		return
	}
	//过滤下数据构造插入lCode
	var lCode = stringParse(code)
	//创建trie树
	trieRoot := trie
	trieRoot.Add(lCode,name)
}
//将传入的字符串解析成字符串数组
func stringParse (str string) []string{
	var lCode []string
	for i := 0;i<len(str)/2 ;i++  {
		if str[2*i:2*(i+1)] != "00" {
			lCode = append(lCode,str[2*i:2*(i+1)] )
		}
	}
	return lCode
}
//从树里面读取数据
func (b *BGT2260) SearchGBT2260 (code string) []string {
	var lCode = stringParse(code)
	var newCode = []string{}
	node := trie.Root()
	for i := range lCode{
		r := lCode[i]
		if n,ok := node.children[r]; ok {
			newCode = append(newCode,n.value)
			node = n
		} else {
			fmt.Printf("对不起，您输入的地域码不在列表当中")
		}
	}
	return newCode
}