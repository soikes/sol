import * as THREE from 'three';
import { SubdivisionModifier } from 'three/examples/jsm/modifiers/SubdivisionModifier.js';

export default class GraphicsEffectsComponent {
    constructor(graphicsComp, graphics) {
        this.graphicsComp = graphicsComp;
        this.graphics = graphics;
        this.setupGlow();
    }

    setupGlow() {
        let modifier = new SubdivisionModifier(1);
        let glowMesh = this.graphicsComp.graphicsObject().clone();
        glowMesh.position.copy(this.graphicsComp.graphicsObject().position);
        glowMesh.scale.set(1.4, 1.4, 1.4); //TODO this doesn't quite look right. Fix the size and shader, and add black transparency
        glowMesh.geometry = modifier.modify(glowMesh.geometry);
        
        let pos = new THREE.Vector3();
        glowMesh.getWorldPosition(pos);
        let viewVec = new THREE.Vector3().subVectors(this.graphics.camera.position, pos);

        let material = new THREE.ShaderMaterial({
            uniforms: {
                viewVec: { value: viewVec } 
            },
            vertexShader: `
                uniform vec3 viewVec;
                varying float intensity;
                void main() {
                    gl_Position = projectionMatrix * modelViewMatrix * vec4(position, 1.0);
                    vec3 actual_normal = vec3(modelMatrix * vec4(normal, 0.0));
                    intensity = pow(dot(normalize(viewVec), actual_normal), 6.0);
                }`,
            fragmentShader: `
                varying float intensity;
                void main() {
                    vec3 glow = vec3(1, 1, 1) * intensity;
                    gl_FragColor = vec4(glow, 0.5);
                }
            `
        });
        glowMesh.material.dispose();
        glowMesh.material = material;
        this.glowMesh = glowMesh;
    }

    glow() {
        this.graphics.addToScene(this.glowMesh);
    }

    stopGlow() {
        this.graphics.removeFromScene(this.glowMesh);
    }

    update() {}
}