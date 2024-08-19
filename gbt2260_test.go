package gbt2260_test

import (
	"testing"

	"github.com/ppmoon/gbt2260"
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
	if len(allProvince) != 34 {
		t.Errorf("省份数量不正确")
	}
	t.Log(allProvince)
}

func TestBGT2260_GetCityByProvince(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	allCity := gbt.GetCityByProvince("130000")
	if len(allCity) != 11 {
		t.Error("河北省地级市数量不对")
	}
	t.Log(allCity)
}

func TestBGT2260_GetAreaByCity(t *testing.T) {
	gbt := gbt2260.NewGBT2260()
	area := gbt.GetAreaByCity("130100")
	areaName := area["130102"]
	if areaName != "长安区" {
		t.Error("get area by city error")
	}
}

func TestNewGBT2260(t *testing.T) {
	a := gbt2260.NewGBT2260()
	b := gbt2260.NewGBT2260()
	if a != b {
		t.Error("singleton error")
		return
	}
}
