package utils

import "testing"

func TestStruct2map(t *testing.T) {
	cases := []struct {
		Username string   `json:"username,a"`
		Password string   `json:"b,password"`
		Age      int      `json:"age"`
		Tag      []string `json:"tag"`
	}{
		{
			Username: "张三",
			Password: "123456",
		},
		{
			Username: "张三",
			Age:      18,
			Tag:      []string{"1", "2", "3"},
		},
	}

	for _, v := range cases {
		vFalse := Struct2map(v, "json", false)
		vTrue := Struct2map(v, "json", true)
		t.Logf("%+v", vFalse)
		t.Logf("%+v", vTrue)
	}
}
