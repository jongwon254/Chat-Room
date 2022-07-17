var socket = io();

new Vue({
    el: '#chat-room',
    created() {

    },
    data: {
        message: '',
        messages: []
    },
    methods: {
        sendMessage() {
            socket.emit('msg', this.message)
            this.message = ''
        }

    }
})