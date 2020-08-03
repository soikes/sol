import * as THREE from 'three';
import App from './app';
import Cube from './entities/cube';

window.onload = function() {
  var game = new App();

  var axesHelper = new THREE.AxesHelper(5);
  game.graphics.addToScene(axesHelper);

  var cube = Cube.build(game.graphics, game.input);

  var egeometry = new THREE.SphereGeometry(40, 38, 38);
  var ematerial = new THREE.MeshStandardMaterial({ color: 0x0000ff });
  // ematerial.flatShading = true;
  var earth = new THREE.Mesh(egeometry, ematerial);
  earth.position.set(-80, -80, -40);
  game.graphics.addToScene(earth);

  var rgeometry = new THREE.RingGeometry( 64, 80, 80 );
  var rmaterial = new THREE.MeshPhongMaterial( { color: 0xffff00, side: THREE.DoubleSide } );
  // rmaterial.flatShading = true;
  var rmesh = new THREE.Mesh( rgeometry, rmaterial );
  rmesh.position.set(-80, -80, -40);
  rmesh.rotateX(Math.PI / 2);
  game.graphics.addToScene( rmesh );

  // var light = new THREE.PointLight( 0xff0000, 1, 0 );
  var light = new THREE.PointLight(0xffffff, 3, 0, 2);
  light.position.set(20, 10, 80);
  game.graphics.addToScene(light);

  var sphereSize = 1;
  var pointLightHelper = new THREE.PointLightHelper(light, sphereSize);
  game.graphics.addToScene(pointLightHelper);

  // var alight = new THREE.AmbientLight();
  // game.graphics.addToScene(alight);

  var size = 100;
  var divisions = 100;

  // var gridHelper = new THREE.GridHelper(size, divisions);
  // gridHelper.position.set(0, -5, 0);
  // game.graphics.addToScene(gridHelper);

  game.spawn(cube);
  game.start();
};
