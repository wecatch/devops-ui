import Controller from '@ember/controller';

export default Controller.extend({
    actions: {
        delete(rd) {
            this.store.request.delete('/v1/', {data: {
                
            }}).then(()=>{
                this.model.removeObject(rd);
            });
        }
    }
});
