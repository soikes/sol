export default class GraphicsComponent {
  constructor(graphics, obj, transform) {
    this.graphics = graphics;
    this.obj = obj;
    this.graphics.addToScene(this.obj);
    this.transform = transform;
  }

  update() {
    this.obj.position.x = this.transform.pos.x;
    this.obj.position.y = this.transform.pos.y;
    this.obj.position.z = this.transform.pos.z;

    this.obj.rotation.x = this.transform.rot.x;
    this.obj.rotation.y = this.transform.rot.y;
    this.obj.rotation.z = this.transform.rot.z;
  }

  obj() {
    return this.obj;
  }
}
