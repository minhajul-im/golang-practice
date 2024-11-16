---
date: 2024-10-12
title: Understanding "this" A Comprehensive Guide 1st-part!
description: এই post এ আমরা "this" keyword নিয়ে কথা বলবো; Global Context, Function Context এবং Object method কেমন আচরণ করে deep drive করবো। এইটা 1st part 2nd part আসবে।
author: Minhajul Islam (Minhaj)
published: true
---

# "this" keyword in JavaScript 1st Part!

_"this"_ keyword এটা তো panic একটা word JavaScript এর মধ্যে। যাই হোক; this keyword কে নিয়ে একটু ভেঙে চুরে দেখার চেষ্টা করবো; কোথায় কেমন আচরণ করে চলেন! আমি দুই part করেছি একটা post করলে অনেক বড় হয়ে যায়।

_1st part_

**Note**: তবে একটা কথা বলে রাখি, কেউ যদি আপনাকে বলে `this` এর output কি হবে? মনে রাখবেন _**কিভাবে এবং কোথায়**_ call হচ্ছে `this` keyword.
_মনে রাখবেন আমি এই কথাটা বার বার বলবো কিন্তু_

"this" keyword অনেক ভাবে refer করে context; সেগুলা চলেন দেখি।

## **Global / Window context**:

যদি কোন scope বা function এর মধ্যে this কে `log` না করেন; just file এর মধ্যে log করেন তাহলে দেখবেন যে browser এ window object এবং Node.js এ global object কে refer করে।

```javascript
console.log(this); // window or global;
```

আবার [Runtime Environment](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/Runtime-Environments.md) এর নিজস্ব Object থাকে; যেমন টাই বলছিলাম browser এ window এবং Node.js এ global;

## **Function context**:

function context এ this এখানে দুই রকম আচরণ করে;

_চলেন দেখি কোড করে দেখি_

```javascript
console.log(this); // window or global
function regularFunc() {
  console.log(this); // window or global
}
regularFunc();
```

এইখানে দেখেন global এবং function দুইটাই কিন্তু একই output দেখাচ্ছে। আপনি যদি regularFunc এর মধ্যে this global / window object কে refer করে।

_চলেন আবার একটা function এর example দেখি_

```javascript
"use strict";
console.log(this); // window or global
function regularFunc() {
  console.log(this); // undefined
}
regularFunc();
```

এইবার result কিন্তু একে রকম না! এইবার regularFunc এর মধ্যে undefined দেখাচ্ছে। এর কারণ হল যে strict mode এবং non-strict mode; _strict mode & non strict এর মধ্যে পার্থক্য কি?_ non-strict mode হল যে আপনি চাইলে কিন্তু `x = "non-strict-mode"` আপনি var, let, const না দিয়েও x নামে একটা variable declare করতে পারবেন; কোন error দিবে না, তখন globally declare হয়ে যাবে। আপনি কিন্তু strict mode এ `x = "strict-mode"` এইভাবে declare করতে পারবেন না, আপনাকে syntax error দিয়ে দিবে। মূলত ভালো programmer রা _strict mode use_ করে। আপনি অবশ্যই ভালো programmer। আপনি strict and non strict mode নিয়ে আরও deep-dive করতে পারেন। non-strict mode এ this substitution করে window or global কে। strict mode এ this substitution করতে পারে না। দেখেন, আমি আগেই note এ বলেছি যে this কে _*কোথায় এবং কিভাবে*_ ?

**এবার আসেন Arrow Function দেখি**

```javascript
const arrowFunc = () => {
  console.log(this); // window or global
};
arrowFunc();
```

মনে রাখবেন arrow function এর নিঃস্ব কোন **`this`** নেই; তাহলে window or global কেন output দিল? Great question! arrow function surrounding মানে তার parent হল global or window এর সাথে lexical binding করে `this` এর output দেয়। তাই output আসে window or global; এইখানে strict এবং non-strict mode কোন effect করে না। তবে arrow function আরো একটা example object এর সাথে সেটা নিচে দেখবো।

## **Object method**

_object এর মধ্যে this কি আচরণ করে চলেন code করে দেখি_

```javascript
const user = {
  fName: "Mr. John",
  lName: "Khan",
  fullName: function () {
    console.log(this);
    // {fName: 'Mr. John', lName: 'Khan', fullName: ƒ}
    console.log(this.fName + " ", this.lName);
    // Mr. John Khan
  },
};
user.fullName();
```

এইখানে দেখেন কি হয়েছে; user নামে একটা object নিয়েছি এবং fullName নামে একটা regular function নিয়েছি _(আমি function বলতেছি বুঝার সুবিধাতে আসলে এইটা একটা Method)_ এর মধ্যে যখন this কে log করেছি দেখেন কি output দিয়েছে; `{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}` এখানে কিন্তু কোন window বা global কে substituted করে নি বরং নিজের object কে refer করেছে এবং নিচের log টা দেখেন this দিয়ে কিন্তু আমরা user এর মধ্যে property গুলা access করতে পাচ্ছি তাই তো output `Mr. John Khan` পেয়েছি। তাহলে আমরা regular function এর মত কে declare করলে this এর মধ্যে own object কে refer করে; এটা strict mode এবং non strict একই behavior;

_আবার আসেন code টা একটু change করে দেখি_

```javascript
const user = {
  fName: "Mr. John",
  lName: "Khan",
  fullName: () => {
    console.log(this); // Window or Global
    console.log(this.fName + " ", this.lName); // undefined
  },
};
user.fullName();
```

তেমন কিছুই change করি নাই just regular function থেকে arrow function করেছি; regular function এর কি output পেয়েছিলাম `{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}` ; arrow function এর _Window or Global_ এটা আসলো কেন? আমরা আগেই দেখেছিলাম যে Arrow Function এর নিজের কোন this নাই। ঐ যে surrounding এর সাথে lexical binding করে Global or Window কে refer করতেছে; `console.log(this.fName + " ", this.lName);` এই `log` টাতে _undefined_ আসলো কেনো? কারণ হল যে global or window তে তো আর `this.fName + " ", this.lName` variable declare করি নাই; তাই access করতে পারব না; তাই undefined দিয়েছে। তাহলে JavaScript কোন কিছু declare না করলে সেটা access করলে তো undefined আসেই।

_এবার আসেন দেখি code টা আরেকটু change করি_

```javascript
const user = {
  fName: "Mr. John",
  lName: "Khan",
  fullName: function () {
    console.log("regular", this);
    //regular:{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}
    const arrowFunc = () => {
      console.log("arrow", this);
      // arrow : {fName: 'Mr. John', lName: 'Khan', fullName: ƒ}
      console.log(this.fName + " ", this.lName);
      // Mr. John Khan
    };
    arrowFunc();
  },
};
user.fullName();
```

এবার দেখেন আমরা just arrowFunc টা regular function এর ভিতরে দিয়েছি; fullName একাটা regular function; তখন কি হল? আমি আগেই বলেছি যে arrow function তার surrounding এর সাথে lexical binding করে `log` গুলা করে result দিয়েছে। arrowFunc এর surrounding কে fullName যেটার output এইটা `regular:{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}`; arrowFunc এর this কি দাঁড়ালো surrounding এর সাথে lexical binding করে output দিয়েছে এইটা `arrow :{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}` ; দুইটা result একই; তাহলে কি হল fullName function তার own object কে refer করে। যার জন্য this এর value `{fName: 'Mr. John', lName: 'Khan', fullName: ƒ}` এবং `this.fName + " ", this.lName // Mr. John Khan` দিয়েছে।

**এইবার একেবারে compare করবো চলেন** : কেন দুই রকম behavior করে? question থাকতে পারে না। তাহলে কেন regular function global or window কে refer করে না; arrow function তার surrounding এর সাথে lexical binding করে refer করে output দেয়!

```javascript
const user = {
  fName: "Mr. John",
  lName: "Khan",
  fullNameRegular: function () {
    console.log(this); // Refers to user
    console.log(this.fName + " ", this.lName);
    // Logs: Mr. John Khan
  },
  fullNameArrow: () => {
    console.log(this); // Refers to the global object
    console.log(this.fName + " ", this.lName);
    // Logs: undefined undefined
  },
};
user.fullNameRegular(); // Calls the regular function
user.fullNameArrow(); // Calls the arrow function
```

এর answer আমিও অনেক খোঁজার চেষ্টা করেছি; যেটা পেয়েছি সেটা হল যে JavaScript এর **rule**; আপনি আরও খুঁজে দেখতে পারেন।

## **DOM**

HTML-এ আপনি `this` দিয়ে element এর তার নিজের সব কিছু access করতে পারবেন;

_চলেন দেখি code করে_

```html
<button onclick="alert(this)">get this object</button>
<button onclick="alert(this.innerText)">get inner text</button>
<button onclick="this.setAttribute('disabled', true)">disabled</button>
<button id="button-id" class="btn" onclick="alert(this.id)">get id</button>
<button
  onclick="this.style.backgroundColor = 'green'; this.style.color = 'red'; this.innerText = 'Clicked!';">
  set color
</button>
```

দেখেন আমি কিছু code করে দিয়েছি আপনিও code করে দেখুন; সত্যি বলতে কখনো লেগেছে এমন টা আমার জানা নাই; একটু জেনে রাখছি আরকি।

### **Conclusion**

1st part এ আমরা "this" keyword-এর বিভিন্ন আচরণ বিশ্লেষণ করেছি। আমরা দেখেছি কিভাবে "this" global context, function context, object methods and DOM কেমন আচরণ করে।
global and window context এ "this" সাধারণত global and window object নির্দেশ করে, তবে strict mode এটি "undefined" হয় regular function এর সময়। এছাড়া, arrow function lexical binding করে।
object methods, "this" object এর নিজস্ব property access করতে সাহায্য করে। dom নিয়েও একটু দেখেছি কিভাবে this কাজ করে।

পরবর্তী 2nd part আমরা "this" এর আরও deep drive করব; যেখানে "call," "apply," এবং "bind" methods গুলার ব্যবহার এবং class মধ্যে "this" এর আচরণ নিয়ে আলোচনা করা হবে।

লিখতে গিয়ে আমার ভুল হতে পারে; দয়া করে comments করে জানাবেন, আমি ঠিক করে নেব।
