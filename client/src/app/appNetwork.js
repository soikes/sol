const MsgType = {
  MsgInput: 0,
  MsgTransform: 1,
  MsgSpawn: 2,
  MsgRegister: 3
}

export default class AppNetwork {
  constructor(uri) {
    this.uri = uri;
    this.socket = new WebSocket(this.uri);
    this.outgoing = [];
    this.incoming = [];
    this.socket.onmessage = (e) => { this.queueIncoming(e) };
  }
  
  queueOutgoing(msg) {
    this.outgoing.push(JSON.stringify(msg));
  }

  flushOutgoing() {
    this.outgoing.forEach(msg => this.socket.send(msg));
    this.outgoing.length = 0;
  }

  queueIncoming(event) {
    let msg = JSON.parse(event.data);
    this.incoming.push(msg);
    // console.log(this.1incoming);
  }

  flushIncoming() {
    let inc = [...this.incoming];
    // console.log(inc);
    this.incoming.length = 0;
    return inc;
  }

  filterSpawns(msgs) {
    return msgs.filter(msg => msg.type == MsgType.MsgSpawn);
  }

  filterTransforms(msgs) {
    return msgs.filter(msg => msg.type == MsgType.MsgTransform);
  }
}