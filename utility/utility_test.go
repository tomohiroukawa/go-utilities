package utility

import (
	"fmt"
	"testing"
)

func TestCreateUUID(t *testing.T) {

	res := CreateUUID().String()

	if len(res) != 36 {
		t.Fatalf("length is not valid %d", len(res))
	}
}

func TestSplitWord(t *testing.T) {

	str := "ワード1 ワード2  ワード3　　　ワード４ABC"

	res := SplitWord(str)

	fmt.Println(res)

	if len(res) != 4 {
		t.Fatalf("length is not valid %d", len(res))
	}
}

func TestHashPassword(t *testing.T) {

	plain := "test_password"

	res := HashPassword(plain)

	if len(res) != 60 {
		t.Fatalf("length is not valid %d", len(res))
	}

	ok := CheckPassword(plain, res)

	if !ok {
		t.Fatalf("result shoud be true")
	}

	ok = CheckPassword(plain+"invalid", res)

	if ok {
		t.Fatalf("result shoud not be false")
	}
}

func TestGenerateToken(t *testing.T) {

	res := GenerateToken()

	if len(res) != 64 {
		t.Fatalf("length is not valid %d", len(res))
	}
}

func TestGetExtension(t *testing.T) {

	ext, err := GetExtension("/path/to/test.JPEG")

	if err != nil {
		t.Fatalf("err should be nil %s", err.Error())
	}

	if ext != "jpeg" {
		t.Fatalf("result is not valid")
	}

	ext, err = GetExtension("/p.a.t.h/to/test.bak.png")

	if err != nil {
		t.Fatalf("err should be nil %s", err.Error())
	}

	if ext != "png" {
		t.Fatalf("result is not valid")
	}

	_, err = GetExtension("/path/to/no_extension")

	if err == nil {
		t.Fatalf("err should not be nil")
	}

}

func TestFileExists(t *testing.T) {

	res := FileExists("./not-exist.txt")

	if res {
		t.Fatalf("result should be false")
	}

	res = FileExists("./utility.go")

	if !res {
		t.Fatalf("result should be true")
	}

	res = FileExists("../utility")

	if !res {
		t.Fatalf("result should be false")
	}
}

func TestPtr(t *testing.T) {

	v := 1
	r := ToPtr(v)

	if fmt.Sprintf("%T", r) != "*int" {
		t.Fatalf("result should be *int")
	}

	v2 := 1.90472
	r2 := ToPtr(v2)

	if fmt.Sprintf("%T", r2) != "*float64" {
		t.Fatalf("result should be *float64")
	}

	v3 := "hoge hoge"
	r3 := ToPtr(v3)

	if fmt.Sprintf("%T", r3) != "*string" {
		t.Fatalf("result should be *string")
	}

}

func TestInArray(t *testing.T) {
	if !InArray([]string{"A", "B"}, "B") {
		t.Fatalf("result should be true")
	}
	if InArray([]string{"A", "B"}, "C") {
		t.Fatalf("result should be false")
	}
	if !InArray([]int64{123, 456}, 123) {
		t.Fatalf("result should be true")
	}
	if !InArray([]float64{12.3342, 7128.35325, 456.432}, 7128.35325) {
		t.Fatalf("result should be true")
	}
}
