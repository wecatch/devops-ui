import Service from '@ember/service';
import {set} from '@ember/object';
import {run} from '@ember/runloop';
import {guidFor} from '@ember/object/internals';



export default Service.extend({
    loading: false,
    elementId: null,
    lastElementId: null,
    trigger(obj, status) {
        let elementId = guidFor(obj);
        //start loading
        if (this.elementId === null) {
            set(this, 'elementId', elementId);
            set(this, 'loading', true);
            run.later(this, function() {
                if (this.lastElementId !== elementId) {
                    //again wait for start loading
                    set(this, 'elementId', null);
                    set(this, 'loading', false);
                }
            }, 10000);
            return;
        }

        if (status === false) {
            //stop loading, only the same element can cacel loading
            if (this.elementId === elementId) {
                set(this, 'elementId', null);
                set(this, 'loading', false);
                set(this, 'lastElementId', elementId);
            }
        }
    }
});