import Route from '@ember/routing/route';

export default Route.extend({
    model(){
        return this.store.request.get('/v1/gitlab/groups');
    }
});
