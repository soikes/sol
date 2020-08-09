import * as THREE from 'three';

export default class AppWorld {
    constructor() {
        this.colliders = [];
        //TODO how will this work? This needs to hold every object's position to be able to calculate collision and gravity physics
    }

    addCollider(cc){ 
        this.colliders.push(cc);
    }

    calculateCollisions() {
        for(let i = 0; i < this.colliders.length; i++) {
            let obj = this.colliders[i];
            for(let j = i + 1; j < this.colliders.length; j++) {
                let obj2 = this.colliders[j];
                if (obj.intersects(obj2)) {
                    console.log("collision detected!");
                }
            }
            // iterate through all objects
            // check if bounding box X, Y or Z is colliding with any other bounding box
            // signal to both collision components that they have collided with each other
            // each collision component will use that state to determine what to do
            // e.g. for a projectile, a collision may mean to: 1. trigger damage on its target, 2. trigger a particle effect, 3. destroy itself
        }
    }
}