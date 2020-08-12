import { Vector3 } from 'three';

export default class GameObject {
  constructor(...components) {
    components.forEach(component => component.gameObject = this);
    this.components = components;
    this.id = 1;
  }
  
  update(dt) {
    this.components.forEach(component => component.update(dt));
  }
  
  attachComponent(component) {
    component.gameObject = this;
    this.components.push(component);
  }

  getComponentThatCan(method) {
    return this.components.find(component => component[method]);
  }

  destroy() {
    this.components.forEach(component => component.destroy());
  }
}