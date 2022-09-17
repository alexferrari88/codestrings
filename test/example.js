function greet(name) {
  console.log(`Hello ${name}`);
}

var fname = "John";
const lname = "Doe";
let age = 42;
let address = {
  street: "123 Main St",
  city: "New York",
  state: "NY",
};
const today = new Date();
const person = {
  fname,
  lname,
  age,
  address,
  today,
  getBirthYear: function () {
    return today.getFullYear() - age;
  },
};
console.log(person);
if (person.age >= 21) {
  console.log("You can drink");
}
if (person.fname === "John") {
  console.log("You are John");
}
const a = "Hello";
const b = "World";
const c = "!";
const d = a + " " + b + c;
console.log(d);
const e = `${a} ${b}${c}`;
console.log(e);

const empty = "";
const empty2 = ``;
const empty3 = ``;
const space = " ";
const space2 = ` `;
const space3 = ` `;

const withInsideDoubleQuotes = 'This is a "quote"';
const withInsideSingleQuotes = "This is a 'quote'";
const withInsideBackticks = `This is a 'quote'`;
