import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import TransformComponent from '../components/transformComponent';
import CollisionComponent from '../components/collisionComponent';
import GameObject from './gameObject';
import CollisionDamageComponent from '../components/collisionDamageComponent';

export default class Asteroid {
    static build(graphics, world) {
        let asteroidTransform = new TransformComponent(new THREE.Vector3(20, 0, 20), new THREE.Vector3(), new THREE.Vector3());
        let asteroidGeometry = new THREE.SphereGeometry(10, 10, 10);
        let asteroidMaterial = new THREE.MeshStandardMaterial({ color: 0x00f8c7 });
        let asteroidMesh = new THREE.Mesh(asteroidGeometry, asteroidMaterial);
        asteroidMesh.castShadow = true;
        asteroidMesh.receiveShadow = true;

        let asteroidGraphics = new GraphicsComponent(graphics, asteroidMesh, asteroidTransform);
        let asteroidCollisionDmg = new CollisionDamageComponent(1);
        let asteroidCollision = new CollisionComponent(
            asteroidGraphics.object(), 
            asteroidTransform, 
            graphics, 
            null, 
            asteroidCollisionDmg.onCollide.bind(asteroidCollisionDmg));
        world.addCollider(asteroidCollision);

        return new GameObject(asteroidTransform, asteroidGraphics, asteroidCollision, asteroidCollisionDmg);
    }
}