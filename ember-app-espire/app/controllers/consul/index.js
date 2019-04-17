import Controller from '@ember/controller';
import {paginationController} from 'ember-app-espire/mixins/paging';
import {computed} from '@ember/object';


export default Controller.extend(paginationController, {
    name: "",
    filterModel: computed('name', function(){
        let name = this.name;
        let model = this.model;
        let ret = {}
        if(name != ""){
            if(model[name]) {
                ret[name] = model[name]
            }
            return ret;
        }else {
            return model;
        }
    })
});
