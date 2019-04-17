import EmberObject from '@ember/object';
// import A from '@ember/array';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;


export default EmberObject.extend(model, {
    url: "/v1/app/deploy",
    init(){
        this._super(...arguments);
        this.model = {
            id: attr('number'),
            app_id: attr('number'),
            max:  attr({defaultValue: function(){
                return 1;
            }}),
            commit_id:  attr('string'),
            rollback_id:  attr('string'),
            desc: attr('string'),
            commit_tag: attr('string'),
            rollback_tag: attr('string'),
            hosts: attr('array'),
            status: attr('string'),
            binary_url: attr('binary_url'),
            interval:  attr({defaultValue: function(){
                return 60;
            }}),
        }
    }
});
