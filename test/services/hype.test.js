const assert = require('assert');
const app = require('../../src/app');

describe('\'hype\' service', () => {
  it('registered the service', () => {
    const service = app.service('hype');

    assert.ok(service, 'Registered the service');
  });
});
