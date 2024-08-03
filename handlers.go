package main

import (
	"encoding/json"
	"go.mau.fi/whatsmeow/types"
	"html/template"
	"net/http"
	"time"
)

func selectContactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		selectedJIDs := r.Form["contacts"]
		for _, jid := range selectedJIDs {
			parsedJID, _ := types.ParseJID(jid)
			client.SubscribePresence(parsedJID)
		}
		time.Sleep(1 * time.Second)

		http.Redirect(w, r, "/status", http.StatusSeeOther)
		return
	}
	contacts, _ := client.Store.Contacts.GetAllContacts()
	for jid, contact := range contacts {
		if len(contact.FullName) > 0 {
			userNames[jid.String()] = contact.FullName
		} else {
			userNames[jid.String()] = contact.PushName
		}

	}
	tmpl := template.Must(template.ParseFiles("select_contacts.html"))
	tmpl.Execute(w, contacts)
}

func statusUpdateHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	updates := []StatusUpdate{}
	for jid, logs := range userStatusLog {
		onlineRanges := calculateOnlineRanges(logs)
		isOnline := false
		if len(logs) > 0 {
			lastLog := logs[len(logs)-1]
			isOnline = lastLog.Status == "Online"
		}
		username := userNames[jid]
		updates = append(updates, StatusUpdate{
			JID:          jid,
			Username:     username,
			OnlineRanges: onlineRanges,
			IsOnline:     isOnline,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updates)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	statuses := make(map[string]ContactData)
	for jid, status := range userStatus {
		statuses[jid] = ContactData{
			CurrentStatus: status,
			Log:           userStatusLog[jid],
		}
	}
	mu.RUnlock()

	tmpl := template.Must(template.ParseFiles("status.html"))
	err := tmpl.Execute(w, statuses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
