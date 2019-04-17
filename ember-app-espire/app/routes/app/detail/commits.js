import Route from '@ember/routing/route';

export default Route.extend({
    model(params){
        return this.store.request.get('/v1/gitlab/commits', {
            data: {'pid': params.repository_id}
        });
    }    
});
