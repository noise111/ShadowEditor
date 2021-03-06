/*
 * Copyright 2017-2020 The ShadowEditor Authors. All rights reserved.
 *
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 * 
 * For more information, please visit: https://github.com/tengge1/ShadowEditor
 * You can also visit: https://gitee.com/tengge1/ShadowEditor
 */
import BaseHelper from '../BaseHelper';
import VolumePointLightHelper from './VolumePointLightHelper';
import global from '../../global';

/**
 * 点光源帮助器
 * @author tengge / https://github.com/tengge1
 */
function PointLightHelpers() {
    BaseHelper.call(this);

    this.helpers = [];
}

PointLightHelpers.prototype = Object.create(BaseHelper.prototype);
PointLightHelpers.prototype.constructor = PointLightHelpers;

PointLightHelpers.prototype.start = function () {
    global.app.on(`objectAdded.${this.id}`, this.onObjectAdded.bind(this));
    global.app.on(`objectRemoved.${this.id}`, this.onObjectRemoved.bind(this));
    global.app.on(`objectChanged.${this.id}`, this.onObjectChanged.bind(this));
    global.app.on(`storageChanged.${this.id}`, this.onStorageChanged.bind(this));
};

PointLightHelpers.prototype.stop = function () {
    global.app.on(`objectAdded.${this.id}`, null);
    global.app.on(`objectRemoved.${this.id}`, null);
    global.app.on(`objectChanged.${this.id}`, null);
    global.app.on(`storageChanged.${this.id}`, null);
};

PointLightHelpers.prototype.onObjectAdded = function (object) {
    if (!object.isPointLight) {
        return;
    }

    var helper = new VolumePointLightHelper(object, 1);

    helper.visible = global.app.storage.showPointLight;

    this.helpers.push(helper);

    global.app.editor.sceneHelpers.add(helper);
};

PointLightHelpers.prototype.onObjectRemoved = function (object) {
    if (!object.isPointLight) {
        return;
    }

    var index = this.helpers.findIndex(n => {
        return n.light === object;
    });

    if (index === -1) {
        return;
    }

    global.app.editor.sceneHelpers.remove(this.helpers[index]);

    this.helpers[index].dispose();

    this.helpers.splice(index, 1);
};

PointLightHelpers.prototype.onObjectChanged = function (object) {
    if (!object.isPointLight) {
        return;
    }

    var index = this.helpers.findIndex(n => {
        return n.light === object;
    });

    if (index === -1) {
        return;
    }

    this.helpers[index].update();
};

PointLightHelpers.prototype.onStorageChanged = function (key, value) {
    if (key !== 'showPointLight') {
        return;
    }

    this.helpers.forEach(n => {
        n.visible = value;
    });
};

export default PointLightHelpers;