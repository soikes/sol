import * as THREE from 'three';

export default class TransformComponent {
  constructor(pos, rot, scale) {
    this.changed = false;
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

  observe(observer) {
    this.observer = observer;
    this.notify(this.pos);
  }

  stopObserving() {
    this.observer = null;
  }

  notify(val) {
    if (this.observer) {
      this.observer.notify(val);
    }
  }

  addPos(pos) {
    this.pos.add(pos);
    this.changed = true;
  }

  setPos(x, y, z) {
    this.pos.setX(x)
    this.pos.setY(y)
    this.pos.setZ(z)
    this.changed = true;
  }

  update() {
    if (this.changed) {
      this.notify(this.pos);
    }
    this.changed = false;
  }

  position() {
    return this.pos;
  }

  listenToNetwork(network) {
    this.network = network;
  }
}