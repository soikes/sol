import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import TransformComponent from '../components/transformComponent';
import CollisionComponent from '../components/collisionComponent';
import GameObject from './gameObject';
import CollisionBounceComponent from '../components/collisionBounceComponent';
import GravityEffectComponent from '../components/gravityEffectComponent';

export default class Asteroid {
    static build(graphics, world) {
        let asteroidTransform = new TransformComponent(new THREE.Vector3(20, 0, 20), new THREE.Vector3(), new THREE.Vector3());
        let asteroidGeometry = new THREE.SphereGeometry(10, 10, 10);
        let asteroidMaterial = new THREE.MeshStandardMaterial({ color: 0x00f8c7 });
        let asteroidMesh = new THREE.Mesh(asteroidGeometry, asteroidMaterial);
        asteroidMesh.castShadow = true;
        asteroidMesh.receiveShadow = true;

        let asteroidGraphics = new GraphicsComponent(graphics, asteroidMesh, asteroidTransform);
        
        // let asteroidCollisionBounce = new CollisionBounceComponent(0.05);
        // let asteroidCollision = new CollisionComponent(
        //     asteroidGraphics.object(), 
        //     asteroidTransform, 
        //     graphics);
        // asteroidCollision.onCollisionStart(asteroidCollisionBounce.collideStart.bind(asteroidCollisionBounce));
        // asteroidCollision.onCollisionStop(asteroidCollisionBounce.collideStop.bind(asteroidCollisionBounce));
        // world.addCollider(asteroidCollision);

        let asteroidGravityEffect = new GravityEffectComponent(asteroidTransform, 10);
        let asteroidGravityCollision = new CollisionComponent(
            asteroidGraphics.graphicsObject(), 
            asteroidTransform, 
            graphics,
            0.4 //TODO this does not scale, it ADDs X to the size... wrong
        );
        asteroidGravityCollision.onCollisionStart(asteroidGravityEffect.collideStart.bind(asteroidGravityEffect));
        asteroidGravityCollision.onCollisionStop(asteroidGravityEffect.collideStop.bind(asteroidGravityEffect));
        world.addCollider(asteroidGravityCollision);

        return new GameObject(
            asteroidTransform, 
            asteroidGraphics, 
            // asteroidCollision, 
            // asteroidCollisionBounce, 
            asteroidGravityCollision, 
            asteroidGravityEffect
        );
    }
}