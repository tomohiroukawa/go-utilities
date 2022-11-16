package utility

import (
	"reflect"
	"testing"
)

func TestCreateUUID(t *testing.T) {

	res := CreateUUID().String()

	if len(res) != 36 {
		t.Fatalf("length is not valid %d", len(res))
	}
}

func TestSplitWord(t *testing.T) {

	str := "ワード1 ワード2  ワード3　　　 　  ワード4"

	res := SplitWord(str)

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

func TestInArrayString(t *testing.T) {

	arr := []string{"test1", "test2", "test3", "日本語"}

	if InArrayString(arr, "ない") {
		t.Fatalf("result shoud be false")
	}

	if !InArrayString(arr, "test2") {
		t.Fatalf("result shoud be true")
	}

	if !InArrayString(arr, "日本語") {
		t.Fatalf("result shoud be true")
	}
}

func TestInArrayInt(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 0}

	if InArrayInt(arr, 6) {
		t.Fatalf("result shoud be false")
	}

	if !InArrayInt(arr, 1) {
		t.Fatalf("result shoud be true")
	}

	if !InArrayInt(arr, 0) {
		t.Fatalf("result shoud be true")
	}
}

func TestInArrayInt32(t *testing.T) {

	arr := []int32{1874623, 298355872, 3875367, 489764597, 57857863, 947262}

	if InArrayInt32(arr, 8747862) {
		t.Fatalf("result shoud be false")
	}

	if !InArrayInt32(arr, 1874623) {
		t.Fatalf("result shoud be true")
	}

	if !InArrayInt32(arr, 489764597) {
		t.Fatalf("result shoud be true")
	}
}

func TestInArrayInt64(t *testing.T) {

	arr := []int64{18756362672776, 275635352542, 387578278287, 47895978627938, 5578783782, 868934223445353}

	if InArrayInt64(arr, 557478928783782) {
		t.Fatalf("result shoud be false")
	}

	if !InArrayInt64(arr, 387578278287) {
		t.Fatalf("result shoud be true")
	}

	if !InArrayInt64(arr, 868934223445353) {
		t.Fatalf("result shoud be true")
	}
}

func TestInArrayFloat64(t *testing.T) {

	arr := []float64{1.020292893, 2.3875632, 3.3856271, 4123.3756352, 545.4873627462, 36.383726}

	if InArrayFloat64(arr, 64.28278) {
		t.Fatalf("result shoud be false")
	}

	if !InArrayFloat64(arr, 2.3875632) {
		t.Fatalf("result shoud be true")
	}

	if !InArrayFloat64(arr, 545.4873627462) {
		t.Fatalf("result shoud be true")
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

func TestPtrString(t *testing.T) {

	str := "not pointer"

	res := PtrString(str)

	typeOf := reflect.TypeOf(res)

	if typeOf.String() != "*string" {
		t.Fatalf("result should be *string")
	}
}
