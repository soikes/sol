export default class Observable {
    constructor() {
      this._listeners = [];
    }
  
    notify(value) {
      this._listeners.forEach(listener => listener(value));
    }
  
    subscribe(listener) {
      this._listeners.push(listener);
    }
  }