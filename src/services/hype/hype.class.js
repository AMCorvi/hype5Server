const Hype5 = require("hype5");

/* eslint-disable no-unused-vars */
class Service {
  constructor (options) {
    this.options = options || {};
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
      let H = await Hype5[id]()
      .then( data => output = data )
      .catch( err => console.error( err ) );
      return output;
    })();

    return Promise.all([ id, data ]);

  }


}

module.exports = function (options) {
  return new Service(options);
};

module.exports.Service = Service;
