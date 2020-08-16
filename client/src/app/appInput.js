import * as THREE from 'three';

const Key = {
  _pressed: {},

  LEFT: 65,
  UP: 87,
  RIGHT: 68,
  DOWN: 83,
  SPACE: 32,

  isDown: function (keyCode) {
    return this._pressed[keyCode];
  },

  onKeydown: function (event) {
    this._pressed[event.keyCode] = true;
  },

  onKeyup: function (event) {
    delete this._pressed[event.keyCode];
  }
};

export default class AppInput {
  constructor(camera, entities) {
    this.raycaster = new THREE.Raycaster();
    this.mouse = new THREE.Vector2();
    this.camera = camera;
    this.entities = entities;
    this.intersects = [];
    this.lastIntersects = [];

    window.addEventListener('keyup', (event) => { Key.onKeyup(event); }, false);
    window.addEventListener('keydown', (event) => { Key.onKeydown(event); }, false);
    window.addEventListener('mousemove', (event) => {
      this.mouse.x = (event.clientX / window.innerWidth) * 2 - 1;
      this.mouse.y = - (event.clientY / window.innerHeight) * 2 + 1;
    }, false);
  }

  updateMouseIntersects() {
    this.raycaster.setFromCamera(this.mouse, this.camera);
    let components = this.entities.map(entity => entity.getComponentThatCan("graphicsObject"));
    let objects = components.map(entity => entity.graphicsObject());
    this.intersects = this.raycaster.intersectObjects(objects).map(i => i.object.id);
    // let diff = this.lastIntersects
    //   .filter(i => !this.intersects.includes(i))
    //   .concat(this.intersects.filter(i => !this.lastIntersects.includes(i)));

    if (this.intersects.length > 0 ) {
      let a = this.lastIntersects.filter(i => !this.intersects.includes(i));
      let b = this.intersects.filter(i => !this.lastIntersects.includes(i));

      // at t = 1:  a = [] b = [], diff = []
      // at t = 2:  a = [] b = [obj], diff = [obj]
      // at t = 3:  a = [obj] b = [obj] diff = []

      let diff = a.concat(b);

      if (diff.length > 0) {
        let intersected = diff[0];
        for (let i = 0; i < components.length; i++) {
          if (intersected === components[i].graphicsObject().id) {
            console.log(`mouse is over component ${components[i]}`);
          }
        }
      }

      this.lastIntersects = this.intersects;
    } else {
      for (let i = 0; i < this.lastIntersects; i++) {
        for (let j = 0; j < components.length; j++) {
          if (this.lastIntersects[i] === components[j].graphicsObject().id) {
            console.log(`mouse stopped being over component ${components[j]}`);
          }
        }
      }
      this.lastIntersects.length = 0;
    }

    // for ( var i = 0; i < this.intersects.length; i++ ) {
    //   this.intersects[i].object.material.color.set(0xff0000);
    //   // console.log("object mouseover");
    // }
  }

  forwardPressed() {
    return Key.isDown(Key.UP);
  }

  backwardPressed() {
    return Key.isDown(Key.DOWN);
  }

  turnLeftPressed() {
    return Key.isDown(Key.LEFT);
  }

  turnRightPressed() {
    return Key.isDown(Key.RIGHT);
  }
}
