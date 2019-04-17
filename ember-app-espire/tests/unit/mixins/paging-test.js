import EmberObject from '@ember/object';
import PagingMixin from 'ember-app-espire/mixins/paging';
import { module, test } from 'qunit';

module('Unit | Mixin | paging', function() {
  // Replace this with your real tests.
  test('it works', function (assert) {
    let PagingObject = EmberObject.extend(PagingMixin);
    let subject = PagingObject.create();
    assert.ok(subject);
  });
});
