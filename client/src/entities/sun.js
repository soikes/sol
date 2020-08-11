import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';

export default class Sun {
    static build(graphics) {
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
            let light = new THREE.PointLight(0xffffff, 1, 0, 2);
            light.position.copy(points[i]);
            light.castShadow = true;
            // light.add(sunGraphics.object());
            graphics.addToScene(light);
            var sphereSize = 1;
            var pointLightHelper = new THREE.PointLightHelper(light, sphereSize);
            graphics.addToScene(pointLightHelper);
        }

        var ambLight = new THREE.AmbientLight(0xffffff, 0.05);
        ambLight.position.set(-150, 10, -150);
        ambLight.castShadow = false;
        graphics.addToScene(ambLight);
        
        

        return new GameObject(sunTransform, sunGraphics);
    }
}