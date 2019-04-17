import {get} from '@ember/object';
import { inject } from '@ember/service';
import { defineProperty } from '@ember/object';
import {A} from '@ember/array';


export function initialize(appInstance) {
    let store = appInstance.lookup('service:store');
    // https://github.com/emberjs/ember.js/issues/16519
    defineProperty(store, 'flashMessages', inject('flashMessages'));
    defineProperty(store, 'flashLoading', inject('flashLoading'));
    store.ajaxSettings = {
        traditional: false,
        dataType: 'json',
        timeout: 10000,
        contentType: 'application/json',
    };

    store.RESTSerializer=function(data){
        if(data && data.code===0){
            return data['resp'];
        }else if (data && data.code === 2) {
            return window.location.reload();
        } else {
            let msg = data && data.message || data;
            get(store, 'flashMessages').warning(msg);
            throw String(msg);
        }
    };

    store.on('ajaxError', (/*error, ajaxSettings, jqxhr*/) => { // jshint ignore:line
        get(store, 'flashMessages').warning('网络请求错误，请重试');
    });

    store.on('ajaxSuccess', () => {
        get(store, 'flashMessages').success('success');
    });

    store.on('RESTSerializerError', (e) => {
        get(store, 'flashMessages').warning(e || '数据解析有误，请确保网络连接正常');
    });

    store.on('ajaxStart', () => {
        get(store, 'flashLoading').trigger(store, true);
    });

    store.on('ajaxDone', () => {
        get(store, 'flashLoading').trigger(store, false);
    });

    store.set('needSerializedMethod', A(['put', 'post', 'delete']));
}

export default {
    name: 'reopen-store',
    initialize: initialize
};
