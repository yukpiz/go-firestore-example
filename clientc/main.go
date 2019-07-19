package main

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/yukpiz/go-firestore-example/domain"
	"golang.org/x/xerrors"
	"google.golang.org/api/option"
)

func main() {
	msg := flag.String("msg", "", "message")
	flag.Parse()
	if len(strings.TrimSpace(*msg)) == 0 {
		panic(xerrors.New("empty message[required -msg option]"))
	}
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
	ref, res, err := cli.Collection("messages").Add(ctx, &domain.Message{
		Message:  *msg,
		SenderID: "clientc",
		SentAt:   time.Now().Unix(),
	})
	log.Println(ref, res)
}
