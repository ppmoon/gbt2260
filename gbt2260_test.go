package gbt2260

import "testing"

//func TestSearchTrieTree(t *testing.T) {
//	local := SearchTrieTree("110000")
//	fmt.Println(local)
//}
func TestBGT2260_SearchGBT2260(t *testing.T) {
	gbt2260 := NewGBT2260()
	lCode := gbt2260.SearchGBT2260("130104")
	compareCode := []string{"河北省", "石家庄市", "桥西区"}
	for i := range lCode {
		if lCode[i] != compareCode[i] {
			t.Errorf("测试错误，地域码更改或不存在导致的程序错误")
			t.Fail()
		}
	}
}

//下面这个测试用来生成地域配置文件用的
func TestCreateGBT2260Table(t *testing.T) {
	CreateGBT2260Table()
}
