export default class CollisionBounceComponent {
    constructor(knockback) {
        this.knockback = knockback;
        this.active = false;
        this.physics = null;
        this.collidingWith = []; //TODO be able to keep track of everything that is being collided with, can't track one single health component
    }

    collideStart(c2) {
        let physics = c2.gameObject.getComponentThatCan("bounce");
        if (physics != undefined) {
            this.active = true; //TODO Need to have a onCollideStop so that I can stop damaging
            this.physics = physics;
            this.physics.bounce(this.knockback);
        }
    }

    collideStop(c2) {
        this.active = false;
        this.physics = null;
    }

    update(dt) {
        // if (this.active && this.physics) {
        //     this.physics.bounce();
        // }
    }
}