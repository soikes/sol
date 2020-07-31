import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import InputComponent from '../components/inputComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';
import CameraFollowComponent from '../components/cameraFollowComponent';
import SpinComponent from '../components/spinComponent';
import PhysicsComponent from '../components/physicsComponent';

export default class Cube {
  static build(graphics, input) {
    var cubeTf = new TransformComponent(new THREE.Vector3(), new THREE.Vector3(), new THREE.Vector3());
    var geometry = new THREE.CubeGeometry();
    var material = new THREE.MeshLambertMaterial({ color: 0x00ff00 });
    var cubeGfx = new GraphicsComponent(graphics, new THREE.Mesh(geometry, material), cubeTf);
    var cubeFollow = new CameraFollowComponent(cubeTf, graphics);
    var cubePhys = new PhysicsComponent(new THREE.Vector3(), 0.2, 0.005, 0.05, cubeTf);
    var cubeInput = new InputComponent(cubeTf, input, cubePhys);
    // var cubeSpin = new SpinComponent(cubeTf);

    return new GameObject(cubeTf, cubeGfx, cubeFollow, cubeInput, cubePhys);
  }
}
