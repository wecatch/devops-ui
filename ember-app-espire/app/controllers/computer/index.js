import Controller from '@ember/controller';
import {paginationController} from 'ember-app-espire/mixins/paging';

export default Controller.extend(paginationController, {
    queryParams: ['app_name'],
    app_name: null,
    actions: {
        searchApp(){
            let page = this.page;
            let appname = this.app_name;
            this.transitionToRoute({
                queryParams: {
                    page: page,
                    app_name: appname
                }
            });
        }
    }
});
