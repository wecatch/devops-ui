import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | app/detail/deploy-detail', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:app/detail/deploy-detail');
    assert.ok(route);
  });
});
