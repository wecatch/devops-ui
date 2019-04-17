import Route from '@ember/routing/route';

export default Route.extend({
    model(){
        return this.store.createRecord('app');
    },
    afterModel(/*model*/){
        return this.store.find('tag').then((data)=>{
            this.controllerFor('app/new').set('tagOptions', data);
        });
    }
});
