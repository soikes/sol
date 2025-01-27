import { MsgType, Msg } from "./msg";

const SpawnType = {
  PlayerSpawn: 0,
};

export default class AppNetwork {
  uri: string;
  socket: WebSocket;
  outgoing: Array<string>;
  incoming: Array<Msg>;

  constructor(uri: string) {
    this.uri = uri;
    this.socket = new WebSocket(this.uri);
    this.outgoing = [];
    this.incoming = [];
    this.socket.onmessage = (e) => {
      this.queueIncoming(e);
    };
  }

  queueOutgoing(msg: Msg) {
    this.outgoing.push(JSON.stringify(msg));
  }

  flushOutgoing() {
    this.outgoing.forEach((msg) => this.socket.send(msg));
    this.outgoing.length = 0;
  }

  queueIncoming(event: any) {
    let msg = JSON.parse(event.data);
    if (this.isMsg(msg)) {
      this.incoming.push(msg as Msg);
    }
  }

  isMsg(msg: any): msg is Msg {
    return (
      msg &&
      typeof msg == "object" &&
      typeof msg.id == "string" &&
      typeof msg.seq == "number" &&
      typeof msg.ts == "number"
    );
  }

  flushIncoming() {
    let inc = [...this.incoming];
    // console.log(inc);
    this.incoming.length = 0;
    return inc;
  }

  // filterSpawns(msgs) {
  //   return msgs.filter(
  //     (msg) =>
  //       msg.type == MsgType.MsgSpawn && msg.data.type != SpawnType.PlayerSpawn,
  //   );
  // }

  // filterRegistered(msgs) {
  //   return msgs.filter((msg) => msg.type == MsgType.MsgRegistered);
  // }

  // filterTransforms(msgs) {
  //   return msgs.filter((msg) => msg.type == MsgType.MsgTransform);
  // }

  // filterPlayerSpawns(msgs) {
  //   return msgs.filter(
  //     (msg) =>
  //       msg.type == MsgType.MsgSpawn && msg.data.type == SpawnType.PlayerSpawn,
  //   );
  // }
}
