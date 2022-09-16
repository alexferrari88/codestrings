function greet(name) {
  console.log(`Hello ${name}`);
}

var fname = "Alex";
const lname = "Ferrari";
let age = 34;
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
if (person.fname === "Alex") {
  console.log("You are Alex");
}
