import Mixin from '@ember/object/mixin';

var paginationController = Mixin.create({
    queryParams: ['page'],
    limit: 50,
    page: 1,
    actions: {
        prevPage: function() {
            var page = this.page - 1;
            console.log(this.page);
            if (page < 1) {
                page = 1
            }
            this.set('page', page);
            this.transitionToRoute({
                queryParams: {
                    page: page,
                }
            });
        },
        nextPage: function() {
            var page = this.page + 1;
            // if have next page
            // if (this.get('model').length % this.limit !== 0) {
            //     return;
            // }

            this.transitionToRoute({
                queryParams: {
                    page: page
                }
            });
        }
    }
});


var paginationRoute = Mixin.create({
    queryParams: {
        page: {
            refreshModel: true
        }
    }
});


export { paginationRoute, paginationController};
