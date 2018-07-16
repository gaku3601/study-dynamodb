package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	Name        string    `dynamo:"Name"`
	CreatedTime time.Time `dynamo:"created_time"`
}

func main() {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})
	table := db.Table("Test")

	u := User{Name: "Gaku", CreatedTime: time.Now().UTC()}

	if err := table.Put(u).Run(); err != nil {
		panic(err.Error())
	}
}
