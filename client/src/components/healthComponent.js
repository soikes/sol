export default class HealthComponent {
    constructor(min, max, amount) {
        this.changed = false;
        this.min = min;
        this.max = max;
        this.amount = amount;
    }

    damage(amount) {
        this.amount -= amount;
        this.changed = true;
    }

    dead() {
        return (this.amount <= this.min);
    }

    observe(observer) {
        this.observer = observer;
        this.notify(this.amount);
    }

    stopObserving() {
        this.observer = null;
    }

    notify(val) {
        if (this.observer) {
            this.observer.notify(val);
        }
    }

    update() {
        if (this.changed) {
            this.notify(this.amount);
        }
        this.changed = false;
        this.damage(0.01);
    }
}