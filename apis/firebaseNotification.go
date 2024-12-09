package apis

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

func NotifyFirebase(tokens []string, title string, body string) error {
	app := initializeFirebaseApp()
	ctx := context.Background()

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Ошибка создания клиента Messaging: %v", err)
		return err
	}

	// Формируем сообщение
	message := &messaging.MulticastMessage{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Tokens: tokens,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	// Отправляем сообщение
	response, err := client.SendEachForMulticast(ctx, message)
	if err != nil {
		log.Fatalf("Ошибка отправки уведомления: %v", err)
	}

	for _, r := range response.Responses {
		fmt.Println(r.Error)
	}
	return nil
}

func initializeFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("apis/frist-firebase-project-cfe29-firebase-adminsdk-vp6tp-f160840574.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	return app
}
