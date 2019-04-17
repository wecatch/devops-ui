export function initialize(application) {
    // application.inject('route', 'foo', 'service:foo');
    application.inject('route', 'session', 'service:session');
    application.inject('controller', 'session', 'service:session');
    application.inject('component', 'session', 'service:session');
    application.inject('router', 'session', 'service:session');

    application.inject('route', 'flashMessages', 'service:flashMessages');
    application.inject('controller', 'flashMessages', 'service:flashMessages');
    application.inject('component', 'flashMessages', 'service:flashMessages');
    application.inject('router', 'flashMessages', 'service:flashMessages');

    application.inject('route', 'flashLoading', 'service:flashLoading');
    application.inject('controller', 'flashLoading', 'service:flashLoading');
    application.inject('component', 'flashLoading', 'service:flashLoading');
    application.inject('router', 'flashLoading', 'service:flashLoading');

    application.inject('component', 'store', 'service:store');
}

export default {
    name: 'session',
    initialize
};
