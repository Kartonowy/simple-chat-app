import WebSocket from "ws";

const ws = new WebSocket("ws://localhost:8080/echo");

ws.on("open", function open() {
  ws.send("i love macklemore");
});

ws.on("message", function message(message) {
  console.log(new TextDecoder().decode(message));
});
