export default class AppHud {
    constructor(container) {
        this.container = container;
        this.position = {
            x: container.querySelector("#position #u"),
            z: container.querySelector("#position #w") //TODO do proper data binding dynamically, with attributes
        };
    }

    updatePos(pos) {
        this.position.x.innerText = pos.x.toFixed(3) + "Gm";
        this.position.z.innerText = pos.z.toFixed(3) + "Gm";
    }
}