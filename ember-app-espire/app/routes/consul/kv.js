import Route from '@ember/routing/route';
import {paginationRoute} from 'ember-app-espire/mixins/paging';

export default Route.extend(paginationRoute,{
    queryParams: {
        dc: {
            refreshModel: true
        },
        prefix: {
            refreshModel: true
        },
        separator: {
            refreshModel: true
        },
    },
    dc: "dc1",
    prefix: "",
    separator: "/",
    model(params){
        return this.store.request.get('/v1/consul/kv/keys', params)
    }
});
