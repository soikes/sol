export default class AppState {
  constructor() {
    this.states = [];
    this.currentState = null;
  }

  pushState(state) {
    this.states.push(state);
  }
}