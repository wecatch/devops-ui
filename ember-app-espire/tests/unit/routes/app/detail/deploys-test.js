import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | app/detail/deploys', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:app/detail/deploys');
    assert.ok(route);
  });
});
