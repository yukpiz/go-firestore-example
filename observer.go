package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
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

	q := cli.Collection("messages").OrderBy("SentAt", firestore.Asc)
	qsiter := q.Snapshots(ctx)
	for {
		qsnap, err := qsiter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}

		for _, change := range qsnap.Changes {
			if change.Kind != firestore.DocumentAdded {
				continue
			}

			d := change.Doc.Data()
			fmt.Printf("%s: %s\n", d["SenderID"], d["Message"])
		}

	}
}
