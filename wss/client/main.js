const WebSocket = require("ws")
const socket = new WebSocket("ws://localhost:3000/ws/123?v=1.0")

socket.onopen = function (e) {
  socket.send("hello")
}


socket.onmessage = async function (e) {

}

socket.onerror = function(e) {
  console.log(e)
}