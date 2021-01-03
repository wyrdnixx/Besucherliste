<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
     <h2>Con: {{ constate }}</h2>
<div id="example">
  <button v-on:click="greet">Greet</button>
  <button v-on:click="connectToWebsocket">connectToWebsocket</button>
</div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  components: {},
  props: {
    msg: String
  },
  data() {
   return {
      constate: "", 
      ws: null,
      serverUrl: "ws://localhost:8080/ws"      
   }
  },
 
    methods: {
       //greet: function (event) {
       greet: function () {
          var websocketServerLocation = "ws://" + window.location.hostname + ":8080/ws";

    //  var websocketServerLocation = "ws://localhost:8080/ws";
   
        var socket = null;
        socket = new WebSocket(websocketServerLocation)
        socket.OPEN;
    //alert('Socket ' + socket.readyState + '!')
        this.constate = socket.readyState;
    },
    
    connectToWebsocket() {
        this.ws = new WebSocket( this.serverUrl );
        this.ws.addEventListener('open', (event) => { this.onWebsocketOpen(event) });
      },
      onWebsocketOpen() {
        console.log("connected to WS!");
      }
    }
}

 
    




</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
