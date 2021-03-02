package conversion

import (
	"fmt"
	"os"
	"log"
	"path/filepath"
	"image"
	_ "image/jpeg"
	"image/png"
)
func assert(err error ,msg string){
    if err != nil {
		// エラー時の処理
		println(msg)
		log.Fatal(err)
	}
}

func Convert(srcPath string, output_fmt string){
	//画像のフォーマットによって変換方法を変える
	//ファイル読み込み
	// ファイルオープン
	file, err := os.Open(srcPath)
	assert(err, "Invalid image file path " + srcPath)
	defer file.Close()

	// ファイルオブジェクトを画像オブジェクトに変換
	img, _, err :=  image.Decode(file)


	//出力ファイル生成
	output_path := srcPath[:len(srcPath)-len(filepath.Ext(srcPath))] + "." + output_fmt 
    out, err := os.Create(output_path)
	assert(err,"ファイル作成時")
	defer out.Close()
	//ファイル出力(変換)

	switch filepath.Ext(output_path) {
	case ".png":
	    png.Encode(out,img)
	default :
		fmt.Println("指定した拡張子がまちがっている可能性があります")
		os.Exit(1)
	}
}

