import * as THREE from 'three';
import App from './app';
import Ship from './entities/ship';
import Planet from './entities/planet';

window.onload = function() {
  var game = new App(
    document.querySelector("#app #rcontainer"), 
    document.querySelector("#app #hud")
  );

  var axesHelper = new THREE.AxesHelper(5);
  game.graphics.addToScene(axesHelper);

  Ship.build(game.graphics, game.input, game.hud).then(ship => { game.spawn(ship); });

  var planet = Planet.build(game.graphics);

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

  game.spawn(planet);

  game.start();
};
