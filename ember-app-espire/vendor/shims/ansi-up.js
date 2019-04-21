(function() {
  function vendorModule() {
    'use strict';

    return {
      'default': self['ansi-up'],
      __esModule: true,
    };
  }

  define('ansi-up', [], vendorModule);
})();
