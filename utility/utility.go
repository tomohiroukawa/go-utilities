package utility

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ktnyt/go-moji"

	"golang.org/x/crypto/bcrypt"

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

	// 全角スペースを半角スペースに変換して半角スペースでsplit
	words := strings.Split(moji.Convert(str, moji.ZS, moji.HS), " ")

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

// InArray は配列存在チェックを、ジェネリクスで実装したもの
func InArray[T comparable](haystack []T, needle T) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}

	return false
}

// GetExtension はファイル名から拡張子を取得（ドットは含まない）
func GetExtension(filename string) (string, error) {

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

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// ToPtr は引数のアドレスを返す
// ジェネリクスで汎用化
func ToPtr[T comparable](x T) *T {
	return &x
}
