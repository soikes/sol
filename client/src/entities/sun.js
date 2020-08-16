import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';
import CollisionDamageComponent from '../components/collisionDamageComponent';
import CollisionComponent from '../components/collisionComponent';

export default class Sun {
    static build(graphics, world) {
        var sunTransform = new TransformComponent(new THREE.Vector3(-200, 0, -200), new THREE.Vector3(), new THREE.Vector3());

        var sunGeometry = new THREE.SphereGeometry(80, 48, 48);
        var sunMaterial = new THREE.MeshLambertMaterial({ color: 0xffffff, emissive: 0xffffff });
        var sunMesh = new THREE.Mesh(sunGeometry, sunMaterial);
        var sunGraphics = new GraphicsComponent(graphics, sunMesh, sunTransform);
        
        sunGeometry.computeBoundingSphere();
        let sphere = sunGeometry.boundingSphere;
        let points = [
            // sunTransform.pos.clone().add(new THREE.Vector3(sphere.radius, 0, 0)),
            // sunTransform.pos.clone().sub(new THREE.Vector3(sphere.radius, 0, 0)),
            sunTransform.pos.clone().add(new THREE.Vector3(0, sphere.radius / 3, 0)),
            // sunTransform.pos.clone().sub(new THREE.Vector3(0, sphere.radius, 0)),
            // sunTransform.pos.clone().add(new THREE.Vector3(0, 0, sphere.radius)),
            // sunTransform.pos.clone().sub(new THREE.Vector3(0, 0, sphere.radius))
        ];
        for (let i = 0; i < points.length; i++) {
            let light = new THREE.PointLight(0xffffff, 1.3, 0, 2);
            light.position.copy(points[i]);
            light.castShadow = true;
            // light.add(sunGraphics.graphicsObject());
            graphics.addToScene(light);
            var sphereSize = 1;
            var pointLightHelper = new THREE.PointLightHelper(light, sphereSize);
            graphics.addToScene(pointLightHelper);
        }

        var ambLight = new THREE.AmbientLight(0xffffff, 0.05);
        ambLight.position.set(-150, 10, -150);
        ambLight.castShadow = false;
        graphics.addToScene(ambLight);
        
        let sunCollisionDmg = new CollisionDamageComponent(3);
        let sunCollision = new CollisionComponent(
            sunGraphics.graphicsObject(), 
            sunTransform, 
            graphics);
        sunCollision.onCollisionStart(sunCollisionDmg.collideStart.bind(sunCollisionDmg));
        sunCollision.onCollisionStop(sunCollisionDmg.collideStop.bind(sunCollisionDmg));
        world.addCollider(sunCollision);

        return new GameObject(sunTransform, sunGraphics, sunCollision, sunCollisionDmg);
    }
}