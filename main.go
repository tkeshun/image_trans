package main
import (
  "flag"
  "fmt"
  "path/filepath"
  "image_trans/conversion"
  "image_trans/findFile"
)


var input_fmt *string
var output_fmt *string

func init(){
  
  input_fmt  = flag.String("i_fmt","jpeg","画像の入力フォーマットを指定") 
  output_fmt = flag.String("o_fmt","png","画像の出力フォーマットを指定") 
}

func main(){
	//フラグの読み込み
	flag.Parse()
	args := flag.Args()
	if len(args) > 2 {
		fmt.Println("Only one directory can be specified")
	}
	
    
	//ディレクトリ内の対象ファイル一覧を取得
	f_list := findFile.Search(flag.Arg(0),*input_fmt)
	
	//画像変換
	for _,srcPath :=range(f_list) {
		fmt.Println(srcPath," >> ",srcPath[:len(srcPath)-len(filepath.Ext(srcPath))] + "." + *output_fmt )
        conversion.Convert(srcPath,*output_fmt)
	}
    
}