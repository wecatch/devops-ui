import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | computer/only', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:computer/only');
    assert.ok(route);
  });
});
