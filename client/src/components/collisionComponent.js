import * as THREE from 'three';

export default class CollisionComponent {
    constructor(object, transform, graphics, physics, ...collisionHandlers) {
        this.box = new THREE.Box3();
        this.object = object;
        this.box.setFromObject(object);
        this.helper = new THREE.Box3Helper( this.box, 0xffff00 );
        graphics.addToScene(this.helper);

        this.transform = transform;
        this.physics = physics;
        this.colliding = false;
        if (collisionHandlers === undefined) {
            this.collisionHandlers = [];
        } else {
            this.collisionHandlers = collisionHandlers;
        }
    }

    computeBounding() {
        this.box.setFromObject(this.object);
    }

    collide() {
        this.colliding = true;
    }

    stoppedColliding() {
        this.colliding = false;
    }

    collideStart(c2) {
        this.collisionHandlers.forEach(handler => {
            handler(c2);
        });
    }

    intersects(collision) {
        if (this.box.intersectsBox(collision.box)) {
            this.collide();
            collision.collide();
            return true;
        } else {
            this.stoppedColliding();
        }
        return false;
    }

    update() {
        this.box.setFromObject(this.object);
        this.helper.updateMatrixWorld();
    }
}