var socket = io();

new Vue({
    el: '#chat-room',
    created() {
        socket.on("chat message", (msg) => {
            this.messages.push({
                text: msg,
                date: new Date().toLocaleString()
            })
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
        messages: []
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