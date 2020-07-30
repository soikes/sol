export default class CameraFollowComponent {
  constructor(entity, graphics) {
    this.entity = entity;
    this.graphics = graphics;
    this.following = true;
  }

  update() {
    if (this.following) {
      this.graphics.camera.obj.lookAt(this.entity.position);
    }
  }

  follow() {
    this.following = true;
  }

  unfollow() {
    this.following = false;
  }
}