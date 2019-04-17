import EmberObject from '@ember/object';
import A from '@ember/array';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;


export default EmberObject.extend(model, {
    url: "/v1/domain/list",
    init(){
        this._super(...arguments);
        this.model = {
            name: attr('string'),
            ip:  attr({defaultValue: function(){
                return A();
            }}),
            host: attr('string'),
            created_at: attr('string'),
            updated_at: attr('string')
        }
    }
});
