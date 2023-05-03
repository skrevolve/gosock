const WebSocket = require("ws")
const readline = require("readline")

const socket = new WebSocket("ws://localhost:3000/ws/123?v=1.0")

let clientMessage

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
})

rl.on("line", (line) => {

  console.log("input: ", line)
  clientMessage = line

  socket.onopen = (e) => {

    socket.send(clientMessage)

    socket.onmessage = (e) => {
      console.debug("client received a message")
      // console.debug(e)
    }

    socket.onclose = (e) => {
      console.debug("client notified socket has closed")
      // console.debug(e)
    }

    socket.onerror = (e) => {
      console.debug(e.message) // console.log(e)
    }

    // socket.close()
  }

  // rl.close()
})

rl.on("close", () => {
  process.exit()
})