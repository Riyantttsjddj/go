package main

import (
    "golang.org/x/net/websocket"
    "io"
    "log"
    "net"
    "net/http"
)

func handleWS(ws *websocket.Conn) {
    // Menghubungkan ke SSH server di localhost (port 22)
    sshConn, err := net.Dial("tcp", "127.0.0.1:22") // Ganti dengan port atau alamat server yang sesuai
    if err != nil {
        log.Println("Gagal terhubung ke SSH:", err)
        return
    }
    defer sshConn.Close()

    // Menyalurkan data dua arah antara WebSocket dan SSH
    go io.Copy(sshConn, ws)
    io.Copy(ws, sshConn)
}

func main() {
    // Menjalankan server WebSocket pada port 80
    http.Handle("/", websocket.Handler(handleWS))
    log.Println("WebSocket server berjalan di port 80...")
    err := http.ListenAndServe(":80", nil) // Pastikan server mendengarkan pada port 80
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
