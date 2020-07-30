import AppGraphics from './appGraphics';

export default class App {
  constructor() {
    this.graphics = new AppGraphics();

    this.entities = [];
  }

  spawn(ent) {
    this.entities.push(ent);
  }

  run(t) {
    this.stopFrame = window.requestAnimationFrame((t) => { this.run(t) });
    this.update();
    this.graphics.render();
  }

  start() {
    window.requestAnimationFrame((t) => { this.run(t) });
  }

  stop() {
    window.cancelAnimationFrame(() => { this.stopFrame() });
  }

  update() {
    this.entities.forEach(ent => ent.update());
  }
};