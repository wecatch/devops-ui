import Component from '@ember/component';
import { formComponent } from 'ember-easy-orm/mixins/form';
import {set} from '@ember/object';

export default Component.extend(formComponent, {
    init(){
        this._super(...arguments);
        set(this, 'tagOptions', [
            {
                name: 'base',
                value: 'base'
            },
            {
                name: 'business',
                value: 'business'
            },
        ])
    }
});
