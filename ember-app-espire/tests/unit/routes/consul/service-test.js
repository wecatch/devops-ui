import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | consul/service', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:consul/service');
    assert.ok(route);
  });
});
