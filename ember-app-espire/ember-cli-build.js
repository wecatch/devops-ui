'use strict';

const EmberApp = require('ember-cli/lib/broccoli/ember-app');

module.exports = function(defaults) {

    var options = {
      minifyJS: {
        enabled: true
      },
      fingerprint: {
        enabled: false
      },
      minifyCSS: {
        enabled: true
      },
      sourcemaps: {
        enabled: false
      },
      storeConfigInMeta: false,
      'ember-composable-helpers': {
        only: ['slice'],
      },
      hinting: false
  };

  if (process.env.EMBER_ENV === 'development') {
      options.minifyJS = {
        enabled: false
      };
      options.sourcemaps = {
        enabled: true
      };
  }

  let app = new EmberApp(defaults, options);
  app.import('node_modules/ansi_up/ansi_up.js');
  app.import('vendor/shims/ansi-up.js');

  // Use `app.import` to add additional libraries to the generated
  // output files.
  //
  // If you need to use different assets in different
  // environments, specify an object as the first parameter. That
  // object's keys should be the environment name and the values
  // should be the asset to use in that environment.
  //
  // If the library that you are including contains AMD or ES6
  // modules that you would like to import into your application
  // please specify an object with the list of modules as keys
  // along with the exports of each module as its value.

  return app.toTree();
};
