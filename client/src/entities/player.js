import * as THREE from "three";
import GraphicsComponent from "../components/graphicsComponent";
import InputComponent from "../components/inputComponent";
import TransformComponent from "../components/transformComponent";
import GameObject from "./gameObject";
import CameraFollowComponent from "../components/cameraFollowComponent";
import PhysicsComponent from "../components/physicsComponent";
import Observable from "../util/observable.js";

import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader.js";
import HealthComponent from "../components/healthComponent";
import CollisionComponent from "../components/collisionComponent";

export default class Player {
  static build(id, graphics, input, hud, world, network) {
    return new Promise((resolve, reject) => {
      let shipTf = new TransformComponent(
        new THREE.Vector3(),
        new THREE.Vector3(),
        new THREE.Vector3(),
      );
      let hudPosObserver = new Observable();
      hudPosObserver.subscribe(hud.updatePos.bind(hud));
      shipTf.observe(hudPosObserver);

      let shipHealth = new HealthComponent(0, 100, 100);
      let hudHealthObserver = new Observable();
      hudHealthObserver.subscribe(hud.updateHealth.bind(hud));
      shipHealth.observe(hudHealthObserver);

      let shipFollow = new CameraFollowComponent(shipTf, graphics);
      let shipPhys = new PhysicsComponent(
        new THREE.Vector3(),
        15,
        8,
        0.05,
        shipTf,
      );
      let shipInput = new InputComponent(shipTf, input, shipPhys, network);
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
              shipFollow,
              shipInput,
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
