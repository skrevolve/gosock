const WebSocket = require("ws")
const url = "ws://localhost:3000/login"
const socket = new WebSocket(url)

let data = {
    email: "test@email.com",
    password: "12345!@#$%"
}

socket.onopen = async () => console.log("socket connected..")

socket.send(JSON.stringify(data))

socket.onmessage = async (e) => {
    try {
        console.debug("client received a message")
        console.debug(e)
        // if (e !== null && e !== undefined) {
        //     const resData = await JSON.parse(e.data)
        //     console.log(resData)
        // }
    } catch (e) {
        console.log(e.message)
    }
}

socket.onclose = async (e) => {
    console.debug("client notified socket has closed")
    console.debug(e.message)
}

socket.onerror = async (e) => {
    console.log("onerror")
    console.debug(e.message)
}

socket.close()