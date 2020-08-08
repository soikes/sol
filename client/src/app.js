import AppGraphics from './appGraphics';
import AppInput from './appInput';
import AppHud from './appHud';

export default class App {
  constructor(container, hud) {
    this.graphics = new AppGraphics(container);
    this.input = new AppInput();
    this.hud = new AppHud(hud);
    this.entities = [];
  }

  spawn(ent) {
    this.entities.push(ent);
  }

  run(t) {
    this.stopFrame = window.requestAnimationFrame((t) => { this.run(t); });
    this.update();
    this.graphics.render();
  }

  start() {
    window.requestAnimationFrame((t) => { this.run(t); });
  }

  stop() {
    window.cancelAnimationFrame(() => { this.stopFrame(); });
  }

  update() {
    this.entities.forEach(ent => ent.update());
  }
}
