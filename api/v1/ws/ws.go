package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/services"
	"gochat/utils"
)

type MessageObject struct {
	Data  string `json:"data"`
	From  string `json:"from"`
	Event string `json:"event"`
	To    string `json:"to"`
}

type ContactObject struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SocketIORouter(app fiber.Router) {

	clients := make(map[string]string)

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			c.Locals("Authorization", c.GetReqHeaders()["Authorization"])
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		fmt.Printf("Connection event 1 - User: %s \n", ep.Kws.GetStringAttribute("user_id"))
		userService := services.UserService{}
		user, err := userService.GetUserById(ep.Kws.GetStringAttribute("user_id"))
		if err != nil {
			ep.Kws.Emit([]byte("User Not Found"))
			ep.Kws.Close()
			return
		}
		ep.Kws.Emit([]byte(fmt.Sprintf("%v", user.Contacts)))

	})

	// On message event
	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {

		//fmt.Printf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data))

		message := MessageObject{}

		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
			return
		}
		message.From = ep.SocketUUID
		if message.Event != "" {
			ep.Kws.Fire(message.Event, []byte(message.Data))
			return
		}

		payload, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		socketUid, ok := clients[message.To]
		if !ok {
			ep.Kws.Emit([]byte(fmt.Sprintf("User Not Active")), socketio.TextMessage)
			return
		}
		err = ep.Kws.EmitTo(socketUid, payload, socketio.TextMessage)
		if err != nil {
			fmt.Println(err)
		}
	})

	socketio.On("Contact", func(ep *socketio.EventPayload) {
		contactObject := ContactObject{}
		err := json.Unmarshal(bytes.Replace(ep.Data, []byte("'"), []byte("\""), -1), &contactObject)
		if err != nil {
			fmt.Println(err)
			return
		}

		userService := services.UserService{}
		user, err := userService.UpdateContacts(ep.Kws.GetStringAttribute("user_id"), contactObject.Name, contactObject.Email)
		if err != nil {
			ep.Kws.Emit([]byte(fmt.Sprintf(err.Error())), socketio.TextMessage)
			return
		}
		ep.Kws.Emit([]byte(fmt.Sprintf("Updated Successfully %v", user.Contacts)), socketio.TextMessage)

	})
	// On disconnect event
	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		ep.Kws.Close()
		fmt.Println(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On error event
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		fmt.Println(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	app.Get("/ws/", socketio.New(func(kws *socketio.Websocket) {

		authToken := kws.Locals("Authorization")
		token, err := utils.VerifyJWTToken(authToken.([]string)[0])
		if err != nil {
			kws.Emit([]byte(fmt.Sprintf("Authentication Error")), socketio.TextMessage)
			kws.Close()
			return
		}

		userId := token["id"].(string)
		_, ok := clients[userId]

		if ok {
			kws.Emit([]byte(fmt.Sprintf("Already Active")), socketio.TextMessage)
			kws.Close()
			return
		}

		clients[userId] = kws.UUID

		kws.SetAttribute("user_id", userId)

		//kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true, socketio.TextMessage)

		kws.Emit([]byte(fmt.Sprintf("Hello user: %s with UUID: %s", userId, kws.UUID)), socketio.TextMessage)
	}))

}
