"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const surrealdb_js_1 = require("surrealdb.js");
const db = new surrealdb_js_1.Surreal("http://127.0.0.1:8000/rpc");
async function main() {
    try {
        await db.signin({
            user: "root",
            pass: "root"
        });
        await db.use('test', 'test');
        await db.delete("food");
        await db.create("food:pizza", {
            toppings: ["cheese"],
            rating: 10,
        });
        await db.create("food:pasta", {
            toppings: ["basil"],
            rating: 9,
        });
        let pizza = await db.query("SELECT * FROM food WHERE rating > 9;");
        console.log(JSON.stringify(pizza));
        let pasta = await db.query("SELECT * FROM food WHERE toppings CONTAINS 'basil'");
        console.log(JSON.stringify(pasta));
        let toppings = pizza[0].result[0].toppings;
        toppings.push("pineapple");
        let updated = await db.change("food:pizza", {
            toppings,
        });
        console.log(updated);
    }
    catch (e) {
        console.error(e);
    }
}
main();
//# sourceMappingURL=main.js.map