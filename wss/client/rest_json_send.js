const WebSocket = require("ws")
const readline = require("readline")
const loginUrl = "ws://localhost:3000/user/login"
const socket = new WebSocket(loginUrl)

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
})

socket.onopen = async (e) => {
    try {
        console.log("socket connected..")
    } catch(e) {
        console.log(e.message)
    }
}

rl.on("line", (line) => {

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
    }

    socket.onerror = async (e) => {
        console.debug(e.message)
    }

    // socket.close()
})