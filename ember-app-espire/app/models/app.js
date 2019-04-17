import EmberObject from '@ember/object';
// import A from '@ember/array';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;


export default EmberObject.extend(model, {
    url: "/v1/app",
    primaryKey: 'id',
    urlForFind: function() {
        return this.get('api') + '/list';
    },
    urlForSave: function(id){
        if(id){
            return this.get('api') + '/'+ id;
        }else {
            return this.get('api') + '/create';
        }
    },
    init(){
        this._super(...arguments);
        this.model = {
            name: attr('string'),
            url: attr('string'),
            internal_url: attr('string'),
            desc: attr('string'),
            tag: attr('string'),
            repository_url: attr('string'),
            deploy_dir: attr('string'),
            monitor_url: attr('string'),
            repository_id: attr('number'),
            port: attr('number'),
            update_code_cmd: attr('string'),
            reload_service_cmd: attr('string'),
            check_service_cmd: attr('string'),
            cmd_name: attr('string'),
            cmd_dir: attr('string'),
        }
    }
});