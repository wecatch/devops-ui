import Controller from '@ember/controller';
import {godForm} from 'ember-easy-orm/mixins/form';

export default Controller.extend(godForm, {
    modelName: "app",
    actions: {
        cancel(){
            this.transitionToRoute('app');
        },
        success(){
            this.transitionToRoute('app');
        }
    }
});
