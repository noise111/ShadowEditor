/*
 * Copyright 2003-2006, 2009, 2017, United States Government, as represented by the Administrator of the
 * National Aeronautics and Space Administration. All rights reserved.
 *
 * The NASAWorldWind/WebWorldWind platform is licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
attribute vec4 vertexPoint;
attribute vec4 vertexTexCoord;

uniform mat4 mvpMatrix;
uniform mat4 texSamplerMatrix;
uniform mat4 texMaskMatrix;

uniform int column;
uniform int row;
uniform int level;
uniform sampler2D heightmap;

varying vec2 texSamplerCoord;
varying vec2 texMaskCoord;

int add(int x, int y, int z) {
    return x + y + z;
}

void main() {
    add(column, row, level);
    gl_Position = mvpMatrix * vertexPoint;
    /* Transform the vertex texture coordinate into sampler texture coordinates. */
    texSamplerCoord = (texSamplerMatrix * vertexTexCoord).st;
    /* Transform the vertex texture coordinate into mask texture coordinates. */
    texMaskCoord = (texMaskMatrix * vertexTexCoord).st;
}