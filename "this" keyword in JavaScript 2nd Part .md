---
date: 2024-10-12
title: Understanding "this" A Comprehensive Guide 2nd-part!
description: এই post এ আমরা "this" keyword নিয়ে কথা বলবো; Call, Apply & Bind এবং Class কেমন আচরণ করে deep drive করবো। এইটা 2nd part 1st part এর link দিয়ে দিব।
author: Minhajul Islam (Minhaj)
published: true
---

# "this" keyword in JavaScript 2nd Part!

_"this"_ keyword নিয়ে আমরা [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) Global Context, Function Context, Object Methods এবং DOM নিয়ে দেখেছিলাম এই part এ আমরা (Call, Apply & Bind) & Class নিয়ে দেখব। এইটা 2nd part; [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) টা highly request থাকবে [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) দেখে then 2nd part দেখবেন তাহলে ভালো হবে।

_2nd part_

## **Call, Apply & Bind**:

_this কিভাবে কাজ করে সেটাই দেখব চলেন code করি_

```javascript
const user = { name: "Mr. John Khan" };
function funcThis() {
  console.log(this);
  console.log(`Hello ${this.name}`);
}
funcThis();
```

এখানে একটা object নিয়েছি যার মধ্যে property আছে name। আপনি চাইলে আরো রাখতে পারেন আপনার requirement অনুযায়ী। আরেকটা funcThis নামে function আছে যার মধ্যে দুইটা log আছে; এখন আমরা যদি funcThis() invoke/call করি তাহলে কি output আসবে?

_চলেন দেখি_

```javascript
const user = { name: "Mr. John Khan" };
function funcThis() {
  console.log(this); // Window or Global
  console.log(`Hello ${this.name}`); // Hello
}
funcThis();
```

যদি strict mode থাকে তাহলে ১ম log `this` undefined, ২য় log syntax error দিবে; কারণ হল যে `this.name` নামে কিছুই নাই Window or Global এর মধ্যে। এইটুকু তো জানার কথা আমরা [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) তো বিস্তর কথা বলছি এই নিয়ে। এইবার আসেন non-strict mode যদি হয় তাহলে output তো দিয়েই দিয়েছি comments করে; এটা আপনার পারার কথা ২য় log এ কেন Hello এরপরে কিছু আসে নাই? কারণ হল যে `this` তো global or window কে refer করে তাই global or window তে name নামে কোন property নাই তাই Hello প্রিন্ট হয়েছে। আমি নতুন কিছুই বলি নাই এখনো সবগুলা [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) এর কথা; আপনি যাতে ভুলে না যান।

_চলেন code টা কে একটু modify করি_

```javascript
const user = { name: "Mr. John Khan" };
function funcThis() {
  console.log(this); // {name: 'Mr. John Khan'}
  console.log(`Hello ${this.name}`); // Hello Mr. John Khan
}
funcThis.call(user);
funcThis.apply(user);
funcThis.bind(user);
```

সব কিছু একই আছে তবে funcThis function টা কে direct call/invoke না করে Function Methods (_methods কি? function নাম এর পর dot দিলে অনেক আসবে ওইগুলা methods_) দিয়ে call/invoke করেছি সাথে user object পাঠাচ্ছি। `call`, `apply`, `bind` এই function methods গুলা কি করতেছে? `call`, `apply`, `bind` methods গুলা funcThis এর যে আগে `this` যেটা Window or Global refer করতো সেটাকে replace করে দিয়ে user নামে যেই object টা pass করেছি; সেই user object এর `this` কে refer করতেছে; তাই output দেখেন change হয়ে গেছে `{name: 'Mr. John Khan'}; Hello Mr. John Khan`

strict or non-strict mode কোন effect ফেলবে না, যদি আপনি with context দিয়ে call/invoke করেন; তবে মনে রাখবেন আপনি funcThis function কে without context call/invoke করতে পারবেন না তাহলে error দিবে যেটা উপরে বলা হয়েছে `strict and non strict mode`

_আপনি আরেকবার পড়েন: উপরের পার্টে কিভাবে `this` এর reference পরিবর্তিত হল_

**note :** Call, Apply & Bind নিয়ে আমি details বলতে চাচ্ছি না; এটা আলাদা একটা part.

**Call**: আপনি function এর মধ্যে parameter নিতে পারেন সেগুলা argument আকারে pass করতে পারবেন।

_চলেন code করে দেখি_

```javascript
function funcThis(age, address) {}
funcThis.call(user, age, address);
```

_Note : মনে রাখবেন_ যখন funcThis কে call/apply/bind methods দিয়ে call/invoke করবেন; সবার আগে this মানে নতুন একটা this এর references দিতে হবে; তারপর আপনি যত গুলা parameter নিয়েছিলেন ততগুলা argument pass করবেন।

**Apply**: এটা একদম Call এর মতই, তবে আপনাকে arguments গুলা array তে pass করতে হবে।

_চলেন code করে দেখি_

```javascript
function funcThis(age, address) {}
funcThis.call(user, [age, address]);
```

এইভাবে pass করতে হবে।

**Bind**: এটা একদম Call এর মতই, আপনি Call এর মত করে argument pass করতে পারবেন; তবে bind method একটা function তৈরি করে নেয়।

_চলেন code করে দেখি_

```javascript
const user = { name: "Mr. John Khan" };
function funcThis() {
  console.log(this); // {name: 'Mr. John Khan'}
  console.log(`Hello ${this.name}`); // Hello Mr. John Khan
}
const bindFunc = funcThis.bind(user);
bindFunc(22);
```

call and apply কিন্তু `bindFunc` এইখানে result রাখত bind কিন্তু function make করে রাখছে। `bindFunc` function always মনে রাখবে object এর `this`; যখন দরকার পড়বে `bindFunc` call/invoke করার, তখন arguments pass করলেই result পাবেন; funcThis এর default `this` Window or Global থাকবে না। সব সময়ের জন্য user object এর `this` references করে থাকবে। `bindFunc` কখন ভুলে যাবে না পাঠানো this এর কথা। এই জন্য Call & Apply থেকে Bind আলাদা।

**Note**: Call & Apply এই দুইটা methods যখন call/invoke হবে তখন immediately with a specified `this` দিয়ে শেষ হয়ে যাবে; মনে করেন আপনি উপরের funcThis কে user object pass করে call/invoke করলে immediately result দিয়ে user এর `this` শেষ হয়ে যাবে funcThis function এর মধ্যে থেকে। funcThis function user এর `this` এর কথা মনে রাখবে না ঐটা completely শেষ। এই জায়গায় `Bind` method আলাদা, সে মনে রাখতে পারে যা তাকে দেওয়া হয়েছিল user object এর this সে কখনো ভুলে যাবে না; call and apply তা পারে না ভুলে যায় এবং function make করে রাখে। আমি just concept টুকু বলার চেষ্টা করলাম। আপনার ইচ্ছা হলে আপনি more details দেখতে পারেন;

## **Class**:

_Class_ এর মধ্যে `this` এর আচরণ কেমন? এখানে বুঝতে তেমন অসুবিধা হওয়ার কথা না, আমরা already এই বিষয় গুলা নিয়ে আলোচনা করেছি [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md)। (আমি class এর ভিতরে function নাম বলছি, আসলে ওইগুলা methods বুঝার সুবিধাতে function বলতেছি)

_চলেন code করে দেখি_

```javascript
class ThisClass {
  constructor(name) {
    this.name = name;
  }
  arrowFunc = () => {
    console.log("arrow", this);
    // arrow: ThisClass {name: 'Mr. John Khan', arrowFunc: ƒ}
  };
  regularFunc() {
    console.log("regular", this);
    // regular: ThisClass {name: 'Mr. John Khan', arrowFunc: ƒ}
  }
}
const instance = new ThisClass("Mr. John Khan");
instance.arrowFunc();
instance.regularFunc();
```

এইখানে কি output আসবে সেটা তো এতক্ষণ জানা কথা, তাই না? দুইটা function declare করেছি arrowFunc এবং regularFunc নামে; দুইটার মধ্যেই `this` কে log করেছি; **"class কিন্তু একটা Object"** regularFunc কি করবে object কে output দিয়ে দিবে যেটা আমরা আগে দেখেছি object method. arrowFunc এর নিজস্ব কোন `this` নাই, তাই সে surrounding এর সাথে lexical binding করে output দিবে; surrounding কে আছে class object; তাই দুইটার output একই আসবে। কিন্তু output কেমন দেখাচ্ছে, দেখছেন তো? ভালো করে দেখেন কিন্তু `ThisClass {name: 'Mr. John Khan', arrowFunc: ƒ}` দুই জায়গায় একই output, তবে arrowFunc কিন্তু দুই জায়গায় আছে, regularFunc কিন্তু নাই।

_কেন, question হতে পারে না? এইখানে কি হয়েছে, চলেন দেখি।_

```javascript
const instance = new ThisClass("Mr. John Khan");
```

new দিয়ে instance create করে নিয়েছি। [new কি?](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/new) আপনি দেখে নিয়েন, আমি আর deep এ যাচ্ছি না; new creates a new instance of a class (or constructor function)। আপনি যখন instance কে `console.log(instance);` শুধু এইটাই না আমরা যে আগে regularFunc এবং arrowFunc funtion কে call/invoke করেছিলাম তখন সেগুলার log এর মধ্যে `[[Prototype]]:Object` click করলে দেখতে পারতাম।

_এমন output দেখতে পাবেন_

```javascript
ThisClass {name: 'Mr. John Khan', arrowFunc: ƒ}
arrowFunc:() => { console.log("arrow", this); }
name:"Mr. John Khan"
[[Prototype]]:Object // এইটাতে আপনি click করলেই দেখতে পারবেন।
constructor:class ThisClass
regularFunc:ƒ regularFunc()
[[Prototype]]:Object
```

অবশেষে regularFunc দেখতে পাইলাম; তবে regularFunc সাথে constructor আছে।
তার মানে হল যে এই যে regularFunc function টা ThisClass class এর Prototype হয়ে গেছে automatically; [_Prototype কি?_](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Object_prototypes) সহজ বাংলায় method reuse করা যায়; যেটা কে Inheritance বলে। regularFunc টা **specific একটা point এ memory তে address** করেছে। এখন আপনি যদি অনেক গুলা ThisClass এর `const instance = new ThisClass("Mr. John Khan");` এইরকম instance create করেন, তাহলে regularFunc বার বার create হবে না। regularFunc just reference হয়ে কাজ করবে। এই জন্যই Prototype হয়।
arrowFunc যদি অনেক গুলা এইরকম `const instance = new ThisClass("Mr. John Khan");` instance create করেন; তাহলে দেখবেন প্রতিটি instance এর জন্য একটা করে function create করবে এবং নতুন নতুন point এ memory address করবে; arrowFunc স্বাধীন ভাবে কাজ করবে।

_চলেন একটু compare করে দেখি_

```javascript
class ThisClass {
  constructor(name) {
    this.name = name;
  }
  arrowFunc = () => {
    console.log("arrow", this);
  };
  regularFunc() {
    console.log("regular", this);
  }
}
const instance1 = new ThisClass("Mr. John Khan");
const instance2 = new ThisClass("Mr. John Khan");

// arrow function compare
if (instance1.arrowFunc === instance2.arrowFunc) {
  console.log("arrow function", true);
} else {
  console.log("arrow function", false); // false
}

// regular function compare
if (instance1.regularFunc === instance2.regularFunc) {
  console.log("regular function", true); // true
} else {
  console.log("regular function", false);
}
```

দেখেন arrowFunc result false এবং regularFunc result true; আশা করি বুঝতে পারছেন, আপনি code করে করে দেখবেন, তাহলে বুঝতে আরও সুবিধা হবে।

_এইবার চলেন দেখি কিভাবে regular function গুলা Prototype হয়ে যায়_

```javascript
ThisClass.prototype.checkRegularFunc = function () {
  console.log("check regular share function");
};
const instance1 = new ThisClass("Mr. John Khan");
const instance2 = new ThisClass("Mr. John Khan");

if (instance1.checkRegularFunc === instance2.checkRegularFunc) {
  console.log(true); // result আপনি true পাবেন।
} else {
  console.log(false);
}
```

এইভাবে কিন্তু method add করা যায়, এই জন্য class এর Prototype হয়ে যায় regular function গুলা।
_আসলে এইগুলা কে আমি function বলেছি এইগুলা class এর method আগেও বলেছি_

### **Conclusion**

আমরা দেখেছি কিভাবে `Call`, `Apply`, এবং `Bind` methods `this` এর reference পরিবর্তন করে এবং কিভাবে এগুলো আমাদের function invocation এর behavior কে প্রভাবিত করে। class ক্ষেত্রে, আমরা লক্ষ্য করেছি যে `this` arrow functions এর জন্য lexical scope ব্যবহার করে, যা prototype methods এর সাথে সম্পর্কিত behavior এর সাথে ভিন্ন। regular function কিভাবে prototype method হয়ে যায় আবার reference হয়ে কাজ করে। আপনি যদি [1st-part](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/%22this%22%20keyword%20in%20JavaScript%201st%20Part.md) না পড়ে থাকেন তাহেল পড়ে আসেন।
লিখতে গিয়ে আমার ভুল হতে পারে; দয়া করে comments করে জানাবেন, আমি ঠিক করে নেব।
