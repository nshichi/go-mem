package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s dir\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	dir := os.Args[1]

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		ext := filepath.Ext(filename)
		if ext != ".ORF" && ext != ".JPG" {
			continue
		}

		mmdd, err := parseFilename(filename)
		if err != nil {
			// fmt.Printf("[スキップ] %s: ファイル名解析エラー: %v\n", filename, err)
			continue
		}

		subdir := filepath.Join(dir, mmdd)
		if err := os.MkdirAll(subdir, os.ModePerm); err != nil {
			log.Printf("[エラー] ディレクトリ作成失敗 (%s): %v\n", subdir, err)
			continue
		}

		old := filepath.Join(dir, filename)
		new := filepath.Join(subdir, filename)
		if err := os.Rename(old, new); err != nil {
			log.Printf("[エラー] ファイル移動失敗 (%s -> %s): %v\n", old, new, err)
			continue
		}

		count++
	}
}

// "PA050058.ORF" -> "1005"
func parseFilename(filename string) (string, error) {
	if len(filename) < 6 {
		return "", fmt.Errorf("filename too short: %s", filename)
	}

	month, err := strconv.ParseInt(filename[1:2], 16, 32)
	if err != nil {
		return "", err
	}

	day, err := strconv.ParseInt(filename[2:4], 10, 32)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%02d%02d", month, day), nil
}

/*
golang
画像ファイルの分別ツール

コマンドライン引数に画像ファイルを含んだディレクトリ名

ファイル名のサンプル "PA050058.ORF"
拡張子は .ORF と .JPG がありうる。

ファイル名の二文字めは月、ただし A は 10月、B は 11月、C は 12月。
ファイル名の３文字目から四文字は日。
mmdd形式のサブディレクトリを作り、そこに画像ファイルを移動させる。

対象外ファイルは残留。
*/
