import * as THREE from 'three';

export default class CollisionComponent {
    constructor(object, transform, graphics, sizeFactor) {
        this.box = new THREE.Box3();
        this.object = object;
        this.box.setFromObject(object);
        this.sizeFactor = sizeFactor;
        this.scaledSize = new THREE.Vector3();
        if (this.sizeFactor) {
            this.box.expandByVector(this.box.getSize(this.scaledSize).multiplyScalar(this.sizeFactor));
        }
        this.helper = new THREE.Box3Helper(this.box, 0xffff00);
        graphics.addToScene(this.helper);

        this.transform = transform;
        this.colliding = false;
        this.collideStartHandlers = [];
        this.collideStopHandlers = [];
    }

    onCollisionStart(handler) {
        this.collideStartHandlers.push(handler);
    }

    onCollisionStop(handler) {
        this.collideStopHandlers.push(handler);
    }

    collideStart(c2) {
        this.collideStartHandlers.forEach(handler => {
            handler(c2);
        });
    }

    collideStop(c2) {
        this.collideStopHandlers.forEach(handler => {
            handler(c2);
        });
    }

    computeBounding() {
        this.box.setFromObject(this.object);
    }

    intersects(collision) {
        if (this.box.intersectsBox(collision.box)) {
            return true;
        }
        return false;
    }

    update() {
        this.box.setFromObject(this.object);
        if (this.sizeFactor) {
            this.box.expandByVector(this.box.getSize(this.scaledSize).multiplyScalar(this.sizeFactor));
        }
        this.helper.updateMatrixWorld();
    }
}