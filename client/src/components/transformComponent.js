export default class TransformComponent {
  constructor(pos, dir, scale) {
    if (this.pos === null) {
      this.pos = new THREE.Vector3();
    } else {
      this.pos = pos;
    }
    
    if (this.rot === null) {
      this.rot = new THREE.Vector3();
    } else {
      this.rot = dir;
    }

    this.scale = scale;
  }

  update() {}
}