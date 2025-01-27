import * as THREE from "three";
import App from "./app/app";
import Planet from "./entities/planet";
import TexturedPlanet from "./entities/texturedPlanet";
import Asteroid from "./entities/asteroid";
import Sun from "./entities/sun";

window.onload = function () {
  window.game = new App( // TODO should not be in window, did for debugging
    document.querySelector("#app #rcontainer"),
    document.querySelector("#app #hud"),
    document.location.host,
  );

  var axesHelper = new THREE.AxesHelper(5);
  window.game.graphics.addToScene(axesHelper);

  // Ship.build('1', game.graphics, game.input, game.hud, game.world, game.network)
  //   .then(ship => { game.spawn(ship); });

  let asteroid = Asteroid.build("2", game.graphics, game.world);
  window.game.spawn(asteroid);

  var planet = Planet.build("3", game.graphics);
  window.game.spawn(planet);

  // TexturedPlanet.build("4", game.graphics).then((earth) => {
  //   window.game.spawn(earth);
  // });

  var sun = Sun.build("5", game.graphics, game.world);
  window.game.spawn(sun);

  // var sun = Sun.build("6", game.graphics, game.world);
  // window.game.spawn(sun);

  var alight = new THREE.PointLight();
  alight.position.setX(0, 0, 0);
  game.graphics.addToScene(alight);

  // var size = 100;
  // var divisions = 100;

  // var gridHelper = new THREE.GridHelper(size, divisions);
  // gridHelper.position.set(0, -5, 0);
  // game.graphics.addToScene(gridHelper);
  window.game.start();
};
