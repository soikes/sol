import * as THREE from 'three';
import App from './app/app';
import Ship from './entities/ship';
import Planet from './entities/planet';
import Asteroid from './entities/asteroid';
import Sun from './entities/sun';

window.onload = function() {
  var game = new App(
    document.querySelector("#app #rcontainer"), 
    document.querySelector("#app #hud"),
    document.location.host
  );

  var axesHelper = new THREE.AxesHelper(5);
  game.graphics.addToScene(axesHelper);

  // Ship.build('1', game.graphics, game.input, game.hud, game.world, game.network)
  //   .then(ship => { game.spawn(ship); });
  
  let asteroid = Asteroid.build('2', game.graphics, game.world);
  game.spawn(asteroid);
  
  var planet = Planet.build('3', game.graphics);
  game.spawn(planet);

  var sun = Sun.build('4', game.graphics, game.world);
  game.spawn(sun);

  // var alight = new THREE.AmbientLight();
  // game.graphics.addToScene(alight);

  // var size = 100;
  // var divisions = 100;

  // var gridHelper = new THREE.GridHelper(size, divisions);
  // gridHelper.position.set(0, -5, 0);
  // game.graphics.addToScene(gridHelper);
  game.start();
};
