---
date: 2024-10-09
title: Understanding Scope, Scope Chain & Lexical Environment A Comprehensive Guide!
description: Scope, Scope Chain & Lexical Environment সম্পর্কে বিস্তারিত আলোচনা করা হয়েছে। আপনি জানবেন Scope, Scope Chain & Lexical Environment কি, কিভাবে এটি কাজ করে? এই বিষয় গুলা নিয়ে in-depth যেতে চান, তাহলে এই পোস্টটি আপনার জন্য।
author: Minhajul Islam (Minhaj)
published: true
---

## Scope কি?

Scope হল একটি নির্দিষ্ট boundary (সীমানা)। ধরেন একটা ফুটবল মাঠে ভিতরে বল থাকলে player খেলতে পারে বাহিরে গেলে out হয়ে যায়; আশা করি এইটুকু বুঝতে পারছেন।
JavaScript এ দুই ধরনের Scope আছে;

1.  Global Scope
2.  Local Scope

## Global Scope:

JavaScript by default Global Scope code run হয় যা যেকোনো জায়গা থেকে access করা যায়;

চলেন code করে দেখি;

```javascript
var aVar = "Scope";
function aFunc() {
  console.log(aVar); // Scope
}
aFunc();
```

কি হল code এ মনে করেন এইটা একটা scope.js name এ একটা file এইটা মধ্যে আপনি যেই code গুলা line by line লিখবেন function এর ভিতরে ছাড়া এই গুলা global scope এ থাকবে; আপনি যদি aFunc এর ভিতরে আর অনেক function নেন তাহলেও access করতে পারবেন aVar variable কে এইগুলা নিয়ে আর কথা বলবো নিচে।

## Local Scope কি:

local scope হল ভাল করে মনে রাখবেন এইটা খুব important কিন্তু; যতক্ষণ না কোন function call বা invoke না হবে ততক্ষণ কোন local scope তৈরি হবে না। তার মানে হল এই যে যদি function call বা invoke হয় তাহলে local scope তৈরি হয়; !!মনে থাকে জেন!! এখন যদি একটা function এর মধ্যে আরও ৫ টা function থাকে তাহলে আরও ৫ টা scope তৈরি হবে।

চলেন code করে দেখি

```javascript
var aVar = "Scope";
function aFunc() {
  var bVar = "Local Scope";
  console.log(bVar); // Local Scope
}
aFunc();
```

আচ্ছা এইখানে কি দেখতে পাচ্ছেন; aVar variable এবং aFunc টা হল global scope এর মধ্যে; তবে আমি আগেই বলেছিলাম যে scope আরে কিছুই নয় boundary এইখানে boundary বলতে কি বুঝায় curly brackets এর মধ্যে যেই code গুলা আছে ওইগুলা হল একটা boundary এর মধ্যে so এইটা একটা scope;

চলেন একটা quiz দেন আগে নিচে পড়বেন না আগে নিজে চেষ্টা করেন না পারলে পড়বেন:

```javascript
var aVar = "Scope";
function aFunc() {
  var bVar = "Local Scope";
  console.log(bVar);
}
```

এইখানে কি কি scope আছে নিজে নিজে answear করেন? পারতে হবে না পারলে তো ভুলেই গেলেন কি হল তাহলে?
এইখানে একটা scope আছে সেটা global scope সেটা by default থাকে আগেই বলেছি; আর local scope তৈরি হয় নাই; কারণ হল যে aFunc name function টা তো call বা invoke এ করি নাই তাহলে local scope তৈরি হবে কেমনে ভুলে গেলে চলবে না; তাই দেখেন আমি কিন্তু output দেই নাই।

তাহলে আমরা কিন্তু scope কি বুঝলাম global scope by default এবং function call বা invoke হলে local scope হয়।

## Scope Chain কি জিনিস?

Scope Chain নতুন কিছু না তাহলে কি? চলেন দেখি কি?
মনে করেন আপনার দাদা এর কিছু সম্পদ আছে। সেগুলা আপনার দাদা কিন্তু চাইলে নিয়ে খরচ করতে পারেন simple জিনিস তার সম্পদ তিনি চাইলে খরচ করবেন; আবার দেখা গেল যে আপনার বাবা মানে হল আপনার দাদা এর সন্তান; এখন আপনার বাবা কিছু টাকা দরকার পরল তার কাছে টাকা নাই কিভাবে খরচ করবেন; তখন বাবা এর কাছ থেকে নিয়ে খরচ করতে পারবেন। এইবার দেখা গেল যে আপনার নিজের কিছু টাকা দরকার একটা R15 কিনবেন তখন টাকা দরকার আপনার বাবা এর কাছে টাকা নিতে পারবেন আবার আপনার দাদা এর কাছ থেকেও নিয়ে কিনতে পারবেন তাই নয় কি! তবে আপনার যদি কিছু সম্পদ থাকে ওইগুলা আপনার বাবা এবং দাদা কেউ নিতে পারবে না; আপনার বাবার ও সম্পদ কিন্তু আপনার দাদা নিতে পারবে না; উপর এর পানি যতই নিচে যায় সবাই ব্যাবহার করতে পারে তবে নিচের পানি উপরে উঠে না; যদিও বাস্থব জীবনে সম্পদ নিতে পারে আমি বুঝার সুবিধাতে বলছি; এইটা কি দেখেলেন যে আপনার কাছে টাকা নাই তো বাবার আছে তার কাছে নিলেন; বাবার নাই তো দাদার আছে তার কাছে নিলেন; দাদার কাছে নাই তো আপনার R15 কেনা হল না।

এইখানে কি একটা scope chain লক্ষ করতে পারছেন? হ্যাঁ ঠিক এইটাই হল scope chain;
চলেন code করে দেখি

```javascript
var সরকারি_সম্পদ = "সরকারি সম্পদ";
console.log(সরকারি_সম্পদ);
function দাদা() {
  var দাদার_সম্পদ = "দাদার সম্পদ";
  console.log(দাদার_সম্পদ, সরকারি_সম্পদ);
  function বাবা() {
    var বাবার_সম্পদ = "বাবার সম্পদ";
    console.log(বাবার_সম্পদ, দাদার_সম্পদ, সরকারি_সম্পদ);
    function আপনি() {
      var আপনার_সম্পদ = "আপনার সম্পদ";
      console.log(আপনার_সম্পদ, বাবার_সম্পদ, দাদার_সম্পদ, সরকারি_সম্পদ);
    }
    আপনি();
  }
  বাবা();
}
দাদা();
```

দেখেন আপনি ভাল করে আগে; কি হচ্ছে code? আচ্ছা এইবার আবার আপনি নিজে নিজেই একটা challenge নিয়ে বলেন যে কি কি scope create হইছে কি না? আর হইলেও কয়টা scope তৈরি হইছে?

দেখেন আপনি এমন জায়গায় বসেছেন আপনি সবকিছু access করতে পারতেছেন আপনার বাবা আপনার boundary এর মধ্যে সম্পদ নিতে পাচ্ছে না; যদি নিতে যায় তাহলে error খেয়ে বসে থাকবে কেমনে error খাবে এইটা নিয়ে বলতেছি। আশা করি code তো বুঝেছেন নাকি?
যার যার boundary যা আছে তা তো ব্যাবহার করতে পারবে; তবে তার ভিতরে আবার boundary থাকলে তখন কিন্তু তার ভিতরে যেতে পারবে না।

```javascript
var সরকারি_সম্পদ = "সরকারি সম্পদ";
function দাদা() {
  function বাবা() {
    function আপনি() {
      console.log(সরকারি_সম্পদ); // সরকারি সম্পদ
    }
    আপনি();
  }
  বাবা();
}
দাদা();
```

এইখানে হল কি? দাদা function কে call / invoke হওয়ার পর একটা scope তৈরি হয়; আবার বাবা function call / invoke হওয়ার পর একটা scope তৈরি হয়; একই ভাবে আপনি যখন call / invoke হবেন তখন একটা scope তৈরি হবে; আপনার local scope গিয়ে দেখবে একটা log আছে log করতে গিয়ে দেখবে যে সরকারি_সম্পদ variable আছে সেটা আপনার local scope এর মধ্যে খোঁজা শুরু করবে দেখবে নাই তখন আপনার বাবা এর boundary মানে আপনার বাবা এর সম্পদ এ দেখা শুরু করবে সেখানেও নাই; তখন দাদা এর কাছে গিয়ে দেখবে নাই তার পর আবার file মানে global এ গিয়ে খোঁজবে তখন গিয়ে পাবে; তারপর log হইয়ে যাবে। এইখানে যে chain টা তৈরি হল বুঝতে পারছেন তো?
আপনি -> বাবা -> দাদা -> সরকারি -> কিছু নাই!

এইবার চলেন একটু change করে দেখি কি হয়?

```javascript
function দাদা() {
  function বাবা() {
    function আপনি() {
      console.log(সরকারি_সম্পদ);
      //Uncaught ReferenceError: সরকারি_সম্পদ is not defined
    }
    আপনি();
  }
  বাবা();
}
দাদা();
```

একই রকম chain খুঁজতে থাকবে যখন দেখবে নাই তখন "Uncaught ReferenceError: সরকারি*সম্পদ is not defined" দিয়ে বসে থাকবে।
আপনি -> বাবা -> দাদা -> global -> null;
হ্যাঁ ঠিক global এ যখন সরকারি*সম্পদ variable নাই তখন খুঁজার মত আর কিছু নাই; কারন global এর কোন parent নাই তাই null কে point আছে।

এইভাবে scope chain হয়। আচ্ছা scope আমি তো scope তৈরির কথায় বলেছি তবে মনে রাখবেন প্রতিটি function call / invoke একটি করে execution context তৈরি হয় । এইটা নিয়ে আমি আগে post করেছিলাম।

## Lexical Environment কি জিনিস?

Lexical Environment আসলে কিছুই না আমরা আগেই পড়ে এসেছি; ভালো করে scope chain টা পড়লেই হয়ে যাবে; lexical name টাই একটু অন্যরকম। আমরা কি দেখেছি যে আপনি আপনার বাবার সম্পদ নিতে পারবেন এই যে আপনি আপনার বাবার সম্পদ নিতে পারেন আপনার বাবা আপনার সম্পদ নিতে পারে না; আপনি যে access করতে পারেন আপনার parent এর জিনিস গুলা plus আপনার সম্পদ মিলে হল lexical environment।
চলেন code করে দেখি

```javascript
function বাবা() {
  var বাবার_সম্পদ = "বাবার সম্পদ";
  console.log(বাবার_সম্পদ);
  function আপনি() {
    var আপনার_সম্পদ = "আপনার সম্পদ";
    console.log(আপনার_সম্পদ, বাবার_সম্পদ);
  }
  আপনি();
}
বাবা();
```

এইখানে আপনার বাবা হল global lexical environment তবে যখন বাবা function call / invoke হবে তখন আবার create হবে scope; আপনি হবেন lexical environment আপনার বাবা এর সাথে; মনে রাখবেন আপনার বাবার এর variable যখন ব্যাবহার করবেন তখন আবার নতুন একটা concept চলে আসবে সেটা হল closure এইটা নিয়ে next post হবে ইনশাআল্লাহ্‌।
আশা করি বুঝতে পারছেন আপনি এবং আপনার বাবা মিলে lexical environment.

## Scope Chain এবং Lexical Environment মধ্যে তাহলে difference কোথায়?

Great Question! চলেন দেখি;
Scope Chain নিয়ে তো বলেছি যে তার নিজের local scope এ খুঁজে আগে; যদি না পায় তাহলে সে তার outer এর কাছে যায় টাও যদি না পায় তাহলে তার outer এর কাছে যায়। মুল কথা হল যে সে খুঁজতেই থাকে কোথায় পাবে না পাইলে error দেয় যেটা আমরা উপরে দেখেছি; বাস্তবিক একটা example দিয়েছি; দাদা, বাবা এবং আপনি মিলে।
Lexical Environment হল যে function তার nested function দুইটা মিলে lexical environment তৈরি করে; এবং parent function এর Environment Record (যেখানে variable ও তাদের মান থাকে) এই জায়গায় closure চলে আসে এইটা নিয়ে আমি উপরে বলেছি ইনশাআল্লাহ্‌ next post হবে।

## Conclusion:

আমরা এই পোস্টে Scope, Scope Chain এবং Lexical Environment-এর কিভাবে হয় এবং তাদের মধ্যে সম্পর্ক নিয়ে আলোচনা করেছি। Scope আমাদের code এর ভেতর variable এবং function access boundary নির্ধারণ করে, যেখানে Global Scope এবং Local Scope দুটিই গুরুত্বপূর্ণ ভূমিকা পালন করে।

Scope Chain আমাদের বুঝতে help করে কিভাবে ভেরিয়েবল খোঁজা হয়, যেখানে একটি function ভেতর আরেকটি function ভেরিয়েবলগুলো access করা যায়। এটি একটি chain তৈরি করে, যা আমাদের nested scopes হয় আমরা বাস্তবিক example দেখেছি।

Lexical Environment parent এবং child মিলে make করে। parent এর variable গুলার, access করতে পারে child, যা closure-এর concept চলে আসে; ইনশাআল্লাহ্‌ next post হবে।

আশা করি বুঝতে পেরেছেন; তবে আপনি নিজে code করে debug করে দেখবেন তাহলে clear হয়ে যাবে।

আমার ভুল হতে পারে, comments করে আমাকে জানাবেন; আমি ঠিক করবো ইনশাল্লাহ।
