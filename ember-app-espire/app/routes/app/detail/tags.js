import Route from '@ember/routing/route';


export default Route.extend({
    model(params, /*transition*/){
        // let app_id = transition.params["app.detail"].app_id;
        return this.store.request.get('/v1/gitlab/tags', {
            data: {'pid': params.repository_id}
        });
    }
});
