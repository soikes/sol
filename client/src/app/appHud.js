export default class AppHud {
    constructor(container) {
        this.container = container;
        this.position = {
            x: container.querySelector("#position #u"),
            z: container.querySelector("#position #w") //TODO do proper data binding dynamically, with attributes
        };
        this.health = {
            min: container.querySelector("#health #min"),
            max: container.querySelector("#health #max"),
            amount: container.querySelector("#health #amount")
        };
    }

    updatePos(pos) {
        this.position.x.innerText = pos.x.toFixed(3) + "Gm";
        this.position.z.innerText = pos.z.toFixed(3) + "Gm";
    }

    updateHealth(amount) {
        // this.health.min.innerText = min;
        // this.health.max.innerText = max;
        this.health.amount.innerText = amount.toFixed(2) + "%";
    }
}