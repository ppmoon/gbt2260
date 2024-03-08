package gbt2260_test

import (
	"github.com/ppmoon/gbt2260"
	"testing"
)

func TestBGT2260_SearchGBT2260(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	lCode := gbt.SearchGBT2260("130104")
	compareCode := []string{"河北省", "石家庄市", "桥西区"}
	for i := range lCode {
		if lCode[i] != compareCode[i] {
			t.Errorf("测试错误，地域码更改或不存在导致的程序错误")
			t.Fail()
		}
	}
}

func TestBGT2260_GetAllProvince(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	allProvince := gbt.GetAllProvince()
	t.Log(allProvince)
}

func TestBGT2260_GetCityByProvince(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	allCity := gbt.GetCityByProvince("110100")
	t.Log(allCity)
}

func TestBGT2260_GetAreaByCity(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	area := gbt.GetAreaByCity("110100")
	t.Log(area)
}

func TestNewGBT2260(t *testing.T) {
	a := gbt2260.NewGBT2260()
	b := gbt2260.NewGBT2260()
	if a != b {
		t.Error("singleton error")
		return
	}
}
