const WebSocket = require("ws")
const readline = require("readline")
const loginUrl = "ws://localhost:3000/user/login?method=POST"
const socket = new WebSocket(loginUrl)

let message = ""

let data = {
  email: "test@email.com",
  password: "12345!@#$%"
}

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
})

socket.onopen = async () => console.log("socket connected..")

rl.on("line", (line) => {

  // message = line

  // socket.send(message)
  socket.send(JSON.stringify({
    email: line,
    password: "12345!@#$%"
  }))

  socket.onmessage = async (e) => {
    try {
      if (e !== null && e !== undefined) {
        const resData = await JSON.parse(e.data)
        console.log(resData)
      }
    } catch (e) {
      console.log(e.message)
    }
    // console.debug("client received a message")
    // console.debug(e)
  }

  // socket.onclose = async (e) => {
  //   console.debug("client notified socket has closed")
  //   // console.debug(e)
  // }

  socket.onerror = async (e) => {
    console.debug(e.message)
  }

  // socket.close()


  // rl.close()
})

rl.on("close", () => {
  process.exit()
})