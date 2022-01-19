package app

import (
	"strings"

	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"golang.org/x/xerrors"
)

// Translate は、orderLang の順に inputText を翻訳した結果を返します.
func (a *TApp) Translate(orderLang []string, inputText string) ([]TTranslation, error) {
	var err error

	translations := make([]TTranslation, len(orderLang)-1)

	if a.Force["TransError"] {
		return nil, xerrors.New("forced error for translation")
	}

	transText := ""
	langFrom := ""
	langTo := ""

	// orderLang 順に再帰的に翻訳
	for i, nameLang := range orderLang {
		nameLang = strings.ToUpper(nameLang) // API に合わせて言語を大文字に統一

		// 最初の言語をセット
		if i == 0 {
			langFrom = nameLang
			transText = inputText

			continue
		}

		langTo = nameLang

		obj := NewTranslation(langFrom, langTo, transText)

		if !a.Force["NoTrans"] {
			// from -> to へ翻訳
			transText, _, err = a.Engine.Translate(transText, langFrom, langTo)
			if err != nil {
				return nil, err
			}

			utils.LogDebug("Translate 結果:", transText)
		}

		obj.Translated = transText // Set translated text
		translations[i-1] = obj    // Set tranlated object
		langFrom = langTo          // Update lang
	}

	return translations, err
}
