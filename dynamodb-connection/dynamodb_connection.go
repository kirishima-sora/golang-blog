package main

import (
    "fmt"
    "net/http"
	"encoding/json"
	// Go用のLambdaプログラミングモデル
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/guregu/dynamo"
)

type user struct {
    UserName   string `json:"username"`
    Birthday  string `json:"birthday"`
}

// メソッドに応じて処理分岐
func DBOperateAPI(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    db := dynamo.New(session.New(), &aws.Config{Region: aws.String("ap-northeast-1")})
    table := db.Table("UserTable")
	switch req.HTTPMethod {
    case "GET":
		username := req.QueryStringParameters["UserName"]
        var result Event
        err := table.Get("UserName", username).One(&result)
        

    // case "POST":
    default:
        return clientError(http.StatusMethodNotAllowed)
    }
}

func main() {
    lambda.Start(DBOperateAPI)
}