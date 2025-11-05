---
layout: top-title
---

:: title ::

# Future Vision

<v-click>

Different strategies can be used for different testing needs. For geolocation:

</v-click>

- One implement could return a single point.
- Another could replay a GPX track
  - E.g., your application implements geofencing in the UI.

:: content ::

Gost-DOM provides core browser functionalities. Plugins can support more web
APIs. E.g.

- Local storage/session storage/DB
- Geolocation


---
layout: top-title
color: slate
---

<v-drag pos="351,466,247,34">
Donations are also welcome. 
</v-drag>

<ArrowHeads color="pink" width="300px" v-drag="[279,278,346,214]" />
<StickyNote v-drag="[583,418,196,66,10]">
EU VAT invoices can be provided
</StickyNote>

:: title ::

# How to Contribute

:: content ::

- Help build new features
- Feedback
- Create examples
- Spread the word
- Support for non-core web APIs
- Building a web site

---
layout: top-title
color: slate
---

:: title ::

# Inspiration

:: content ::

- [Testing Library](https://testing-library.com/)
  - Writing tests coupled to accessibility 
- [jscom](https://github.com/jsdom/jsdom)
- Unmaintained
  - [Zombie](https://zombie.js.org/)
  - [PhantomJS](https://phantomjs.org/)


---
layout: top-title
color: slate
---

:: title ::

# Shoutout

:: content ::

- [rogchap] Original author of v8go
- [tommie] Best current fork of v8go
  - Automatic updates of v8 from chromium sources
- Maintainers of [w3c/webref] always helpful and friendly in replying to
  questions, and offering help in extracting information.

[rogchap]: https://github.com/rogchap
[tommie]: https://github.com/tommie
[w3c/webref]: https://github.com/w3c/webref

---
layout: top-title
color: slate
---

:: title ::

# Support Libraries

:: content ::

A few support projects have 0

- [gost-dom/shaman]

  Helper library coupling tests to user interaction rather than
  implementation details. Facilitates safe refactoring and encourages testing
  for accessibility.

- [gost-dom/surgeon]

  Surgically replace dependencies. Helps mocking at layer boundaries.

- [gost-dom/v8go]

  Fork of v8go used by Gost-DOM (don't use), exposing essential V8 features not yet available in v8go.

- [gost-dom/webref]

  Exposes parts of [w3c/webref] specs as Go data structures.

<StickyNote v-drag="[751,246,195,86,11]">

E.g., ESM support. Should eventually make its way into [tommie's fork].

</StickyNote>

<StickyNote v-drag="[522,414,266,99,-15]">

Started as part of the code generator, but through continuous refactoring, a
general package seemed to emerge more or less by itself.

</StickyNote>

[gost-dom/browser]: https://github.com/gost-dom/browser
[gost-dom/v8go]: https://github.com/gost-dom/v8go
[gost-dom/shaman]: https://github.com/gost-dom/shaman
[gost-dom/webref]: https://github.com/gost-dom/webref
[gost-dom/surgeon]: https://github.com/gost-dom/surgeon
[w3c/webref]: https://github.com/w3c/webref
[tommie's fork]: https://github.com/tommie/v8go
