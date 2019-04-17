import Controller, { inject } from '@ember/controller';
import {set, get} from '@ember/object';
import {godForm} from 'ember-easy-orm/mixins/form';

export default Controller.extend(godForm, {
    modelName: "deploy",
    parentController: inject('app.detail'),
    actions: {
        add(rd){
            let app = get(this, 'parentController').model.app;
            set(this, 'selectedItem', this.store.createRecord('deploy', {commit_id: rd.short_id, app_id: app.id, status: "new"}));
            this.set('modalShow', true);
        }
    }
});
