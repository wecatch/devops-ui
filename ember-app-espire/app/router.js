import EmberRouter from '@ember/routing/router';
import config from './config/environment';

const Router = EmberRouter.extend({
  location: config.locationType,
  rootURL: config.rootURL,
  didTransition: function() {
    this._super(...arguments);
    let routeName = this.currentURL;
    if (!localStorage['routeList']){
      localStorage['routeList'] = routeName;
      this.set('previousRouteName', routeName);
    }else {
      localStorage['routeList'] = routeName + "," +localStorage['routeList'];
      let routeList = localStorage['routeList'].split(",");
      routeList = routeList.splice(0, 2);
      this.set('previousRouteName', routeList[0]);
      localStorage['routeList'] = routeList.join(",")
    }
  }
});

Router.map(function() {
  this.route('domain', function() {});
  this.route('computer', function() {
    this.route('only');
  });
  this.route('app', function() {
    this.route('detail',{path: '/:app_id/detail'}, function() {
      this.route('tags', {path: '/:repository_id/tags'});
      this.route('commits', {path: '/:repository_id/commits'});
      this.route('deploys');
      this.route('computers');
      this.route('deploy-detail', {path: '/deploy/:deploy_id'});
    });
    this.route('new');
    this.route('edit', {path:'/:app_id'});
  });
  this.route('gitlab', function() {
    this.route('group');
    this.route('project');
  });
  this.route('tag', function() {
    this.route('new');
  });

  this.route('detail', function() {});
  this.route('consul', function() {
    this.route('service');
    this.route('kv');
  });
  this.route('cloud', function() {
    this.route('region');
  });
});

export default Router;
