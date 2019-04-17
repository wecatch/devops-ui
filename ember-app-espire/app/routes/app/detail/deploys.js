import Route from '@ember/routing/route';

export default Route.extend({
    model(_, transition){
        let app_id = transition.params["app.detail"].app_id;
        return this.store.find("deploy", {app_id: app_id});
    }
});
