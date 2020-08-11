import * as THREE from 'three';
import GraphicsComponent from '../components/graphicsComponent';
import TransformComponent from '../components/transformComponent';
import GameObject from './gameObject';

export default class Planet {
    static build(graphics) {
        var pt = new TransformComponent(new THREE.Vector3(-80, -100, -40), new THREE.Vector3(), new THREE.Vector3());
        var pg = new THREE.SphereGeometry(40, 38, 38);
        var pm = new THREE.MeshStandardMaterial({ color: 0xfff8c7 });
        var pms = new THREE.Mesh(pg, pm);
        pms.castShadow = true;
        pms.receiveShadow = true;

        var pgx = new GraphicsComponent(graphics, pms, pt);

        var rt = new TransformComponent(new THREE.Vector3(-80, -100, -40), new THREE.Vector3(Math.PI / 2, 0, 0), new THREE.Vector3());
        var rg = new THREE.RingGeometry( 64, 120, 80 );
        var rm = new THREE.MeshPhongMaterial( { color: 0xf5f3e6, side: THREE.DoubleSide } );
        var rms = new THREE.Mesh(rg, rm);
        rms.receiveShadow = true;

        var sgx = new GraphicsComponent(graphics, rms, rt, new THREE.Vector3());

        return new GameObject(pt, pgx, rt, sgx);
    }
}