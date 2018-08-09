package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
	"os"

	// [START imports]
	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	// [END imports]

	"github.com/astaxie/beego"
	"testProject/models"
)

type TestNaturalLanguageController struct {
	beego.Controller
}

func (this *TestNaturalLanguageController) Prepare() {

}

func (t *TestNaturalLanguageController) GetResult() {
	res := struct{ Result *models.AnalyseResult }{models.Result}
	fmt.Fprintln(os.Stderr, "GetResult : ", res.Result.Text)
	t.Data["json"] = res
	t.ServeJSON()
}

func (t *TestNaturalLanguageController) ProcessText() {

	fmt.Fprintln(os.Stderr, "ProcessText")

	req := struct{ Analyzetext string }{}
	if err := json.Unmarshal(t.Ctx.Input.RequestBody, &req); err != nil {
		t.Ctx.Output.SetStatus(400)
		t.Ctx.Output.Body([]byte("empty text"))
		fmt.Fprintln(os.Stderr, "empty text")
		return
	}

	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	text := req.Analyzetext
	if text == "" {
		fmt.Fprintln(os.Stderr, "empty text")
		return
	}

	fmt.Fprintln(os.Stderr, "ProcessText :", text)

	// APIを叩く
	res, err := analyzeSentiment(ctx, client, text)
	if err != nil {
		log.Fatalf("got %v, want nil err", err)
	}

	fmt.Fprintln(os.Stderr, "ProcessText", res)

	// APIからもらった結果を処理、保存する
	models.ProcessText(text, res)
}

func usage(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Fprintln(os.Stderr, "usage: analyze [entities|sentiment|syntax|entitysentiment|classify] <text>")
	os.Exit(2)
}

// テキストを渡してGOOGLE APIを叩く
func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
	return client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type:     languagepb.Document_PLAIN_TEXT,
			Language: "ja",
		},
	})
}

func printResp(v proto.Message, err error) {
	if err != nil {
		log.Fatal(err)
	}
	proto.MarshalText(os.Stdout, v)
}
