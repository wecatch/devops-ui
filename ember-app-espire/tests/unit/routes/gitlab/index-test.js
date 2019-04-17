import { module, test } from 'qunit';
import { setupTest } from 'ember-qunit';

module('Unit | Route | gitlab/index', function(hooks) {
  setupTest(hooks);

  test('it exists', function(assert) {
    let route = this.owner.lookup('route:gitlab/index');
    assert.ok(route);
  });
});
