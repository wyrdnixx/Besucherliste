
import Test from "./test.vue"


//var app = new Vue({
new Vue({
    el: '#app',
    components: {
        'Test': Test
    },
    data:{
        ws: null,
        serverUrl: "ws://localhost:8080/ws"
    },
    mounted: function() {
        this.connectToWebsocket()
    },
    methods: {
        connectToWebsocket() {
            this.ws = new WebSocket( this.serverUrl);
            this.ws.addEventListener('open', (event) => {
                this.onWebsocketOpen(event)}
                );
        },
        onWebsocketOpen() {
            console.log("connected to WS!");
        }
    },
    render: h => h(App)
}).$mount('#app')