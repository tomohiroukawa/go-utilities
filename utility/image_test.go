package utility

import (
	"testing"

	"github.com/disintegration/imaging"
)

func TestResize(t *testing.T) {

	// 横向きテスト
	req := &ResizeRequest{
		OriginalPath: "../test_data/test.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewWidth:     100,
		NewHeight:    10,
	}

	if err := Resize(req); err != nil {
		t.Fatalf("error: %v", err)
	}

	// サイズチェック
	imgSrc, err := imaging.Open(req.SaveTo)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	w := imgSrc.Bounds().Dx()
	h := imgSrc.Bounds().Dy()

	if w != 15 && h != 10 {
		t.Fatalf("result size is not 210x100. Actually %dx%d", w, h)
	}

	// 縦向きテスト
	req2 := &ResizeRequest{
		OriginalPath: "../test_data/test_rotate.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewWidth:     100,
		NewHeight:    10,
	}

	if err := Resize(req); err != nil {
		t.Fatalf("error: %v", err)
	}

	// サイズチェック
	imgSrc, err = imaging.Open(req2.SaveTo)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	w = imgSrc.Bounds().Dx()
	h = imgSrc.Bounds().Dy()

	if w != 15 && h != 10 {
		t.Fatalf("result size is not 210x100. Actually %dx%d", w, h)
	}

	// 存在しない1
	req3 := &ResizeRequest{
		OriginalPath: "../test_data/not_exist.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewWidth:     100,
		NewHeight:    10,
	}

	if err := Resize(req3); err == nil {
		t.Fatalf("error shoud be occured")
	}

	// 存在しない2
	req4 := &ResizeRequest{
		OriginalPath: "../test_data/test.jpg",
		SaveTo:       "/not_exist.jpg",
		NewWidth:     100,
		NewHeight:    10,
	}

	if err := Resize(req4); err == nil {
		t.Fatalf("error shoud be occured")
	}

}

func TestCropRect(t *testing.T) {

	// 縦テスト
	req := &CropRectRequest{
		OriginalPath: "../test_data/test_rotate.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewSize:      80,
	}

	if err := CropRect(req); err != nil {
		t.Fatalf("error: %v", err)
	}

	// サイズチェック
	imgSrc, err := imaging.Open(req.SaveTo)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	w := imgSrc.Bounds().Dx()
	h := imgSrc.Bounds().Dy()

	if w != 80 && h != 80 {
		t.Fatalf("result size is not 210x100. Actually %dx%d", w, h)
	}

	// 横テスト
	req2 := &CropRectRequest{
		OriginalPath: "../test_data/test.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewSize:      80,
	}

	if err := CropRect(req2); err != nil {
		t.Fatalf("error: %v", err)
	}

	// サイズチェック
	imgSrc, err = imaging.Open(req2.SaveTo)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	w = imgSrc.Bounds().Dx()
	h = imgSrc.Bounds().Dy()

	if w != 80 && h != 80 {
		t.Fatalf("result size is not 210x100. Actually %dx%d", w, h)
	}

	// 存在しない1
	req3 := &CropRectRequest{
		OriginalPath: "../test_data/test_not_exists.jpg",
		SaveTo:       "../test_data/test_result.jpg",
		NewSize:      80,
	}

	if err := CropRect(req3); err == nil {
		t.Fatalf("error shoud be occured")
	}

	// 存在しない2
	req4 := &CropRectRequest{
		OriginalPath: "../test_data/test.jpg",
		SaveTo:       "/cannot_save.jpg",
		NewSize:      80,
	}

	if err := CropRect(req4); err == nil {
		t.Fatalf("error shoud be occured")
	}

}
