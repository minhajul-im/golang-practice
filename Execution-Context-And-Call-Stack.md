---
date: 2024-10-10
title: Understanding Execution Context & CAll Stack A Comprehensive Guide!
description: এই পোস্টে আমরা Execution Context এবং Call Stack এর গুরুত্ব এবং কাজ করার পদ্ধতি আলোচনা করবো। JavaScript এর core mechanics বুঝতে সাহায্য করবে। কিভাবে Execution Context create হয়? কিভাবে call stack যায় এই বিষয় গুলা নিয়ে deep drive করব।
author: Minhajul Islam (Minhaj)
published: true
---

# Execution Context & Call Stack

Execution Context & Call Stack এই দুইটা কি জিনিস? কিভাবে কাজ করে?  
Execution Context & Call Stack নিয়ে জানবো। এই গুলা [JavaScript Engine](https://minhajul.com) এর পার্ট; আপনি যদি [JavaScript Engine](https://minhajul.com) নিয়ে আমার আগের post টা না দেখে থাকেন তাহলে দেখে আসেন; ওই post টাতে আমি explain করেছি একটা JS file কিভাবে machine code পরিণত হয়.

### Execution Context কি?

Execution Context একটি environment যেখানে JavaScript Code run করা যায়। যখন JavaScript code run করা হয়, তখন Execution Context এর মধ্য দিয়েই execute হয়। Execution Context তিন ধরনের হয়; Global Execution Context, Function Execution Context, Eval Execution Context.

### Call Stack কি?

JavaScript code execution sequence গুলো maintenance করার জন্য ব্যবহার করা হয়। এর মূল কাজ হলো function গুলোর invocation এবং তাদের execution এর order track করা। এটা একটা stack data-structure; LIFO (first-in-last-out)।

### Execution Context কি কি?

1. **Global Execution Context**: এটি একটি default execution context। একটা JS file এর মধ্যে যেই code গুলো কোনো function এর ভিতরে নাই, সেগুলো Global Execution Context এ থাকে।
   - এটি একটি Global Object তৈরি করে; যেটা কে আমরা Browser Window এবং Node.js Global নামে জানি। Global Object এটা দেখতে কেমন আপনি `console.log(window/global);` লিখে যদি দেখেন, দেখবেন অনেক গুলা function আছে ওর মধ্যে।
   - this নামে একটি variable তৈরি করে; যা Global Object বা Window Object কে নির্দেশ করে। আপনি যদি দেখেন থাকেন `console.log(window/global);` এইটা আবার আপনি `console.log(this === window/global);` এইটা একটু দেখেন, দেখবেন result `true` দেখাবে; এই জন্য this নামে একটা variable কে নির্দেশ করে। আমি `this` নিয়ে আলাদা একটা post করবো।
2. **Function Execution Context**: Function Execution Context হল, যে যখন কোন function call / function invocation করা হয়; তখন সেই function এর জন্য JS engine একটি নতুন Execution Context তৈরি করে। প্রতিটি function এর জন্য নিজস্ব একটা করে Execution Context তৈরি করে; একাধিক function থাকলে একাধিক Execution Context তৈরি করবে। এইটা কিভাবে হয় সেটা আমরা দেখবো; একই ভাবে function call / function invocation সময় নিউ একটা করে this create হবে যা Global Object বা Window Object কে নির্দেশ করে।

3. **Eval Execution Context**: সত্যি বলতে, আমি কখনই এটা দিয়ে কোনোদিন code করি নাই। যখন আমি Execution Context নিয়ে পড়াশোনা করছিলাম, তখন নাম জেনেছিলাম।

### Execution Context কিভাবে কাজ করে?

Execution Context মূলত দুইটি phase code execution করে। প্রতিটি Execution Context দুইটি phase থাকে।

- Creation Phase
- Execution Phase

আমরা এইটুকু code এর explanation করবো আশা করি তাতেই বুঝে যাব; সব একই রকম ভাবে execution হয়।

```javascript
var hello = "hello world";
function func(parameter) {
  var age = 22;
  console.log(age, parameter);
}
console.log(hello);
func(argument);
```

### Creation Phase: কি হয় চলেন দেখি।

1. this variable create হয় যা Global Object বা Window Object কে নির্দেশ করে। আমি আগেই বলেছি।
2. code এর মধ্যে সকল variable এর value default undefined assign করে এবং সকল function declaretion গুলা heap memory store করে।

চলেন আমরা যে উপরে code টুকু করেছি সেটা creation phase কেমন হবে?

```
var hello = undefined;
func: ƒ func() // এই func টা browser এ open করলে দেখবেন অনেক কিছু আছে!
```

এই রকম দেখতে পারবেন; variable default value হিসাবে undefined assign করে এবং function declaration গুলা heap data structure এর মধ্যে memory তে store করে।
আপনি যদি code browser run করে debug করেন তাহলে দেখতে পারবেন। debug করবেন কেমনে সেটা না জানলে আমি একটা link দিয়ে দিচ্ছি আপনি দেখে নিয়েন [click](https://bdit.community/t/topic/3945);

### Execution Phase: কি হয় চলেন দেখি!

1. JS engine লাইন বাই লাইন code execute করে; ঠিক এই সময় variable value গুলা assign করে এবং function গুলা function call / function invocation করে।
2. যখন function call / function invocation করে তখন function টা call stack এ push হয়ে যায়। call stack নিয়ে আমরা দেখবো।

code টা কেমন হল সেটা একটু দেখি।

```
var hello = "hello world";
console.log(hello); // hello world
func: ƒ func()
```

আমরা দেখলাম যে variable এর value assign করল, সাথে `log` করে result দেখালো। এইটাই হল Global Execution. আপনি চাইলে আরো arithmetic operation করে দেখতে পারেন।

এবার আসি function, কিন্তু আগের মতই আছে। এইবার আসলো Function Execution, আবার নতুন করে একটা Execution Context create হবে। আগে বলেছি যে প্রতিটি function call / function invocation সময় একটা new Execution Context create হয়।
আবার এইটা কিন্তু Creation Phase মনে রাখতে হবে।

```
function func (parameter = undefined){
	var age = undefined;
	console.log(age, parameter);
};
func(argument);// এখানে আপনি যেকোনো value pass করতে পারেন
```

এর মধ্যে অন্য function থাকতে পারে; মনে রাখতে হবে আমরা কিন্তু execution phase থেকেই আবার creation phase চলে আসলাম; এইজন্যই আগেই বলছি; প্রতিটি function এর জন্য নিজস্ব একটা করে Execution Context তৈরি করে; একাধিক function থাকলে একাধিক Execution Context তৈরি করবে।  
**Note** কিন্তু আমরা তো একটা argument pass করেছি, তার পরেও কেন parameter undefined দেখাচ্ছে তার কারণ হল যে creation phase value assign করতে পারবে না।

এইবার হবে execution phase-

```
function func(parameter = argument) {
	var age = 22;
	console.log(age, parameter); // 22, argument
}
func(argument); // এখানে আপনি যেকোনো value pass করতে পারেন
```

দেখেন এইবার সব গুলা value assign হয়ে গেছে, execute করে result হয়ে গেছে। যখন execution শেষ হয়ে যায়, তখন call stack থেকে চলে যায়।

Execution Context কিন্তু execute করার নিয়ম কিন্তু একই রকম; সেটা global বা function দুইটা phase কাজ করে creation & execution.
আপনি যদি visual দেখতে চান, আমি link দিলাম, আপনি দেখেন। [click](https://cdn.hashnode.com/res/hashnode/image/upload/v1662578709635/I7cKZhiOn.gif?w=1600&h=840&fit=crop&crop=entropy&auto=format,compress&gif-q=60&format=webm)

### Call Stack কিভাবে কাজ করে?

এইটা একটা data structure। আপনাকে visual দেখাতে পাচ্ছি না তবে link দিলাম আপনি দেখে নিয়েন [click](https://vaibhavgupta.me/wp-content/uploads/2018/01/callstack.gif).

চলেন একটা example দেখি, মনে করেন আপনাকে ১০ টা কাচের প্লেট এবং একটা box দেওয়া হল এবং বলা হল যে আপনি এইগুলা box এর ভিতরে রাখেন। আপনি কি করবেন? তখন box এর একটা মধ্যে একটা রাখবেন। আবার যদি বলা হয় যে আপনি প্লেট গুলা বের করেন, আপনি কি সব নিচের টা বের করতে পারবেন? অবশ্যই না! আপনাকে উপর থেকে একটা একটা করে বের করতে হবে। যেই প্লেট টা আপনি সবার শেষে রাখলেন, ওইটাই আগে বের করলেন এবং যেই প্লেট টা সবার আগে রাখলেন, সেই প্লেট সবার শেষে বের করলেন। এইটাই হল call stack data structure LIFO:(last-in-first-out)

code করে দেখি একটু চলেন:

```javascript
function funcOne() {
  funcTwo();
  console.log("funcTwo");
}

function funcTwo() {
  funcThree();
  console.log("funcThree");
}

function funcThree() {
  funcFour();
  console.log("funcFour");
}

function funcFour() {
  console.log("function end");
}

funcOne();
```

##### কি দেখা যায় তো?

**funcOne()**:

- প্রথমে `funcOne` function কল করা হয়।

- যখন `funcOne` কল করা হয়, তখন এর জন্য একটি **Execution Context** তৈরি হয়।

- এটি `funcTwo()` কে কল করে, তাই `funcOne` এর Execution Context call stack এ যুক্ত হয়। মনে রাখতে হবে এখনে কিন্তু `log` হয় না।

**funcTwo()**:

- `funcTwo` ফাংশন শুরু হয় এবং এর জন্যও একটি নতুন **Execution Context** তৈরি হয়।

- এটি `funcThree()` কে কল করে, তাই `funcTwo` এর Execution Context call stack এ যুক্ত হয়। মনে রাখতে হবে এখনে কিন্তু `log` হয় না।

**funcThree()**:

- `funcThree` ফাংশন শুরু হয় এবং এর জন্য একটি নতুন **Execution Context** তৈরি হয়।

- এটি `funcFour()` কে কল করে, তাই `funcThree` এর Execution Context call stack এ যুক্ত হয়। মনে রাখতে হবে এখনে কিন্তু `log` হয় না।

**funcFour()**:

- `funcFour` ফাংশন শুরু হয় এবং এর জন্যও একটি নতুন **Execution Context** তৈরি হয়।

- এখানে একটি `console.log` `"function end"` প্রিন্ট হয়।

- `funcFour` শেষ হলে, এর Execution Context call stack থেকে পপ হয়।

তাহলে কি দেখলেন একদম শেষে funcFour call হল এবং প্রিন্ট হয়ে গেলো। চলেন `log` গুলা দেখি একটু কি order এ প্রিন্ট হল।

```
function end
funcFour
funcThree
funcTwo
```

### Conclusion

Execution Context এবং Call Stack বুঝতে পারা JavaScript এর core mechanics grasp করার জন্য অত্যন্ত গুরুত্বপূর্ণ। Execution Context বিভিন্ন ধরণের environment তৈরি করে যেখানে JavaScript কোড কার্যকরী হয়, এবং Call Stack আমাদের ফাংশনগুলোর invocation এর অর্ডার ট্র্যাক করে।
JavaScript Engine কিভাবে কাজ করে এবং কিভাবে আমাদের কোড বিভিন্ন ধাপে execution করে। আশা করি, এই আর্টিকেলটি আপনাকে Execution Context এবং Call Stack এর মধ্যে সংযোগ স্থাপন করতে সহায়তা করবে।

যদি ভুল হয় থাকে তাহলে comments করে জানাবেন আমি ঠিক করবো। আমি একটা video link দিয়ে দিচ্ছি আপনি দেখতে পারেন code & debugging করে দেখিয়েছে। [click](https://www.youtube.com/watch?v=iLWTnMzWtj4)
