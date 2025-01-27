import * as THREE from "three";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
import { DRACOLoader } from "three/examples/jsm/loaders/DRACOLoader";
import GraphicsComponent from "../components/graphicsComponent";
import TransformComponent from "../components/transformComponent";
import GameObject from "./gameObject";

export default class TexturedPlanet {
  static build(id, graphics) {
    return new Promise((resolve, reject) => {
      var etf = new TransformComponent(
        new THREE.Vector3(-80, -100, -40),
        new THREE.Vector3(),
        new THREE.Vector3(),
      );

      let loader = new GLTFLoader();
      const dracoLoader = new DRACOLoader();
      dracoLoader.setDecoderPath("/examples/jsm/libs/draco/");
      loader.setDRACOLoader(dracoLoader);

      loader.load(
        "assets/earth.glb",
        function (gltf) {
          gltf.scene.scale.set(14, 14, 14);
          let earthGfx = new GraphicsComponent(graphics, gltf.scene, etf);
          resolve(new GameObject(id, etf, earthGfx));
        },
        undefined,
        function (error) {
          console.error(error);
          reject(error);
        },
      );
    });
  }
}
