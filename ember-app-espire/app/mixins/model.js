import Mixin from '@ember/object/mixin';
import {
    get
} from '@ember/object';
import {
    inject
} from '@ember/service';

import model, {
    DS
} from 'ember-easy-orm/mixins/model';

// model.reopen({
//     dataRootKey: 'res'
// });

export {
    DS
};

export default Mixin.create(model, {
    flashMessages: inject(),
    flashLoading: inject(),
    RESTSerializer: function(data) {
        if (data && data.code !== undefined) {
            if (data.code === 0) {
                return data['resp'];
            } else {
                if (data.code === 2) {
                    return window.location.reload();
                }
                if (data.msg) {
                    get(this, 'flashMessages').warning(data.msg);
                    throw String(data.msg);
                }else {
                    throw String(data);
                }
                
            }
        } else {
            throw String(data);
        }
    },
    init() {
        this._super(...arguments);
        this.ajaxSettings = {
            traditional: false,
            dataType: 'json',
            timeout: 10000,
            contentType: 'application/json'
        }
        this.on('ajaxError', (error, ajaxSettings, jqxhr) => { // jshint ignore:line
            get(this, 'flashMessages').warning(error);
        });

        this.on('ajaxSuccess', () => {
            get(this, 'flashMessages').success('success');
        });

        this.on('RESTSerializerError', (e) => {
            get(this, 'flashMessages').warning(e || '数据解析有误，请确保网络连接正常');
        });

        this.on('ajaxStart', () => {
            get(this, 'flashLoading').trigger(this, true);
        });

        this.on('ajaxDone', () => {
            get(this, 'flashLoading').trigger(this, false);
        });
    }

});