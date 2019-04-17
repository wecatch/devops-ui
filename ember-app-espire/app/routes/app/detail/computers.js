import Route from '@ember/routing/route';
// import {get} from '@ember/object';

export default Route.extend({
    queryParams: {
        app_id: {
            refreshModel: true
        },
    },
    app_id: 0,
    model(params, transition){
        let app_id = transition.params["app.detail"].app_id;
        return this.store.request.get('/v1/computer/app', {data: {app_id: app_id}});
    }
});
