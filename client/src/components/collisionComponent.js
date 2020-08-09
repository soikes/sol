import * as THREE from 'three';

export default class CollisionComponent {
    constructor(object, transform, graphics, physics, movable) {
        this.box = new THREE.Box3();
        this.object = object;
        this.box.setFromObject(object);
        this.helper = new THREE.Box3Helper( this.box, 0xffff00 );
        graphics.addToScene(this.helper);

        this.transform = transform;
        this.physics = physics;
        this.colliding = false;
        if (movable === undefined) {
            this.movable = false;
        } else {
            this.movable = movable;
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
        if (this.movable && this.colliding) {
            // this.physics.bounce();
        }
    }
}