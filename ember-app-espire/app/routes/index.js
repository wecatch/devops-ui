import EmberObject from '@ember/object';
import Route from '@ember/routing/route';
import {paginationRoute} from 'ember-app-espire/mixins/paging';
import {A} from '@ember/array';

export default Route.extend(paginationRoute,{
    model(params){
        return this.store.request.get('/v1/computer/app/group', {data: params});
    },
    setupController(controller, model) {
        let ret = {
        }
        for (var i = 0; i < model.length; i++) {
            let obj = model.objectAt(i);
            if (!ret[obj.name]) {
                ret[obj.name] = {
                    host: A(),
                    app: obj
                };
            }
            ret[obj.name].host.pushObject(obj);
        }
        controller.set('model', ret);
    }
});
