import EmberObject from '@ember/object';
import model, {
    DS
} from '../mixins/model';

const {
    attr
} = DS;

export default EmberObject.extend(model,{
    url: '/v1/app/tag',
    init() {
        this._super(...arguments);
        this.model = {
            'name': attr('string'),
            'kind': attr('string'),
        }
    }
});