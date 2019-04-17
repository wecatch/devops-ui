// import {set} from '@ember/object';
import Route from '@ember/routing/route';
import { hash } from 'rsvp';

export default Route.extend({
    model(params){
        return hash({
            app: this.store.findOne('app', params.app_id),
        });
    },
    // afterModel(model){
    //     return this.store.find('computer', {appid: model.app.app_id, tag: model.app.tag, name: model.app.name}).then((data)=>{
    //         set(model, 'computer', data);
    //     });
    // }
});
