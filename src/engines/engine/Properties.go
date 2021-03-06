package engine

import (
	"github.com/Qithub-BOT/QiiTrans/src/cache"
)

// Properties は各翻訳エンジンの基本となる構造体です.
// 翻訳エンジンが API を呼び出すのに必要な基本情報と、翻訳に使われるメソッド（クラス関数）を持った構造体です.
type Properties struct {
	// 翻訳済みのテキストをキャッシュするためのオブジェクトのポインタです
	Cache *cache.TCache

	// translate メソッドは、翻訳エンジンがインスタンス作成時に割り当てた関数です.
	// inputText を langFrom から langTo  翻訳した結果を返します。この値はキャッシュされません.
	translate func(properties *Properties, inputText string, langFrom string, langTo string) (string, error)

	// getInfoAPI メソッドは、翻訳エンジンがインスタンス作成時に割り当てた関数です.
	// アクセス・トークンのアカウント情報のうち、AccountInfo 型の情報だけ返します.
	getInfoAPI func(properties *Properties) (AccountInfo, error)

	// NameEngine は利用している翻訳エンジン名です.
	NameEngine string

	// 環境変数からアクセス・トークンを取得する際の環境変数名
	NameVarEnvAPIKey string

	// 自動翻訳 API 用のアクセス・トークン（API キー）です。空の場合は環境変数から取得します。
	apiKey string

	// 1 リクエストごとに sleep させる秒数です（デフォルト: 1）
	TimeInterval int

	// アクセス・トークンが無料アカウントの場合は true、有料アカウントの場合は false にセットします。（デフォルト: true）
	IsAccountFree bool

	// true の場合はキャッシュを更新します。false の場合はなるべくキャッシュを使います（デフォルト: false）
	Update bool
}
