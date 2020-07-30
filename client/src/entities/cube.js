import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import InputComponent from '../components/inputComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';
import CameraFollowComponent from '../components/cameraFollowComponent';

export default class Cube {
  static build(graphics, input) {
    var cubeTf = new TransformComponent(new THREE.Vector3(), new THREE.Vector3(), new THREE.Vector3());
    var geometry = new THREE.CubeGeometry();
    var material = new THREE.MeshLambertMaterial({ color: 0x00ff00 });
    var cubeGfx = new GraphicsComponent(graphics, new THREE.Mesh(geometry, material), cubeTf);
    var cubeFollow = new CameraFollowComponent(cubeTf, graphics);
    var cubeInput = new InputComponent(cubeTf, input);

    return new GameObject(cubeTf, cubeGfx, cubeFollow, cubeInput);
  }
}
