import Controller, { inject } from '@ember/controller';
import {set, get} from '@ember/object';
import {computed} from '@ember/object';

export default Controller.extend({
    parentController: inject('app.detail'),
    app: computed(function(){
        return this.get('parentController').model.app;
    })
});
