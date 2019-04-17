import Route from '@ember/routing/route';

export default Route.extend({
    queryParams: {
        name: {
            refreshModel: true
        },
        tag: {
            refreshModel: true
        }
    },
    name: "",
    tag: "",
    model(params){
        return this.store.request.get('/v1/consul/service/list', {data: params});
    }
});
