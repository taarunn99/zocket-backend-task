package ws

import (
    "github.com/gorilla/websocket"
    "github.com/gin-gonic/gin"
    "net/http"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer conn.Close()
    clients[conn] = true

    for {
        var msg string
        err := conn.ReadJSON(&msg)
        if err != nil {
            delete(clients, conn)
            break
        }
        broadcast <- msg
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func InitWebSockets(r *gin.Engine) {
    go handleMessages()
    r.GET("/ws", handleConnections)
}
