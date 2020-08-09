export default class GraphicsComponent {
  constructor(graphics, obj, transform, offset) {
    this.graphics = graphics;
    this.obj = obj;
    this.graphics.addToScene(this.obj);
    this.transform = transform;
    this.offset = offset;
  }

  update() {
    let pos = this.transform.pos.clone();
    if (this.offset) {
      pos.add(this.offset);
    }
    this.obj.position.copy(pos);

    this.obj.rotation.x = this.transform.rot.x;
    this.obj.rotation.y = this.transform.rot.y;
    this.obj.rotation.z = this.transform.rot.z;
  }

  object() {
    return this.obj;
  }
}
