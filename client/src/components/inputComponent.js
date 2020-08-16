export default class InputComponent {
  constructor(transform, input, physics) {
    this.transform = transform;
    this.input = input;
    this.physics = physics;
  }

  update() {
    if (this.input.forwardPressed()) {
      this.physics.accelerate();
    } else {
      this.physics.stopAccelerating();
    }
    
    if (this.input.turnLeftPressed()) {
      this.physics.rotating = true;
      this.physics.rotateDirection.y = 1; 
    } else if (this.input.turnRightPressed()) {
      this.physics.rotating = true;
      this.physics.rotateDirection.y = -1; 
    } else {
      this.physics.rotating = false;
    }

    this.input.updateMouseIntersects();
  }
}
