"use strict";

var _server = require("./server.js");

var _server2 = _interopRequireDefault(_server);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

describe("Jest is working", function () {
    test("That a test works", function () {
        expect(new Number(300)).toBeInstanceOf(Number);
    });
    test("Jest parse flow", function () {
        expect((0, _server2.default)("AM")).toEqual("AMCorvi");
    });
});
//# sourceMappingURL=server.spec.js.map
