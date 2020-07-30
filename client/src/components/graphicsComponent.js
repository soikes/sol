export default class GraphicsComponent {
  constructor(graphics, obj) {
    this.graphics = graphics;
    this.obj = obj;
    this.graphics.addToScene(this.obj);
  }
  
  update() {
    this.obj.rotation.y += 0.05;
    // this.obj.rotation.x -= 0.1;
  }

  obj() {
    return this.obj;
  }
}