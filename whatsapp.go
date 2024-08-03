package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func setupWhatsAppClient() {
	dbLog := waLog.Stdout("Database", "ERROR", true)
	container, err := sqlstore.New("sqlite3", "file:whatsapp.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "ERROR", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}
	client.SendPresence(types.PresenceAvailable)
}

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Presence:
		jid := v.From.String()
		status := "Online"
		if v.Unavailable {
			status = "Offline"
		}

		mu.Lock()
		userStatus[jid] = status
		userStatusLog[jid] = append(userStatusLog[jid], StatusLog{
			Time:   time.Now().UTC(),
			Status: status,
		})
		mu.Unlock()
	}
}
