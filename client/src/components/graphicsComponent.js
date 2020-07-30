export default class GraphicsComponent {
  constructor(graphics, obj, transform) {
    this.graphics = graphics;
    this.obj = obj;
    this.graphics.addToScene(this.obj);
    this.transform = transform;
  }

  update() {
    this.obj.position.x = this.transform.pos.x;
    this.obj.rotation.y += 0.05;
    // this.obj.rotation.x -= 0.1;
  }

  obj() {
    return this.obj;
  }
}
