export default class CollisionDamageComponent {
    constructor(dps) {
        this.dps = dps;
        this.active = false;
        this.health = null;
    }

    onCollide(c2) {
        let health = c2.gameObject.getComponentThatCan("damage");
        if (health != undefined) {
            this.active = true; //TODO Need to have a onCollideStop so that I can stop damaging
            this.health = health;
        }
    }

    update(dt) {
        if (this.active && this.health) {
            this.health.damage(this.dps * dt);
        }
    }
}