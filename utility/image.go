package utility

import (
	"github.com/disintegration/imaging"
)

// ResizeRequest はサムネイル作成の構造体
type ResizeRequest struct {
	// 元画像のパス
	OriginalPath string
	// 保存先のパス
	SaveTo string
	// リサイズ後の上限幅
	NewWidth int
	// リサイズ後の上限高さ
	NewHeight int
}

// Resize はリサイズしてファイルを指定のパスに保存
func Resize(req *ResizeRequest) error {

	// 画像を開く
	imgSrc, err := imaging.Open(req.OriginalPath)

	if err != nil {
		return err
	}

	bounds := imgSrc.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// リサイズ後のサイズ計算
	w, h := calcSize(width, height, req.NewWidth, req.NewHeight)

	// リサイズ
	dstSrc := imaging.Resize(imgSrc, w, h, imaging.Linear)

	if err := imaging.Save(dstSrc, req.SaveTo); err != nil {
		return err
	}

	return nil
}

// CropRectRequest は正方形切り抜きの構造
type CropRectRequest struct {
	// 元画像のパス
	OriginalPath string
	// 保存先のパス
	SaveTo string
	// リサイズ後の上限幅/高さ
	NewSize int
}

// CropRect はリサイズしてファイルを正方形にトリミング
func CropRect(req *CropRectRequest) error {

	// スマホの画像などに対応
	opts := imaging.AutoOrientation(true)

	// 画像を開く
	imgSrc, err := imaging.Open(req.OriginalPath, opts)

	if err != nil {
		return err
	}

	bounds := imgSrc.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 小さい方の辺に合わせてリサイズ
	newSize := width
	if height < width {
		newSize = height
	}

	// 切り抜き
	croppedSrc := imaging.CropAnchor(imgSrc, newSize, newSize, imaging.Center)

	// 切り抜き後のサイズを取得
	bounds = croppedSrc.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()

	// リサイズ後のサイズ計算
	w, h := calcSize(width, height, req.NewSize, req.NewSize)

	// リサイズ
	dstSrc := imaging.Resize(croppedSrc, w, h, imaging.Linear)

	if err := imaging.Save(dstSrc, req.SaveTo); err != nil {
		return err
	}

	return nil
}

// calcSize はサイズを計算
func calcSize(w, h, maxW, maxH int) (int, int) {

	if w > maxW {
		// 幅がサイズ超過している場合
		percentage := float64(maxW) / float64(w)
		w = int(float64(w) * percentage)
		h = int(float64(h) * percentage)
	}

	if h > maxH {
		// 高さがサイズ超過している場合
		percentage := float64(maxH) / float64(h)
		w = int(float64(w) * percentage)
		h = int(float64(h) * percentage)
	}

	return w, h
}
