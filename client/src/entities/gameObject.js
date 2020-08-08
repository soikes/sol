import { Vector3 } from 'three';

export default class GameObject {
  constructor(...components) {
    this.components = components;
    this.id = 1;
  }
  
  update(dt) {
    this.components.forEach(component => component.update(dt));
  }
  
  attachComponent(component) {
    this.components.push(component);
  }

  destroy() {
    this.components.forEach(component => component.destroy());
  }
}