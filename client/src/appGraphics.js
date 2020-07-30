import * as THREE from 'three';
import GraphicsComponent from './components/graphicsComponent';

export default class AppGraphics {
  constructor() {
    const renderer = new THREE.WebGLRenderer();
    renderer.setSize(window.innerWidth, window.innerHeight);
    document.body.appendChild(renderer.domElement);
    this.renderer = renderer;

    const scene = new THREE.Scene();
    this.scene = scene;

    const aspect = window.innerWidth / window.innerHeight;
    let d = 20;
    let camera = new THREE.OrthographicCamera(- d * aspect, d * aspect, d, - d, 1, 1000);
    camera.position.set(20, 20, 20); // all components equal
    camera.lookAt(this.scene.position);
    this.camera = camera;

    this.handleResize();
  }

  addToScene(obj) {
    this.scene.add(obj);
  }

  handleResize() {
    window.addEventListener('resize', () => {
      let d = 20;
      let aspect = window.innerWidth / window.innerHeight;
      this.renderer.setSize(window.innerWidth, window.innerHeight);
      this.camera.left = -d * aspect;
      this.camera.right = d * aspect;
      this.camera.top = d;
      this.camera.bottom = -d;
      this.camera.near = 1;
      this.camera.far = 1000;
      this.camera.updateProjectionMatrix();
    }, false);
  }

  render() {
    this.renderer.render(this.scene, this.camera);
  }
}