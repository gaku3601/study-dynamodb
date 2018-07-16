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
	URLs        []string
	Children    []Children
}

type Children struct {
	Name string
	Age  int
}

func main() {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"), // "ap-northeast-1"等
	})
	table := db.Table("Test")

	u := User{Name: "Gaku", CreatedTime: time.Now().UTC(), URLs: []string{
		"http://gakusmemo1.com/",
		"http://www.gakusmemo.com/",
		"http://www.gakusmemo1.com/",
		"http://www.gakusmemo2.com/",
		"http://www.gakusmemo2.com/",
		"http://www.gakusmemo2.com/",
	},
		Children: []Children{
			{Name: "gaku", Age: 25},
			{Name: "gakuko", Age: 50},
			{Name: "gaku0", Age: 51},
			{Name: "gaku1", Age: 52},
			{Name: "gaku2", Age: 53},
			{Name: "gaku3", Age: 54},
		},
	}

	if err := table.Put(u).Run(); err != nil {
		panic(err.Error())
	}
}
