export default class SpinComponent {
    constructor(transform) {
        this.transform = transform;
    }

    update() {
        this.transform.rot.y += 0.05;
    }
}