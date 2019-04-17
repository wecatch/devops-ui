import Route from '@ember/routing/route';
import {paginationRoute} from 'ember-app-espire/mixins/paging';

export default Route.extend(paginationRoute, {
    queryParams: {
        app_name: {
            refreshModel: true
        }
    },
    model(params){
        return this.store.request.get('/v1/computer/app', {data: params});
    }
});
