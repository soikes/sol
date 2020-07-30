export default class InputComponent {
  constructor(transform, input) {
    this.transform = transform;
    this.input = input;
  }

  update() {
    if (this.input.forwardPressed()) {
      this.transform.pos.x += 0.1;
    }
  }
}
