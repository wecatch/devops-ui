import Route from '@ember/routing/route';

export default Route.extend({
    model(params){
        return this.store.findOne('app', params.app_id);
    },
    afterModel(/*model*/){
        return this.store.find('tag').then((data)=>{
            this.controllerFor('app/edit').set('tagOptions', data);
        });
    }
});
