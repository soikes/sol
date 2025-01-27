import * as THREE from "three";

function main() {
  const canvas: HTMLCanvasElement | null = document.querySelector("#app");
  if (!canvas) {
    throw new Error("No canvas element found.");
  }
  const renderer: THREE.WebGLRenderer = new THREE.WebGLRenderer({
    antialias: true,
    canvas,
  });

  const fov: number = 75;
  const aspect: number = 2;
  const near: number = 0.1;
  const far: number = 5;
  const camera: THREE.PerspectiveCamera = new THREE.PerspectiveCamera(
    fov,
    aspect,
    near,
    far,
  );

  camera.position.z = 2;

  const scene: THREE.Scene = new THREE.Scene();

  const boxWidth: number = 1;
  const boxHeight: number = 1;
  const boxDepth: number = 1;
  const geometry: THREE.BoxGeometry = new THREE.BoxGeometry(
    boxWidth,
    boxHeight,
    boxDepth,
  );

  const material: THREE.MeshBasicMaterial = new THREE.MeshPhongMaterial({
    color: 0x44aa88,
  });

  const cube: THREE.Mesh = new THREE.Mesh(geometry, material);
  scene.add(cube);

  const color = 0xffffff;
  const intensity = 3;
  const light = new THREE.DirectionalLight(color, intensity);
  light.position.set(-1, 2, 4);
  scene.add(light);

  renderer.render(scene, camera);
  requestAnimationFrame(renderMesh(cube, renderer, scene, camera));
}

function renderMesh(
  obj: THREE.Mesh,
  renderer: THREE.Renderer,
  scene: THREE.Scene,
  camera: THREE.PerspectiveCamera,
) {
  const render = function (time: number) {
    time *= 0.001;

    const canvas = renderer.domElement;
    const pixelRatio = window.devicePixelRatio;
    const width = (canvas.clientWidth * pixelRatio) | 0;
    const height = (canvas.clientHeight * pixelRatio) | 0;
    const needResize = canvas.width !== width || canvas.height !== height;
    if (needResize) {
      renderer.setSize(width, height, false);
      camera.aspect = canvas.clientWidth / canvas.clientHeight;
      camera.updateProjectionMatrix();
    }

    obj.rotation.x = time;
    obj.rotation.y = time;

    renderer.render(scene, camera);

    requestAnimationFrame(render);
  };
  return render;
}

main();
