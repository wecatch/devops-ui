import Route from '@ember/routing/route';
import { hash } from 'rsvp';
import {paginationRoute} from 'ember-app-espire/mixins/paging';


export default Route.extend(paginationRoute, {
    queryParams: {
        search_value: {
            refreshModel: true
        }
    },
    model(params){
        return hash({
            computer: this.store.request.get('/v1/computer/list', {data: params}),
            app: this.store.find('app'),
        })
    }
});
