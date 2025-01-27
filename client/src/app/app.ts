import AppGraphics from "./appGraphics";
import AppInput from "./appInput";
import AppHud from "./appHud";
import AppWorld from "./appWorld";
import AppNetwork from "./appNetwork";
// import AppEvents from './appEvents';
import Player from "../entities/player";
import Ship from "../entities/ship";
import GameObject from "../entities/gameObject";
import { MsgType } from "./msg";

export default class App {
  graphics: AppGraphics;
  hud: AppHud;
  world: AppWorld;
  entities: Map<string, any>;
  input: AppInput;
  network: AppNetwork;
  fps: number;
  lastTs: number;
  interval: number;
  stopFrame: any;

  constructor(container: any, hud: AppHud, host: string) {
    this.graphics = new AppGraphics(container);
    this.hud = new AppHud(hud);
    this.world = new AppWorld();
    this.entities = new Map();
    this.input = new AppInput(this.graphics.camera, this.entities);
    this.network = new AppNetwork(`ws://${host}/ws`);
    // this.events = new AppEvents();
    this.configure();
  }

  configure() {
    this.fps = 30;
    this.lastTs = performance.now();
    this.interval = 1000 / this.fps;
  }

  spawn(ent: GameObject) {
    this.entities[ent.id] = ent;
  }

  run(ts) {
    this.stopFrame = window.requestAnimationFrame((ts) => {
      this.run(ts);
    });

    let now = performance.now();
    let dt = now - this.lastTs; // seconds, so calculations are all consistent e.g. m/s
    // this.lastTs = ts;
    // let now = Date.now();
    // let dt = now - this.lastTs;

    if (dt > this.interval) {
      this.lastTs = now - (dt % this.interval);
      this.update(dt);
      this.graphics.render();
    }
  }

  start() {
    window.requestAnimationFrame((ts) => {
      this.run(ts);
    });
  }

  stop() {
    let ts = this.stopFrame();
    window.cancelAnimationFrame(ts);
  }

  update(dt: number) {
    // let playerSpawns = this.network.filterPlayerSpawns(msgs);
    // playerSpawns.forEach((s) => {
    //   console.log("player spawn: " + s);
    //   Player.build(
    //     s.data.id,
    //     this.graphics,
    //     this.input,
    //     this.hud,
    //     this.world,
    //     this.network,
    //   ).then((spawn) => {
    //     console.log(spawn);
    //     this.spawn(spawn);
    //   });
    // });
    // let spawns = this.network.filterSpawns(msgs);
    // spawns.forEach((s) => {
    //   console.log("spawn: " + s);
    //   Ship.build(s.data.id, this.graphics, this.world).then((spawn) => {
    //     console.log(spawn);
    //     this.spawn(spawn);
    //   });
    // }); //TODO this does not belong here.
    // // These should be processed asynchronously and
    // // placed in a queue to be consumed by the game loop.
    // let transforms = this.network.filterTransforms(msgs);
    // this.world.calculateCollisions();
    // this.entities.forEach((ent) => {
    //   transforms.forEach((t) => {
    //     if (ent.id == t.data.id) {
    //       let t1 = ent.getComponentThatCan("setPos");
    //       t1.setPos(t.data.pos.x, t.data.pos.y, t.data.pos.z);
    //       let t2 = ent.getComponentThatCan("setRot");
    //       t2.setRot(t.data.rot.x, t.data.rot.y, t.data.rot.z);
    //     }
    //   });
    //   ent.update(dt);
    // });
    // this.network.flushOutgoing();
  }

  processIncomingMsgs() {
    let msgs = this.network.flushIncoming();

    for (var i = 0; i < msgs.length; i++) {
      let m = msgs[i];
      switch (m.type) {
        case MsgType.Transform:
          if (this.entities[m.data.id]) break;
        case MsgType.Spawn:
          break;
      }
    }
  }
}
