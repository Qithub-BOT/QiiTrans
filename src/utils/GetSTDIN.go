package utils

import (
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"
)

var (
	// ValSTDINDummy は GetSTDIN() の挙動を mock するための値です.
	//
	// この値が空ではない場合、そのセットされた値を返します.
	ValSTDINDummy string = ""

	// ForceErrorGetSTDIN はテスト用の変数です。true の場合、GetSTDIN は強制的にエラーを返します.
	//
	// この設定はテストで強制的にエラーを発生したい場合に利用されます.
	ForceErrorGetSTDIN bool = false
)

// GetSTDIN はパイプ渡し or リダイレクトされた標準入力からのデータを返します.
//
// 対話式で標準入力を取得したい場合は InteractSTDIN() を利用してください.
func GetSTDIN() (stdin string, err error) {
	if ForceErrorGetSTDIN {
		return "", xerrors.New("forced to return error")
	}

	if ValSTDINDummy != "" {
		return ValSTDINDummy, nil
	}

	value, err := ioutil.ReadAll(os.Stdin)

	return string(value), err
}
