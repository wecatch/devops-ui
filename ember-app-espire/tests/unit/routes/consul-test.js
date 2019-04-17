import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | consul', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:consul');
    assert.ok(route);
  });
});
