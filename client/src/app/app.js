import AppGraphics from './appGraphics';
import AppInput from './appInput';
import AppHud from './appHud';
import AppWorld from './appWorld';
import AppNetwork from './appNetwork';
// import AppEvents from './appEvents';

import Ship from '../entities/ship'

export default class App {
  constructor(container, hud, host) {
    this.graphics = new AppGraphics(container);
    this.hud = new AppHud(hud);
    this.world = new AppWorld();
    this.entities = [];
    this.input = new AppInput(this.graphics.camera, this.entities);
    this.network = new AppNetwork(`ws://${host}/ws`);
    // this.events = new AppEvents();
    this.configure();
  }

  configure() {
    // this.fps = 60;
    this.lastTs = performance.now();
    this.interval = 1000 / this.fps;
  }

  spawn(ent) {
    this.entities.push(ent);
  }

  run(ts) {
    this.stopFrame = window.requestAnimationFrame((ts) => { this.run(ts); });

    let dt = (ts - this.lastTs) / 1000; // seconds, so calculations are all consistent e.g. m/s
    this.lastTs = ts;
    // let now = Date.now();
    // let dt = now - this.lastTs;

    // if (dt > this.interval) {
      // this.lastTs = now - (dt % this.interval);
      this.update(dt);
      this.graphics.render();
    // }
  }

  start() {
    window.requestAnimationFrame((ts) => { this.run(ts); });
  }

  stop() {
    window.cancelAnimationFrame(() => { this.stopFrame(); });
  }

  update(dt) {
    let msgs = this.network.flushIncoming();
    let spawns = this.network.filterSpawns(msgs);
    spawns.forEach(s => {
      console.log(s)
      Ship.build(s.data.id, this.graphics, this.input, this.hud, this.world, this.network)
        .then(ship => { console.log(ship); this.spawn(ship); });
    }); //TODO this does not belong here

    let transforms = this.network.filterTransforms(msgs);
    this.world.calculateCollisions();
    this.entities.forEach(ent => {
      transforms.forEach(t => {
        if (ent.id == t.data.id) {
          let t1 = ent.getComponentThatCan('setPos');
          t1.setPos(t.data.x, t.data.y, t.data.z);
        }
      })
      ent.update(dt);
    });
    this.network.flushOutgoing();
  }
}
