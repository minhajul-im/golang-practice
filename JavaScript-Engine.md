---
date: 2024-10-09
title: Understanding JavaScript Engine (V8) A Comprehensive Guide!
description: JavaScript Engine (বিশেষ V8) সম্পর্কে বিস্তারিত আলোচনা করা হয়েছে। আপনি জানবেন JavaScript Engine কি, কিভাবে এটি কাজ করে এবং কেন এটি Modern Web Development এ অত্যন্ত গুরুত্বপূর্ণ। V8 Engine এর কাঠামো, এর বিভিন্ন components যেমন Lexer, Parser, Interpreter, এবং JIT Compiler নিয়ে আলোচনা করা হবে। এছাড়াও, Execution Context, Call Stack, এবং Memory Management এর ভূমিকা কিভাবে JavaScript code কার্যকারিতাকে প্রভাবিত করে, সেই বিষয়গুলোও অন্তর্ভুক্ত রয়েছে। আপনি যদি JavaScript Engine এর in-depth যেতে চান, তাহলে এই পোস্টটি আপনার জন্য।
author: Minhajul Islam (Minhaj)
published: true
---

# JavaScript Engine কি?

JavaScript Engine হলো JavaScript Language এর code run করার জন্য একটা machine যেখানে আমরা JavaScript code run করতে পারি। কেন দরকার এই machine এর? এটা যদি বুঝতে চান তাহলে আমি আগে একটা পোস্ট করেছিলাম,আপনি এইটা একটু পড়ে আসেন [Runtime Environments](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/Runtime-Environments.md); আমি recommend করব; তারপর এইটা continue করেন, আপনার জন্য ভালো হবে।

### JavaScript Engine কোনগুলো?

মার্কেটে অনেক ধরনের JavaScript Engine আছে; তবে প্রত্যেকটা Browser এ একটি করে JavaScript Engine রয়েছে। কয়েকটি জনপ্রিয় জাভাস্ক্রিপ্ট ইঞ্জিনগুলোর নাম নিচে দেয়া হলোঃ

- Google Chrome – V8 Engine
- Mozilla Firefox – SpiderMonkey
- Safari – JavaScriptCore (Nitro)
- Microsoft Edge – Chakra
- Node.js & Deno – V8 Engine (Server-Side)
- Bun.js - WebKit's JavaScriptCore (Server-Side)

এই Engine গুলোর মধ্যে V8 Engine কে বর্তমানে সবচেয়ে জনপ্রিয় এবং দক্ষ হিসেবে ধরা হয়। আমি এই লেখায় সবগুলো ইঞ্জিন নিয়ে আলোচনা করা সম্ভব নয়। তাই আমি এখানে চেষ্টা করবো শুধুমাত্র V8 Engine নিয়ে কথা বলতে। এই Engine কিভাবে কাজ করে বুঝতে পারলে আশা করছি বাকি Engine গুলো বোঝা খুব বেশি কঠিন হবে না আপনার জন্য!

### কিভাবে কাজ করে JavaScript Engine?

আমরা যখন কোন JavaScript file Runtime Environments এর মধ্যে run করি; তখন Runtime Environments এর একটা পার্ট হলো JavaScript Engine। Runtime Environments কি? [Runtime Environments](https://github.com/minhajul-im/Blog-Post/blob/main/javascript/Runtime-Environments.md); JS file টা Runtime Environments যাওয়ার সাথে সাথে Runtime Environments এই JS file কে V8 Engine এর কাছে দিয়ে দেয় এবং বলে দেয়, V8 Engine, এইটা আপনার কাজ, আপনি করেন। V8 Engine এই file টাকে নিয়ে তার নিজের ভিতরে অনেক কিছু ব্যবহার করে এই file টাকে machine code করে Computer কে দিয়ে দেয়।

### V8 Engine কি কি থাকে একটা JS file কিভাবে machine code হয়?

আমি আগেই বলেছি V8 Engine নিয়ে বলবো; একটা JS file আসার সাথে সাথে V8 Engine নিয়ে নেয় নিজের কাছে, তারপর যেই কাজগুলো হয়।

1. **Lexer**: এইটা V8 Engine এর একটি component, যার কাজ হলো JS file এর মধ্যে থাকা code গুলোকে tokenize করা। Code দেখতে এমন দেখায় Tokenization of `console.log("hello world");`

   - **`console`** - Identifier
   - **`.`** - Punctuation
   - **`log`** - Identifier
   - **`(`** - Punctuation
   - **`"hello world"`** - String Literal
   - **`)`** - Punctuation
   - **`;`** - Punctuation
   - `["var", "x", "=", "5", "+", "3", ";"]` একটা variable declare এরকম দেখায়।

2. **Parser**: এই Tokenize করা code গুলোকে Parser ব্যবহার করে AST বা (Abstract Syntax Tree) তৈরি করে। আচ্ছা, চলেন দেখি একটা function কে AST করলে কেমন দেখায়;

   - JavaScript Code:

   ```javascript
   function hello() {
     console.log("hello world");
   }
   ```

   - AST বা (Abstract Syntax Tree):

   ```
   Program
   └── FunctionDeclaration
       ├── id: Identifier (hello)
       ├── params: []
       └── body: BlockStatement
           └── body:
               └── ExpressionStatement
                   └── expression: CallExpression
                       ├── callee: MemberExpression
                       │   ├── object: Identifier (console)
                       │   └── property: Identifier (log)
                       └── arguments:
                           └── Literal ("hello world")
   ```

   আপনার ইচ্ছা হলে আপনি এই site গিয়ে explore করতে পারেন [AST](https://astexplorer.net/);

3. **Interpreter**: Interpreter কি? সোজা কথা, দোভাষী। Interpreter আমাদের সোর্স কোডের প্রথম লাইন থেকে শুরু করে লাইন বাই লাইন low-level language রূপান্তর করে। এর সুবিধা হলো কোডের যেকোনো লাইনে যদি কোনো error পাওয়া যায়, তাহলে Interpreter আমাদেরকে সাথে সাথেই জানিয়ে দিতে পারে। অর্থাৎ, Debuging করা অনেক সহজ হয়। V8 Engine এর Interpreter Ignition নামে পরিচিত। Ignition কিভাবে কাজ করে?

   - **Bytecode generate**: AST থেকে Ignition bytecode তৈরি করে।
   - **Memory Usage**: Interpreter ভেতরে Bytecode execute করার জন্য memory হিসেবে রেজিস্টার ব্যবহার করা হয়। Ignition এখানে প্রত্যেকটা object এর জন্য একেকটা ‘shape/আকৃতি’ তৈরি করে; যেটা আসলে object এর মত ছাড়া আর কিছুই না।
   - **Inlining Caching**: Object ‘shape/আকৃতি’ তৈরি করে রাখার কারণে পরবর্তীতে ‘Inlining Caching’ ও অন্যান্য optimization করতে সুবিধা হয়।
   - **Profiler/Monitor**: একইসাথে profiling করে; এটি দেখে কোন অংশগুলি (hot code path) detect করে। কিভাবে detect করে hot code path? Interpreter যখন লাইন ধরে ধরে কোড execute করতে থাকে, profiler তখন কোন স্টেটমেন্ট কতবার run হচ্ছে এটার হিসাব করে রাখে। Code এর একই অংশ যদি একাধিকবার execute হয়, profiler তখন এটিকে ‘ওয়ার্ম (Warm)’ হিসেবে শনাক্ত করে। এই সংখ্যাটি বাড়তে থাকলে এক সময় এটিকে ‘hot code path’ বলা হয়। অর্থাৎ, programmer যে অংশগুলো সবচেয়ে বেশি সংখ্যকবার run করছে, monitor/profiler সেগুলোকে বের করে আনে; monitor/profiler দ্বারা শনাক্ত হওয়া hot code path পার্টগুলো JIT (Just-In-Time) কম্পাইলারের কাছে পাঠানো হয়।

4. **JIT (Just-In-Time)**: Hot-Code-Path থেকে পাওয়া একাধিকবার execute হওয়া code অংশগুলো Baseline Compiler দিয়ে Bytecode generate করে। Bytecode কিন্তু মেশিন কোডের মতো fast নয়। অনেকবার run হওয়া কোড যদি সরাসরি machine code এ রূপান্তর করে সেটাকে বারবার ব্যবহার করা যায়, তাহলে program অনেক ভালো পারফর্ম করবে। Optimizing Compiler আমাদের জন্য এই কাজটিই করে দেয়।

5. **TurboFan**: TurboFan Ignition এবং JIT থেকে যেই code গুলো পায়, সবগুলোকে highly-optimized machine code generate করে execute করে।

6. **Execution Context**: Execution Context একটি environment যেখানে JavaScript Code run করা যায়। যখন JavaScript code run করা হয়, তখন Execution Context এর মধ্য দিয়েই execute হয়। আসলে Execution Context নিয়ে আলাদা একটা পোস্ট লিখব, তখন code করে দেখাবো আমি; এটা একটা বিশাল পার্ট; Just high-level overview দেওয়ার জন্য দিলাম; Execution Context তিন ধরনেরঃ

   1. **Global Execution Context**: এটি একটি default execution context। একটা JS file এর মধ্যে যেই code গুলো কোনো function এর ভিতরে নাই, সেগুলো Global Execution Context এ থাকে। Browser Window নামে থাকে এবং Node.js Global নামে থাকে।
   2. **Function Execution Context**: Function Execution Context হল, যে যখন কোন function call / function invocation করা হয়; তখন সেই function এর জন্য JS engine একটি নতুন Execution Context তৈরি করে। প্রতিটি function এর জন্য নিজস্ব একটা করে Execution Context তৈরি করে; একাধিক function থাকলে একাধিক Execution Context তৈরি করবে।
   3. **Eval Execution Context**: সত্যি বলতে, আমি কখনই এটা দিয়ে কোনোদিন code করি নাই। যখন আমি Execution Context নিয়ে পড়াশোনা করছিলাম, তখন জেনেছিলাম।

7. **Call Stack**: JavaScript code execution sequence গুলো maintenance করার জন্য ব্যবহার করা হয়। এর মূল কাজ হলো function গুলোর invocation এবং তাদের execution এর order track করা। এটা একটা stack data-structure; LIFO (first-in-last-out)। এটাও একটা বড় পার্ট; আমি আগেই বলেছি; Execution Context এর সাথে অনেকটাই related, আমি এই দুইটা একসাথে পোস্ট করব।

8. **Heap**: এটিও একটি data-structure; Dynamic Allocation। JS code করি, তখন কিন্তু আমরা কোনো type বলে দেই না; যতটুকু memory দরকার পরে heap data-structure system থেকে pull করে নিয়ে নেয়। Mark & Sweep algorithm ব্যবহার করে garbage collection করে। এই Mark & Sweep algorithm এর মধ্যেও অনেক কিছু করে; আপনার ইচ্ছা হলে আপনি একটু Google করে দেখতে পারেন।

### Conclusion

এই পোস্টে আমরা JavaScript Engine এর কার্যক্রম এবং বিশেষ করে V8 Engine এর বিভিন্ন components নিয়ে আলোচনা করেছি। আমরা দেখেছি কিভাবে V8 Engine code কে প্রথমে Lexer এবং Parser এর মাধ্যমে AST বা (Abstract Syntax Tree) করে; পরে Interpreter এবং JIT Compiler থেকে সবগুলোকে নিয়ে TurboFan highly-optimized machine code generate করে execute করে।

Execution Context এবং Call Stack এর মতো concept এর উপর কিছু ধারনা দিয়েছি। এছাড়াও, Heap এবং Garbage Collection নিয়েও একটু লিখেছি।

পরবর্তীতে, আমি Execution Context এবং Call Stack নিয়ে আরও in-depth আলোচনা করব; আশা করি এই পোস্টটি আপনাদের JavaScript Engine নিয়ে কিছু নতুন ধারণা দিয়েছে।

আপনার যদি পুরো জিনিসগুলো একটা চার্ট আকারে দেখতে চান তাহলে আমি link দিলাম, আপনি দেখে নিন [JavaScript Engine](https://docs.google.com/presentation/d/1HgDDXBYqCJNasBKBDf9szap1j4q4wnSHhOYpaNy5mHU/edit#slide=id.g1357e6d1a4_0_58)

আমার ভুল হতে পারে, জানাবেন comment করে, আমি ঠিক করবো।
