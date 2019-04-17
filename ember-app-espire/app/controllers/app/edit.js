import Controller from '@ember/controller';
import {godForm} from 'ember-easy-orm/mixins/form';

export default Controller.extend(godForm, {
    modelName: "app",
    actions: {
        cancel(){
            if(this.target.previousRouteName){
                this.transitionToRoute(this.target.previousRouteName);
            }else {
                this.transitionToRoute('app');
            }
        },
        success(){
            if(this.target.previousRouteName){
                this.transitionToRoute(this.target.previousRouteName);
            }else {
                this.transitionToRoute('app');
            }
        }
    }
});
