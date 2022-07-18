var socket = io();

new Vue({
    el: '#chat-room',
    created() {
        socket.on("chat message", (msg) => {
            this.messages.push({
                text: msg,
                date: new Date().toLocaleString()
            })
            if(msg == "User Disconnected.") {
                this.isDisabled = true
            }
        })

        socket.on("welcome", (msg) => {
            this.messages.push({
                text: msg,
                date: new Date().toLocaleString()
            })
        })

        
    },
    data: {
        message: '',
        messages: [],
        isDisabled: false
    },
    mounted: function() {
        this.sendWelcome()
    },
    methods: {
        sendMessage() {
            socket.emit('msg', this.message)
            this.message = ''
        },
        sendWelcome() {
            socket.emit('welcome', "New User Connected.")
        }

    }
})