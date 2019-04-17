import Component from '@ember/component';
import {computed} from '@ember/object';

export default Component.extend({
    tagName: "button",
    classNameBindings: [':ui',':basic',':small','isLoading',':grey',':button'],
    isLoading: computed(function(){
        return this.get('flashLoading').loading;
    }),
    click(event){
        let appId = this.get('app_id');
        let hostId = this.get('host_id');
        this.store.request.delete('/v1/computer/role', {data: {app_id:[appId], host_id: [hostId]}}).then(()=>{
            this.get('model').removeObject(this.get('rd'));
        });
    }
});
