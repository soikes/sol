import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import GameObject from './gameObject';

export default class Cube {
  static build(graphics) {
    var geometry = new THREE.CubeGeometry();
    var material = new THREE.MeshLambertMaterial({ color: 0x00ff00 });
    var cubeGfx = new GraphicsComponent(graphics, new THREE.Mesh(geometry, material));
    return new GameObject(cubeGfx);
  }
}