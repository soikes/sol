import * as THREE from 'three';

export default class PhysicsComponent {
  constructor(velocity, maxSpeed, accelerationFactor, rotationFactor, transformComponent) {
    this.velocity = velocity;
    this.maxSpeed = maxSpeed;
    this.accelerationFactor = accelerationFactor;
    this.accelerating = false;

    this.rotationFactor = rotationFactor;
    this.rotating = false;
    this.rotateDirection = new THREE.Vector3();

    this.transformComponent = transformComponent;
  }

  update(dt) {
    this.updateRotation();
    this.calculateVelocity(dt);
    this.updatePosition(dt); //TODO pass in dt into update() so you can use it here
  }

  updateRotation() {
    if (this.rotating) {
      let rd = this.rotateDirection.clone();
      let r = rd.multiplyScalar(this.rotationFactor);
      this.transformComponent.rot.add(r);
      if (this.transformComponent.rot.x < 0) {
        this.transformComponent.rot.x = 2*Math.PI;
      }
      if (this.transformComponent.rot.x > 2*Math.PI) {
        this.transformComponent.rot.x = 0;
      }
      if (this.transformComponent.rot.y < 0) {
        this.transformComponent.rot.y = 2*Math.PI;
      }
      if (this.transformComponent.rot.y > 2*Math.PI) {
        this.transformComponent.rot.y = 0;
      }
      if (this.transformComponent.rot.z < 0) {
        this.transformComponent.rot.z = 2*Math.PI;
      }
      if (this.transformComponent.rot.z > 2*Math.PI) {
        this.transformComponent.rot.z = 0;
      }
    }
  }

  calculateVelocity(dt) {
    if (this.accelerating) {
      let vix = this.velocity.x;
      let vfx = vix + (this.accelerationFactor * Math.sin(this.transformComponent.rot.y)) * dt;

      let viz = this.velocity.z;
      let vfz = viz + (this.accelerationFactor * Math.cos(this.transformComponent.rot.y)) * dt;

      let dir = new THREE.Vector2(vfx, vfz).normalize();
      let mag = Math.sqrt(vfx*vfx + vfz*vfz);

      if (mag <= this.maxSpeed) {
        this.velocity.x = vfx;
        this.velocity.z = vfz;
      } else {
        this.velocity.x = dir.x * this.maxSpeed;
        this.velocity.z = dir.y * this.maxSpeed; // "y" refers to second value of new Vector2, aka "z"
      }
    }
  }

  updatePosition(dt) {
    let d = new THREE.Vector3(this.velocity.x, 0, this.velocity.z).multiplyScalar(dt);
    this.transformComponent.addPos(d);
  }

  bounce() {
    this.velocity.multiplyScalar(-1); // TODO This doesn't work properly lol
  }

}