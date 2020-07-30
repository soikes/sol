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
  constructor() {
    window.addEventListener('keyup', (event) => { Key.onKeyup(event); }, false);
    window.addEventListener('keydown', (event) => { Key.onKeydown(event); }, false);
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
