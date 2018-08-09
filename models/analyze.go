package models

import (
	"fmt"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	"math"
)

var Result *AnalyseResult
var Score float64
var Magnitude float64

// 絵文字配列
var EmojiBoard = [][]string{
	[]string{"\xF0\x9F\x98\xA4", "\xF0\x9F\x98\x92", "\xF0\x9F\x98\x8C", "\xF0\x9F\x98\x8A", "\xF0\x9F\x98\x83"},
	[]string{"\xF0\x9F\x98\xAB", "\xF0\x9F\x98\x94", "\xF0\x9F\x98\x8F", "\xF0\x9F\x98\x8B", "\xF0\x9F\x98\x84"},
	[]string{"\xF0\x9F\x98\xAD", "\xF0\x9F\x98\x93", "\xF0\x9F\x98\x8A", "\xF0\x9F\x98\x8D", "\xF0\x9F\x98\x85"},
}

// 結果保存用
type AnalyseResult []struct {
	Text      string
	Emoji     string
	Score     float64
	Magnitude float64
}

func init() {
	Result = &AnalyseResult{}
}

func setResult(score float64) {
	Score = score
}

func setMagnitude(magnitude float64) {
	Magnitude = magnitude
}

// APIからデータ取得、処理、保存
func ProcessText(text string, res *languagepb.AnalyzeSentimentResponse) {
	SetText(text)

	m := float64(res.DocumentSentiment.Magnitude)
	s := float64(res.DocumentSentiment.Score)

	for _, elem := range res.Sentences {
		fmt.Printf("Content : %d\n", elem.Sentiment.Score)
	}
	// Magnitute, score に当て嵌まる絵文字の取得
	emoji := ConverseToEmoji(m, s)

	SetMagnitute(m)
	SetScore(s)
	SetEmoji(emoji)
}

// Magnitute, score に当て嵌まる絵文字の取得
func ConverseToEmoji(m float64, s float64) string {

	// flatten magnitude
	var magnitude int = int(m)
	if magnitude > 2 {
		magnitude = 2
	}

	// round score
	score := Round(s)

	fmt.Printf("score : %d\n", score)
	fmt.Printf("magnitude : %d\n", magnitude)
	fmt.Printf("length1 : %d\n", len(EmojiBoard))

	return EmojiBoard[magnitude][score]
}

// Scoreのデータを0.5彫みに分ける
func Round(x float64) int {
	roundUp := (math.Round(x * 2)) / 2
	roundUp = (roundUp + 1) * 2
	var result = int(roundUp)
	return result
}
