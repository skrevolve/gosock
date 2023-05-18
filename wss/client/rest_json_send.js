const WebSocket = require("ws")
const readline = require("readline")
const loginUrl = "ws://localhost:3000/user/login"
const socket = new WebSocket(loginUrl)

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
})

socket.onopen = async (e) => console.log("socket connected..")

rl.on("line", (input_email) => {

    socket.send(JSON.stringify({
        email: input_email,
        password: Math.random().toString(36).slice(2)
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
    }

    socket.onerror = async (e) => {
        console.debug(e.message)
    }

    // socket.close()
})