---
date: 2024-10-10
title: Understanding Hosting in JavaScript A Comprehensive Guide!
description: এই পোস্টে আমরা hosting নিয়ে আলোচনা করবো। JavaScript এর core mechanics বুঝতে সাহায্য করবে। কিভাবে function এবং variable এ এর আচরণ কেমন করে? এই বিষয় গুলা নিয়ে deep drive করব।
author: Minhajul Islam (Minhaj)
published: true
---

## JavaScript এ Hosting কি?

Hosting আসলে কি? কেন দরকার কখন দরকার? এই Hosting নিয়ে আজ আলোচনা করব, চলেন;

Hosting কি? আপনি Google করলে অনেক definition পাবেন; তবে এত কঠিন definition আমার মনে থাকে না; তাই আমি মনে করি JavaScript এ hosting হল behavior বা আচরণ; JavaScript কিন্তু উল্টাপাল্টা language, তাই এর আচরণ (behavior) ধরা একটু ঝামেলা; সহজ ভাষায় behavior; আমার সাথে রাগ হন না, হ্যাঁ এমন করে বলছি; আমি নিজেও JavaScript এ code করি;

**বিঃদ্রঃ** একদম simple রাখতে চাচ্ছি variable এবং function এর name গুলো; আমি এত ভালো নাম দিতে পারি না। আচ্ছা, আপনি কিন্তু আমার সাথে পড়বেন, সাথে code করবেন, এতে আপনার ভালো হবে; আমি কি বললাম, আপনি মিলাতে পারবেন।

চলেন code করে দেখি কি behavior করে JavaScript;

```javascript
var aVar = "hosting";
function aFunc() {
  console.log("hosting func");
}
console.log(aVar); // hosting
aFunc(); // hosting func
```

একদম simple output ; চলেন আরেকটু change করি;

```javascript
console.log(aVar); // undefined
aFunc(); // hosting func
var aVar = "hosting";
function aFunc() {
  console.log("hosting func");
}
```

**এখন আবার কি এমন হল, বলেন তো?** `log` গুলোকে একটু উপরে দিতেই আবার কেমন behavior শুরু করে দিল; কাহিনী হল, এখানে JavaScript এর code দুই phase-এ run হয়; এইটা নিয়ে আমি বিস্তর একটা post লিখছি; execution context, বিশ্বাস করেন, এই post না পড়লে সামনে আগোনো বেশি ভালো বুঝবেন না। তারপরেও একটু বলি কেন এমন হল; এখানে code টা তো আপনি বুঝতে পারছেন? একটা variable এবং একটা function declare করেছি; JavaScript engine কি করে? যতগুলো আপনি variable declare করবেন, সবগুলোকে by default undefined দিয়ে রাখে; function গুলোকে function গুলো just নাম দিয়ে heap data-structure এ রাখে;

```javascript
1. aFunc: ƒ aFunc()
2. aVar: undefined
```

আপনি এই output টা browser এ source এ গিয়ে যদি একটু debug করেন তাহলে দেখেবেন একদম শুরুতেই এই দুইটা দেখাচ্ছে; শুরুতেই পাবেন এই জনা আমি name a দিয়ে দিছি যাতে শুরুর দিকে থাকে;

JavaScript engine নিয়ে যদি না জানেন, তাহলে কি হবে? আগে JavaScript engine নিয়ে পড়ে আসেন, এই আমার আগের post পড়ে আসেন; আবার বলেন যে please, আরে পড়তে আসলাম, hosting খালি এদিক সেদিক পাঠাচ্ছে মিনহাজ; সত্যি বলতে ওইগুলা হল basic.

আশা করবো পড়ে এসেছেন দুইটা post এ; এইবার তো নিঃস্বয় ধারণা পেয়েছেন কি ঘটনা; variable undefined, function, তাহলে output দিল, আপনি নিঃস্বয় পড়েছেন; ঐটা হল execution context। আচ্ছা, চলেন এইবার function কে log করে দেখি কি দেখায়;

```javascript
function aFunc() {
	console.log("hosting func");
}
console.log(aFunc);
// ƒ aFunc() {
  console.log("hosting func");
}
```

কিভাব ছিলেন output কি আসবে এইখানে তো পুরা function টাই দিয়ে দিছে; হ্যাঁ তাই তো হওয়ার কথা আমরা তো আর function টা কে call / invoke করি নাই তাই execution context হয় নাই তাই output ও আসে নাই।

এইবার আসেন আরেকটু change করি;

```javascript
console.log(aVar); // Uncaught ReferenceError: aVar is not defined
aFunc();
```

এইবার কোন কিছুই declare করি নাই; মানে হল যে javascript file aVar কে log করেছি aFunc() function কে call/ invoke করেছি। এইবার output দিল error : aVar is not defined. undefined এবং aVar is not defined দুইটা কিন্তু এক জিনিস না! যখন aVar undefined দিয়েছিল তখন সে stack থেকে memory নিয়ে বসেছিল; execution phase আসলে তখন print হয়ে যাইত; আর এখন বলতেছে aVar is not defined মানে হল javascript জানে যে aVar name কিছু আছে। তাই সে error দিয়েছি আপনি যদি function কে আগে call/invoke করেন দেখবেন যে error দিয়েছে।

আসেন এইবার code আরেকটু change করি;

এইবার কি হবে একটু ভাবেন আগে তো দেখলেন যে variable গুলা by default undefined বসায় রাখে। আপনি যদি এই code টা debug করেন; আপনি aVar variable যদি break point টা দেন next button a click করেন দেখবেন; Global এর মধ্যে এই রকম হয়ে আছে।

```javascript
aFunc: undefined;
aVar: undefined;
```

হ্যাঁ একদম ঠিক দেখতেছেন আগেই তো বলছিলাম যে আপনাকে execution context পরতে হবে; 1st phase সকল variable গুলা by default undefined পরের phase এসে ভালু assign করে তখন দেখে যে aFunc আসলে একটা function;

আমি আবার বলছি hosting বুঝার জন্য Execution context এর বিকল্প কিছুই নাই; তাই আমি বলছি Execution context এর post টা পড়েন; খুব helpful হবে please;

চলেন code টা কে একটু change করি;

```javascript
aFunc();
bFunc(); // Uncaught ReferenceError: bFunc is not defined
function aFunc() {
  function bFunc() {
    console.log("hosting func");
  }
}
```

code কি output দিবে আপনার মনে হয়? bFunc function টা কিন্তু hosted aFunc function সাথে তাই Global object এর মধ্যে bFunc function খুঁজে পাচ্ছে না; তাই error দিচ্ছে। আমি আগেই বলেছিলাম যে hosting হল behavior;

let এবং const নিয়ে কথা বলি নাই; এইগুলা নিয়ে এই post আর বলতেছি না; next post এ আমরা দেখব যে var, let এবং const comparison করে দেখবো;

## Conclusion

hosting কিছুই না; এইটা একটা behavior; javascript code output দেওয়ার দুইটা phase থাকে একটা creation phase and execution phase; এই টা নিয়ে আমি আগেই post দিয়েছি এবং deep drive করার চেষ্টা করেছি.
আপনার উচিৎ হবে আগে ঐটা পড়ে এইটা পড়ে দেখা সাথে হাতে কলমে code করে দেখা।
