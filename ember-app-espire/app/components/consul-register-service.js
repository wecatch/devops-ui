import Component from '@ember/component';
import {computed} from '@ember/object';

export default Component.extend({
    tagName: "button",
    classNameBindings: [':ui',':basic',':small','isLoading','color',':button'],
    isLoading: computed(function(){
        return this.get('flashLoading').loading;
    }),
    color: computed('register_status', function(){
        if (this.get('register_status')){
            return 'red';
        }else {
            return 'green';
        }
    }),
    click(event){
        let appId = this.get('app_id');
        let hostId = this.get('host_id');
        let app_name = this.get('app_name')
        let port = this.get('port')
        let private_ip = this.get('private_ip')
        let tag = this.get('tag')
        let service_id = this.get('service_id') || ""

        let form = {
            app_id:appId, 
            host_id: hostId,
            name: app_name,
            port: port,
            private_ip: private_ip,
            tag: tag,
            service_id: service_id,
        }

        if (!form.name) {
            form.name = form.service_id.split("-")[0];
        }

        if (!this.get('register_status')){
            this.store.request.put('/v1/consul/app/register', {data: form}).then(()=>{
                this.set('register_status', 1);
            });
        }

        if (this.get('register_status')){
            this.store.request.put('/v1/consul/app/unregister', {data: form}).then(()=>{
                this.set('register_status', 0);
            });
        }
    }
});
