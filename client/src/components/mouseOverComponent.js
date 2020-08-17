export default class MouseOverComponent {
    constructor() {
        this.startHandlers = [];
        this.stopHandlers = [];
    }

    onMouseOverStarted(handler) {
        this.startHandlers.push(handler);
    }

    mouseOverStarted(event) {
        this.startHandlers.forEach(handler => {
            handler(event);
        });
    }

    onMouseOverStopped(handler) {
        this.stopHandlers.push(handler);
    }

    mouseOverStopped(event) {
        this.startHandlers.forEach(handler => {
            handler(event);
        });
    }
}