export default class CollisionDamageComponent {
    constructor(dps) {
        this.dps = dps;
        this.active = false;
        this.health = null;
        this.collidingWith = []; //TODO be able to keep track of everything that is being collided with, can't track one single health component
    }

    collideStart(c2) {
        let health = c2.gameObject.getComponentThatCan("damage");
        if (health != undefined) {
            this.active = true; //TODO Need to have a onCollideStop so that I can stop damaging
            this.health = health;
        }
    }

    collideStop(c2) {
        this.active = false;
        this.health = null;
    }

    update(dt) {
        if (this.active && this.health) {
            this.health.damage(this.dps * dt);
        }
    }
}