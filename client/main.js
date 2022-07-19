var socket = io();

new Vue({
    el: '#chat-room',
    created() {
        // listen for chat message and display on frontend
        socket.on("chat message", (msg) => {
            this.messages.push({
                text: msg,
                date: new Date().toLocaleString()
            })
            // if server is closed, disable send button
            if(msg == "User Disconnected.") {
                this.isDisabled = true
            }
        })

        // listen for welcome message and display on frontend
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
    // automatically send welcome message when entering chat room
    mounted: function() {
        this.sendWelcome()
    },
    methods: {
        // method for send button
        sendMessage() {
            socket.emit('msg', this.message)
            this.message = ''
        },
        // method for welcome message
        sendWelcome() {
            socket.emit('welcome', "New User Connected.")
        }
    }
})