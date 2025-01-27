import * as THREE from "three";
import GraphicsComponent from "../components/graphicsComponent";
import TransformComponent from "../components/transformComponent";
import GameObject from "./gameObject";
import PhysicsComponent from "../components/physicsComponent";
import AppWorld from "../app/appWorld";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader.js";
import HealthComponent from "../components/healthComponent";
import CollisionComponent from "../components/collisionComponent";

export default class Ship {
  static build(id: string, graphics: GraphicsComponent, world: AppWorld) {
    return new Promise((resolve, reject) => {
      let shipTf = new TransformComponent(
        new THREE.Vector3(),
        new THREE.Vector3(),
        new THREE.Vector3(),
      );

      let shipHealth = new HealthComponent(0, 100, 100);
      let shipPhys = new PhysicsComponent(
        new THREE.Vector3(),
        15,
        8,
        0.05,
        shipTf,
      );
      let loader = new GLTFLoader();
      loader.load(
        "assets/ship.glb",
        function (gltf) {
          let shipGfx = new GraphicsComponent(graphics, gltf.scene, shipTf);
          let shipCollision = new CollisionComponent(
            shipGfx.graphicsObject(),
            shipTf,
            graphics,
          );
          world.addCollider(shipCollision);
          resolve(
            new GameObject(
              id,
              shipTf,
              shipGfx,
              shipPhys,
              shipHealth,
              shipCollision,
            ),
          );
        },
        undefined,
        function (error) {
          console.error(error);
          reject(error);
        },
      );
    });

    // var geometry = new THREE.CubeGeometry();
    // var material = new THREE.MeshLambertMaterial({ color: 0x00ff00 });
    // var cubeGfx = new GraphicsComponent(graphics, new THREE.Mesh(geometry, material), cubeTf);
    // var sgeometry = new THREE.SphereGeometry();
    // var smaterial = new THREE.MeshLambertMaterial({ color: 0xff0000 });
    // var sphereGfx = new GraphicsComponent(graphics, new THREE.Mesh(sgeometry, smaterial), cubeTf, new THREE.Vector3(-10,0,-10));
    // var cubeSpin = new SpinComponent(cubeTf);
  }
}
