package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(jet.NewOSFileSystemLoader("./html"), jet.InDevelopmentMode())

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Suppose to check if the origin of the initial request is equal to current incoming request
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Home(w http.ResponseWriter, r *http.Request) {

	err := renderPage(w, "home.jet", nil)
	if err != nil {

		log.Println("There was an error 😩", err)

	}

}

type WebSocketConnection struct {
	// The Conn type represents a WebSocket connection
	*websocket.Conn
}

// WebsocketJsonResponse defines the response sent back from the websocket to the client
type WebsocketJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

// WebsocketPayload defines the response sent from the client to the websocket(server)
type WebsocketPayload struct {
	Action     string              `json:"action"`
	Username   string              `json:"username"`
	Message    string              `json:"message"`
	Connection WebSocketConnection `json:"-"`
}

// WebsocketEndPoint upgrades a "regular" connection to websocket level
func WebsocketEndPoint(w http.ResponseWriter, r *http.Request) {

	// The third argument is the response header
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {

		log.Println("There was an error upgrading the connection 😩", err)

	}

	fmt.Println("Client connected to web socket end point 🤟")

	var response WebsocketJsonResponse
	response.Message = `<em><small>Connected to server</em></small>`

	err = ws.WriteJSON(response)
	if err != nil {

		log.Println("There was an error 😩", err)

	}

}

// data = the data we want to pass to our html template(tmpl)
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {

	view, err := views.GetTemplate(tmpl)

	if err != nil {

		log.Println("There was an error 😩", err)
		return err
	}

	// the third argument is for data that can be passed in the session context
	err = view.Execute(w, data, nil)
	if err != nil {

		log.Println("There was an error 😩", err)
		return err
	}

	return nil
}
