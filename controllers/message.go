package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"go-wa-bot-v3/utils"
)


func SendPeriodicMessage(client *whatsmeow.Client, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		phone := os.Getenv("PHONE_ID")
		jid := types.NewJID(phone, "g.us")
		data, err := GetData()
		if err != nil {
			fmt.Println("Error fetching users:", err)
			continue
		}
		message := utils.FormatUsers(data)
		msg := &waProto.Message{
			Conversation: proto.String(message)}
		_, err = client.SendMessage(context.Background(), types.JID(jid), msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		} else {
			fmt.Println("Message sent successfully:", message)
		}
	}
}