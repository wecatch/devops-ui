import Controller from '@ember/controller';
import {paginationController} from 'ember-app-espire/mixins/paging';

export default Controller.extend(paginationController, {
    queryParams: ['name'],
    name: null,
    actions: {
        searchApp(){
            let page = this.page;
            let appname = this.name;
            this.transitionToRoute({
                queryParams: {
                    page: page,
                    name: appname
                }
            });
        }
    }
});
