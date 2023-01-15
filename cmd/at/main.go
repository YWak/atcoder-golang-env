package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}

	var err *error
	switch os.Args[1] {
	case "login":
		// ログイン情報を~/.config/at/configから取得する
		// ログインリクエストを行い、cookieを~/.config/at/cookieに保存する
		break

	case "init":
		// 指定されたURLから情報を取得し、パースしてex.txtとans.txtを生成する
		// templateに設定されたファイルを複製して保存する
		// 最後に処理を実行したディレクトリを~/.config/at/lastに保存する
		break

	case "test":
		// 指定されたディレクトリでテスト実行を行う
		break

	default:
		*err = help()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error on %s: %s\n", os.Args[1], *err)
	}
}

func help() error {
	puts := func(s string) {
		os.Stderr.WriteString(s)
		os.Stderr.WriteString("\n")
	}

	puts("usage: at operation arguments...")
	puts("")
	puts("operations")
	puts("  login    login to atcoder")
	puts("")
	puts("  init     create directory from url")
	puts("    at init https://atcoder.jp/contests/abc285/tasks/abc285_a")
	puts("")
	puts("  test     runs code using all test cases.")
	puts("           if no arguments are provided, recently initialized task will be tested")
	puts("    at test directory")
	return nil
}
