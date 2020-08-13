import * as THREE from 'three';

export default class GravityEffectComponent {
    constructor(transform, strength) {
        this.transform = transform;
        this.strength = strength;
        this.active = false;
        this.physics = null;
        this.collidingWith = []; //TODO be able to keep track of everything that is being collided with, can't track one single health component
        this.direction = new THREE.Vector3();
    }

    collideStart(c2) {
        let physics = c2.gameObject.getComponentThatCan("externalAccelerate");
        let otherPos = c2.gameObject.getComponentThatCan("position");
        if (physics != undefined && otherPos != undefined) {
            this.active = true; //TODO Need to have a onCollideStop so that I can stop damaging
            this.physics = physics;
            this.direction = this.transform.position().clone();
            this.direction.sub(otherPos.position()).normalize();
            this.physics.externalAccelerate(this.strength, this.direction);
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