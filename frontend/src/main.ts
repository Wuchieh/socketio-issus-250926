import {io} from 'socket.io-client'

const ws1 = io('http://192.168.213.13:9090', {path: '/ws', autoConnect: false})
const ws2 = io('http://192.168.213.13:9091', {path: '/ws', autoConnect: false})
const ws3 = io('http://192.168.213.13:9092', {path: '/ws', autoConnect: false})

const ws1Btn = document.querySelector("#ws1")! as HTMLButtonElement
const ws2Btn = document.querySelector("#ws2")! as HTMLButtonElement
const ws3Btn = document.querySelector("#ws3")! as HTMLButtonElement

ws1.on('connect', () => {
    ws1Btn.innerText = 'ws1 connected'
})
ws2.on('connect', () => {
    ws2Btn.innerText = 'ws2 connected'
})
ws3.on('connect', () => {
    ws3Btn.innerText = 'ws3 connected'
})

ws1Btn.addEventListener('click', () => {
    ws1.connect()
})

ws2Btn.addEventListener('click', () => {
    ws2.connect()
})

ws3Btn.addEventListener('click', () => {
    ws3.connect()
})
