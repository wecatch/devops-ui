import Route from '@ember/routing/route';
import {paginationRoute} from 'ember-app-espire/mixins/paging';

export default Route.extend(paginationRoute,{
    model(params){
        return this.store.request.get('/v1/consul/service/all');
    }
});
