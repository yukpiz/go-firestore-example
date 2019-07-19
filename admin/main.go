package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("yukpiz-labo.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(err)
	}

	cli, err := app.Firestore(ctx)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	q := cli.Collection("messages").Where("SenderID", "==", "clienta")
	docs, err := q.Documents(ctx).GetAll()
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		log.Printf("%+v\n", doc.Data())
	}
}
