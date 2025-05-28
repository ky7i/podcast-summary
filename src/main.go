package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(".envファイルの読み込みに失敗しました。")
	}

	CallGeminiAPI()
}
