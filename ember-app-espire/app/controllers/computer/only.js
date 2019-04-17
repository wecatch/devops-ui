import Controller from '@ember/controller';
import {A} from '@ember/array';
import {paginationController} from 'ember-app-espire/mixins/paging';
import {computed, set} from '@ember/object';

export default Controller.extend(paginationController, {
    queryParams: ['search_value'],
    search_value: null,
    actions: {
        saveApp(){
            let choosedApp = this.get('choosedApp')
            let host_id = A();
            this.get('choosedHost').forEach((element, index) =>{
                host_id.pushObject(element.host_id);
            });
            if (choosedApp.length > 0 && host_id.length > 0) {
                this.store.request.post('/v1/computer/role', {data: {app_id: choosedApp, host_id: host_id}}).then(()=>{
                    this.reset();
                });
            }

            this.reset();

        },
        cancelApp(){
            this.reset();
        },
        displayModal(rd){
            this.set('display', true);
        },
        searchApp(searchValue){
            this.set('chooseAll', false);
            let page = this.page;
            this.transitionToRoute({
                queryParams: {
                    page: page,
                    search_value: searchValue.target.value,
                }
            });
        },
        sync(){
            this.set('syncLoading', true);
            this.store.request.post('/v1/cloud/sync', {timeout: 20000}).then(()=>{
                this.set('syncLoading', false);
            },()=>{
                this.set('syncLoading', false);
            });
        },
        chooseHost(checked, rd){
            if(checked){
                this.get('choosedHost').pushObject(rd);
            }else {
                this.get('choosedHost').removeObject(rd);
            }
        },
        chooseAllHost(){
            this.set('chooseAll', true);
            this.get('choosedHost').clear();
            this.get('model.computer').forEach((element, index) => {
                set(element, 'checked', true);
                this.get('choosedHost').pushObject(element);
            });
        },
        cancelChooseAllHost(){
            this.set('chooseAll', false);
            this.get('choosedHost').clear();
            this.get('model.computer').forEach((element, index) => {
                set(element, 'checked', false);
            });
        }
    },
    reset(){
        this.set('display', false);
        this.get('choosedHost').clear();
        this.get('model.app').forEach((element, index)=>{
            set(element, 'checked', false);
        });
        this.get('model.computer').forEach((element, index) => {
            set(element, 'checked', false);
        });
    },
    displayChoosedHost: computed('choosedHost.[]', function(){
        let value = "";
        this.get('choosedHost').forEach((element, index) => {
            value += element.private_ip;
        });
        return value;
    }),
    init(){
        this._super(...arguments);
        this.set('choosedApp',A());
        this.set('choosedHost',A());
    }
});
