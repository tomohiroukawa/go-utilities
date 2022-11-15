package utility

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/text/width"

	"github.com/google/uuid"
)

// CreateUUID はUUIDを返す
func CreateUUID() uuid.UUID {
	u, err := uuid.NewRandom()

	if err != nil {
		// 作れないことはないだろうけど、万が一の場合はpanic
		panic(err)
	}

	return u
}

// SplitWord は全角、半角問わずスペース区切りの文字列をスペースで分割してスライスを返す
func SplitWord(str string) []string {

	// 漢字・ひらがな・カタカナを全角に、英数字を半角に統一してスペースで分割
	words := strings.Split(width.Fold.String(str), " ")

	var res []string

	for _, w := range words {

		s := strings.TrimSpace(w)

		if s == "" {
			continue
		}

		res = append(res, s)
	}

	return res
}

// HashPassword はパスワードをハッシュ
func HashPassword(plain string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)

	if err != nil {
		// 作れないことはないだろうけど、万が一の場合はpanic
		panic(err)
	}

	return string(b)
}

// CheckPassword はパスワードをチェック
func CheckPassword(plain string, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)); err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

// GenerateToken はランダムなUUIDをベースにしたトークンの作成
func GenerateToken() string {
	random := CreateUUID()
	b := sha256.Sum256([]byte(random.String()))
	return fmt.Sprintf("%x", b)
}

// InArrayString はhaystackにneedleが存在するかチェック
func InArrayString(haystack []string, needle string) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// InArrayInt はhaystackにneedleが存在するかチェック
func InArrayInt(haystack []int, needle int) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// InArrayInt32 はhaystackにneedleが存在するかチェック
func InArrayInt32(haystack []int32, needle int32) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// InArrayInt64 はhaystackにneedleが存在するかチェック
func InArrayInt64(haystack []int64, needle int64) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// InArrayFloat64 はhaystackにneedleが存在するかチェック
func InArrayFloat64(haystack []float64, needle float64) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// GetExtension はファイル名から拡張子を取得（ドットは含まない）
func GetExtension(filename string) (string, error) {

	// 拒否ファイルリストのみの定義なので、オリジナルファイルの拡張子を取得するしかない。
	parts := strings.Split(filename, ".")

	if len(parts) <= 1 {
		// 拡張子がない場合エラー
		return "", fmt.Errorf("拡張子が特定できませんでした。")
	}

	// 小文字に統一
	return strings.ToLower(parts[len(parts)-1]), nil
}

// FileExists はファイルの存在チェック
func FileExists(filename string) bool {
	_, err := os.Stat(filename)

	if pathError, ok := err.(*os.PathError); ok {
		if pathError.Err == syscall.ENOTDIR {
			return false
		}
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// PtrString はstringをポインターに変更(aws.String()的な)
func PtrString(str string) *string {
	return &str
}
