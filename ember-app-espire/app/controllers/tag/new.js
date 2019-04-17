import Controller from '@ember/controller';
import {godForm} from 'ember-easy-orm/mixins/form';

export default Controller.extend(godForm, {
    modelName: "tag",
    actions: {
        cancel(){
            this.transitionToRoute('tag');
        },
        success(){
            this.transitionToRoute('tag');
        }
    }
});
