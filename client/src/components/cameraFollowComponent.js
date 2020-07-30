export default class CameraFollowComponent {
  constructor(transform, graphics) {
    this.transform = transform;
    this.graphics = graphics;
    this.following = true;
  }

  update() {
    if (this.following) {
      this.graphics.cameraFollow(this.transform.pos);
    }
  }

  follow() {
    this.following = true;
  }

  unfollow() {
    this.following = false;
  }
}