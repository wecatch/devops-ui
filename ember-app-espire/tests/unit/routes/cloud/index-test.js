import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | cloud/index', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:cloud/index');
    assert.ok(route);
  });
});
