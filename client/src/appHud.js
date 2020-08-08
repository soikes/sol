export default class AppHud {
    constructor(container) {
        this.container = container;
        this.position = {
            x: container.querySelector("#position #u"),
            y: container.querySelector("#position #v"),
            z: container.querySelector("#position #w") //TODO do proper data binding dynamically, with attributes
        };
    }

    updatePos(pos) {
        this.position.x.innerText = pos.x;
        this.position.y.innerText = pos.y;
        this.position.z.innerText = pos.z;
    }
}