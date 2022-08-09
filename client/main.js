// websocket
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
        history: [],
        isDisabled: false,
        show: false
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
        },
        // method for getting history
        getHistory() {
            this.show = true
            fetch("http://localhost:8080/api/messages")
            .then(data => {
                return data.json();
            })
            .then(response => {
                this.history = response
            })
        },
        // method for deleting history
        deleteHistory() {
            this.show = false
            fetch("http://localhost:8080/api/delete", {
                method: 'DELETE'
            })
            .then(data => {
                return data.json();
            })
            .then(response => {
                console.log(response)
            })

            // update empty history
            this.getHistory()
        }
    }
})