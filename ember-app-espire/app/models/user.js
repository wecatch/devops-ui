import EmberObject from '@ember/object';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;

export default EmberObject.extend(model,{
    url: '/v1/user',
    init() {
        this._super(...arguments);
        this.model = {
            'email': attr('string'),
            'nickname': attr('string'),
            'province_id': attr('string'),
            'city_id': attr('string'),
            'area_id': attr('string'),
            'town_id': attr('string'),
            'country_id': attr('string'),
        }
    }
});