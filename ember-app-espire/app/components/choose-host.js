import { inject as service } from '@ember/service';
import Component from '@ember/component';
import {set} from '@ember/object';
import EmberObject from '@ember/object';
import {A} from '@ember/array';
import { observer } from '@ember/object';
import { later } from '@ember/runloop';

export default Component.extend({
    router: service(),
    actions: {
        createDeployJob(rd){
            let app = this.get('app');
            let choosedHosts = this.get('choosedHosts');
            let selectedItem = this.store.createRecord('deploy', {
                commit_id: "",
                app_id: app.app_id,
                status: "new",
                hosts: choosedHosts,
            });
            this.set('selectedItem', selectedItem);
            this.set("modalShow", true);
        },
        chooseAllHost(){
            this.set('isChooseAll', true);
            let choosedHosts = this.get('choosedHosts');
            choosedHosts.clear()
            this.get('hostOptions').forEach((element, index)=>{
                choosedHosts.pushObject(element.private_ip);
                set(element, 'checked', true);
            });
        },
        cancelChooseAll(action, data){
            this.set('modalShow', false);
            this.set('isChooseAll', false);
            this.get('choosedHosts').clear();
            this.get('hostOptions').forEach((element, index)=>{
                set(element, 'checked', false);
            }); 

            later(this, function(){
                if (action == 'create'){
                    let url = `/app/${data.app_id}/detail/deploy/${data.id}`;
                    window.open(url, '_blank');
                }
            }, 500);

        }
    },
    valueObserver: observer('choosedHosts.length', function() {
        if (this.get('choosedHosts').length == 0){
            this.set('isChooseAll', false);
        }
    }),
    init(){
        this._super(...arguments);
        this.set('choosedHosts', A());
        this.set('render', true);
        this.set('modelName', 'deploy');
        this.set('isChooseAll', false);
        this.get('hostOptions').forEach((element, index)=>{
            set(element, 'checked', false);
        });
        this.set('deployModel', null);
        this.set('newDeployJob', false);
    }
});
