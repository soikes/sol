import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import InputComponent from '../components/inputComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';
import CameraFollowComponent from '../components/cameraFollowComponent';
import SpinComponent from '../components/spinComponent';
import PhysicsComponent from '../components/physicsComponent';

import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader.js';

export default class Ship {
  static build(graphics, input) {
    return new Promise((resolve, reject) => {
      var shipGfx;
      var shipTf = new TransformComponent(new THREE.Vector3(), new THREE.Vector3(), new THREE.Vector3());
      var loader = new GLTFLoader();
      loader.load( 'assets/ship.glb', function ( gltf ) {
        shipGfx = new GraphicsComponent(graphics, gltf.scene, shipTf);
        var shipFollow = new CameraFollowComponent(shipTf, graphics);
        var shipPhys = new PhysicsComponent(new THREE.Vector3(), 0.2, 0.005, 0.05, shipTf);
        var shipInput = new InputComponent(shipTf, input, shipPhys);
        resolve(new GameObject(shipTf, shipGfx, shipFollow, shipInput, shipPhys));
      }, undefined, function ( error ) {
        console.error( error );
        reject(error);
      } );
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
