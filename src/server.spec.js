import serve from "./server.js"

describe("Jest is working", () => {
    test("That a test works", () => {
        expect(new Number(300)).toBeInstanceOf(Number);
    });
    test("Jest parse flow",() =>{
        expect(serve("AM")).toEqual("AMCorvi");
    });
})
