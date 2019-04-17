import Route from '@ember/routing/route';
import {paginationRoute} from 'ember-app-espire/mixins/paging';

export default Route.extend(paginationRoute, {
    queryParams: {
        gid: {
            refreshModel: true
        }
    },
    gid: 0,
    model(params){
        return this.store.request.get('/v1/gitlab/projects', {data: params});
    }
});
