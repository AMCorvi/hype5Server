const Hype5 = require("hype5");

/* eslint-disable no-unused-vars */
class Service {
  constructor (options) {
    this.options = options || {};
  }

  find (params) {
    //Determin if ID is equal top || remixes || noremixes

    let data = (async function () {
      let output
      if (params.type) {
        let H = await Hype5.top(params.type)
          .then( data => output = data )
          .catch( err => console.error( err ) );
      } else {
        let H = await Hype5.top()
          .then( data => output = data )
          .catch( err => console.error( err ) );

      }
      return output;
    })();

    return Promise.all([data]);
  }

  get (id, params) {

    //Determin if ID is equal top || remixes || noremixes
    if ( !id === ( "top"||"remixes"||"noremixes" ) ) {
      return Promise.resolve({
        err: {
          message: `"${id}" is invalid value for the Hype-5-Server endpoint!. Please use one of the three following values: "top","remixes","noremixes"`
        }
      });
    }

    let data = (async function () {
      let output
      if (params.type) {
        let H = await Hype5[id](params.type)
          .then( data => output = data )
          .catch( err => console.error( err ) );
      } else {
        let H = await Hype5[id]()
          .then( data => output = data )
          .catch( err => console.error( err ) );

      }
      return output;
    })();

    return Promise.all([data]);

  }

  create (data, params) {
    if (Array.isArray(data)) {
      return Promise.all(data.map(current => this.create(current)));
    }

    return Promise.resolve(data);
  }

  update (id, data, params) {
    return Promise.resolve(data);
  }

  patch (id, data, params) {
    return Promise.resolve(data);
  }

  remove (id, params) {
    return Promise.resolve({ id });
  }
}

module.exports = function (options) {
  return new Service(options);
};

module.exports.Service = Service;
