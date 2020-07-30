import * as THREE from 'three';
import App from './app';
import Cube from './entities/cube'

window.onload = function() {
  var game = new App();

  var axesHelper = new THREE.AxesHelper(5);
  game.graphics.addToScene(axesHelper);

  var cube = Cube.build(game.graphics);

  var egeometry = new THREE.SphereGeometry(40, 38, 38);
  var ematerial = new THREE.MeshStandardMaterial({ color: 0x0000ff });
  var earth = new THREE.Mesh(egeometry, ematerial);
  earth.position.set(-80, -80, -40);
  game.graphics.addToScene(earth);

  // var light = new THREE.PointLight( 0xff0000, 1, 0 );
  var light = new THREE.PointLight();
  light.position.set(10, 0, 0);
  game.graphics.addToScene(light);

  var alight = new THREE.AmbientLight();
  game.graphics.addToScene(alight);

  var size = 100;
  var divisions = 100;

  var gridHelper = new THREE.GridHelper(size, divisions);
  gridHelper.position.set(0, -5, 0);
  game.graphics.addToScene(gridHelper);

  game.spawn(cube);
  game.start();
};