import * as THREE from 'three';

export default class AppGraphics {
  constructor(container) {
    this.cameraDist = 20;
    const renderer = new THREE.WebGLRenderer();
    renderer.setSize(window.innerWidth, window.innerHeight);
    container.appendChild(renderer.domElement);
    this.renderer = renderer;

    const scene = new THREE.Scene();
    this.scene = scene;

    const aspect = window.innerWidth / window.innerHeight;
    let camera = new THREE.OrthographicCamera(- this.cameraDist * aspect, this.cameraDist * aspect, this.cameraDist, - this.cameraDist, 1, 1000);
    camera.position.set(this.cameraDist, this.cameraDist, this.cameraDist); // all components equal
    camera.lookAt(this.scene.position);
    this.camera = camera;

    this.handleResize();
  }

  addToScene(obj) {
    this.scene.add(obj);
  }

  cameraFollow(pos) {
    let p = pos.clone();
    let camPos = p.add(new THREE.Vector3(this.cameraDist, this.cameraDist, this.cameraDist));
    this.camera.position.copy(camPos);
    this.camera.lookAt(pos);
  }

  handleResize() {
    window.addEventListener('resize', () => {
      let aspect = window.innerWidth / window.innerHeight;
      this.renderer.setSize(window.innerWidth, window.innerHeight);
      this.camera.left = -this.cameraDist * aspect;
      this.camera.right = this.cameraDist * aspect;
      this.camera.top = this.cameraDist;
      this.camera.bottom = -this.cameraDist;
      this.camera.near = 1;
      this.camera.far = 1000;
      this.camera.updateProjectionMatrix();
    }, false);
  }

  render() {
    this.renderer.render(this.scene, this.camera);
  }
}