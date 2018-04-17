// Initializes the `hype` service on path `/hype`
const createService = require('./hype.class.js');
const hooks = require('./hype.hooks');
const filters = require('./hype.filters');

module.exports = function () {
  const app = this;
  const paginate = app.get('paginate');

  const options = {
    name: 'hype',
    paginate
  };

  // Initialize our service with any options it requires
  app.use('/hype', createService(options));

  // Get our initialized service so that we can register hooks and filters
  const service = app.service('hype');

  service.hooks(hooks);

  if (service.filter) {
    service.filter(filters);
  }
};
