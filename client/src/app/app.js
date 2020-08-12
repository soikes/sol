import AppGraphics from './appGraphics';
import AppInput from './appInput';
import AppHud from './appHud';
import AppWorld from './appWorld';

export default class App {
  constructor(container, hud) {
    this.graphics = new AppGraphics(container);
    this.input = new AppInput();
    this.hud = new AppHud(hud);
    this.world = new AppWorld();
    this.entities = [];
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
    this.world.calculateCollisions();
    this.entities.forEach(ent => ent.update(dt));
  }
}
