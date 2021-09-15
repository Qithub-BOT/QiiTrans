package app_test

import (
	"testing"

	"github.com/Qithub-BOT/QiiTrans/src/app"
	"github.com/Qithub-BOT/QiiTrans/src/helperfunc"
	"github.com/Qithub-BOT/QiiTrans/src/utils"
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInteractiveTranslation(t *testing.T) {
	stopWord := "stop"
	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"
	dummyArgs := []string{"en", "ja", "es"}

	// 標準入力をモック
	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer()

	// 引数をモック
	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true

	defer func() { utils.IsTerminalDummy = false }()

	appTest := app.New(t.Name())

	appTest.StopWord = stopWord
	appTest.Force["NoTrans"] = true

	// テスト実行
	out := capturer.CaptureOutput(func() {
		err := appTest.InteractiveTranslation(dummyArgs)

		require.NoError(t, err)
	})

	assert.Contains(t, out, stopWord)
	assert.Contains(t, out, appTest.Prefix)
	assert.Contains(t, out, "foo bar\n")
	assert.Contains(t, out, "hoge fuga\n")
}

func TestInteractiveTranslation_fail(t *testing.T) {
	stopWord := "stop"
	userInput := "foo bar\nhoge fuga\n" + stopWord + "\n"
	dummyArgs := []string{"en", "ja", "es"}

	// 標準入力をモック
	funcDefer := helperfunc.MockSTDIN(t, userInput)
	defer funcDefer()

	// 引数をモック
	funcDeferArgs := helperfunc.MockArgs(t, dummyArgs)
	defer funcDeferArgs()

	// utils.IsTerminal のモックとリカバリー
	utils.IsTerminalDummy = true

	defer func() { utils.IsTerminalDummy = false }()

	appTest := app.New(t.Name())

	appTest.StopWord = stopWord
	appTest.Force["TransError"] = true

	// テスト実行
	err := appTest.InteractiveTranslation(dummyArgs)

	require.Error(t, err)
}
