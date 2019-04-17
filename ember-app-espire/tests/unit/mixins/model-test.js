import EmberObject from '@ember/object';
import ModelMixin from 'ember-app-espire/mixins/model';
import { module, test } from 'qunit';

module('Unit | Mixin | model', function() {
  // Replace this with your real tests.
  test('it works', function (assert) {
    let ModelObject = EmberObject.extend(ModelMixin);
    let subject = ModelObject.create();
    assert.ok(subject);
  });
});
