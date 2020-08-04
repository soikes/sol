import * as THREE from 'three';

export default class TransformComponent {
  constructor(pos, rot, scale) {
    if (this.pos === null) {
      this.pos = new THREE.Vector3();
    } else {
      this.pos = pos;
    }
    
    if (this.rot === null) {
      this.rot = new THREE.Vector3();
    } else {
      this.rot = rot;
    }

    this.scale = scale;
  }

  update() {}
}