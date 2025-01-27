export default class GameObject {
  id: string;
  components: any;

  constructor(id: string, ...components) {
    components.forEach((component) => (component.gameObject = this));
    this.components = components;
    this.id = id;
  }

  update(dt: number) {
    this.components.forEach((component) => component.update(dt));
  }

  attachComponent(component) {
    component.gameObject = this;
    this.components.push(component);
  }

  getComponentThatCan(method: string) {
    return this.components.find((component) => component[method]);
  }

  destroy() {
    this.components.forEach((component) => component.destroy());
  }
}
