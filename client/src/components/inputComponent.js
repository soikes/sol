export default class InputComponent {
  constructor(transform, input, physics, network) {
    this.transform = transform;
    this.input = input;
    this.physics = physics;
    this.network = network;
  }

  update() {
    let send = false;
    var msg = {
      type: 0,
      data: {
        id: this.gameObject.id
      }
    }
    if (this.input.forwardPressed()) {
      msg.data.forward = true
      send = true;
      this.physics.accelerate();
    } else {
      this.physics.stopAccelerating();
    }
    
    if (this.input.turnLeftPressed()) {
      msg.data.left = true;
      send = true;
      this.physics.rotating = true;
      this.physics.rotateDirection.y = 1; 
    } else if (this.input.turnRightPressed()) {
      msg.data.right = true;
      send = true;
      this.physics.rotating = true;
      this.physics.rotateDirection.y = -1; 
    } else {
      this.physics.rotating = false;
    }

    if (send) {
      this.network.queueOutgoing(msg);
    }
    this.input.updateMouseIntersects();
  }
}
