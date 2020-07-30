import { Vector3 } from 'three';

export default class GameObject {
  constructor(...components) {
    this.components = components;
    this.id = 1;
  }
  
  update() {
    this.components.forEach(component => component.update());
  }
  
  attachComponent(component) {
    this.components.push(component);
  }
}