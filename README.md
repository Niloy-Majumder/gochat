# Go-Chat

> A WebSocket chat application written in Golang
## Technologies Used

- **Language:** Golang
- **Framework:** GoFiber
- **Database**: MongoDB 
- **Logging:** Implemented proper logging mechanisms to track events and errors
- **Rate-Limiting:** Applied rate-limiting on public routes to prevent abuse and ensure stability

## Getting Started

To start the application in production mode, simply run:

```bash
make start_prod
```

![Server Start](./Assets/Server_start.png)
## REST API for User Authentication

### Register Users

Use the following API to register new users:

![Register_User](./Assets/Register_User.png)
### Login

To authenticate a user, use the login API:

![Login_User](./Assets/Login_User.png)
### JWT Authentication

Upon successful login, a JWT token will be returned in the `token` field. This token should be included in the headers of subsequent requests with the `Authorization` key:

![Authorization](./Assets/Authentication.png)
## WebSocket Communication

### Message Format

The server accepts messages in JSON format with the following structure:

```json
{ 	
	"data": "",
	"event": "",
	"to": ""
}
```

### Adding Contacts

To add other users to your contact list, use `"Contact"` as the event and specify the contact details in the `data` field. For example:

```json
{ 	
	"data": "{'name': 'Jason', 'email': 'abc@xyz.com'}",
	"event": "Contact",
	"to": ""
}
```

![Add_Contacts](./Assets/Add_Contacts.png)

Upon connecting to the WebSocket, you will receive an array of your contacts.

### Sending Messages

To send a message to another user connected to the server, you need to know the user's `userId`, which you can obtain from your saved contacts. Place your message in the `data` field, specify the recipient's `userId` in the `to` field, and leave the `event` field empty.

![User_1](./Assets/User_1.png)![User_2](./Assets/User_2.png)
