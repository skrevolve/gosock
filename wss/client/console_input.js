const WebSocket = require("ws")
const readline = require("readline")
const loginUrl = "ws://localhost:3000/ws/123?v=1.0"
const socket = new WebSocket(loginUrl)

let message = ""

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
})

socket.onopen = async () => console.log("socket connected..")

rl.on("line", (input_msg) => {
  message = input_msg
  socket.send(message)

  socket.onerror = async (e) => {
    console.debug(e.message)
  }

  // socket.close()
  // rl.close()
})

rl.on("close", () => {
  process.exit()
})