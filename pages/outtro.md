---
layout: section
color: slate
---

# Outtro

---
layout: top-title
color: slate
---

:: title ::

# It really does ... (what others claim)

:: content ::

- Completely predictable execution
- Auto wait for all
  - Including back-end processes
- You can reliably **test for the absence** of a reaction with zero waiting.

---
layout: top-title
color: slate
---

:: title ::

# Support Libraries

:: content ::

A few support projects have spawned from the project

- [gost-dom/shaman]

  Helper library coupling tests to user interaction rather than
  implementation details. Facilitates safe refactoring and encourages testing
  for accessibility.

- [gost-dom/surgeon]

  Surgically replace dependencies. Helps mocking at layer boundaries.

- [gost-dom/v8go]

  Fork of v8go used by Gost-DOM (don't use), nesting ground for adding V8 features not yet available in v8go.

- [gost-dom/webref]

  Exposes parts of [w3c/webref] specs as Go data structures.

<StickyNote v-drag="[751,246,195,86,11]">

E.g., ESM support. Should eventually make its way into [tommie's fork].

</StickyNote>

<StickyNote v-drag="[593,407,266,99,-15]">

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
  - Proof that a usable browser-like environment is achievable
- Unmaintained (the original headless browsers)
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
- Maintainers of [w3c/webref] always helpful and friendly and helpful replying to
  questions, and offering help.
- [slidev] Presentations that can be committed to git
- [neversink]: Slidev theme

[rogchap]: https://github.com/rogchap
[tommie]: https://github.com/tommie
[w3c/webref]: https://github.com/w3c/webref
[slidev]: https://sli.dev/
[neversink]: https://github.com/gureckis/slidev-theme-neversink/tree/main

---
layout: top-title
color: slate
---

:: title ::

# Want to Contribute?

:: content ::

<v-clicks>

- Help build new features
- Support for non-core web APIs
- Feedback
- Create examples
- Spread the word
- Building a web site

</v-clicks>

<v-click>
<ArrowHeads color="pink" width="300px" v-drag="[279,278,346,214]" />
<v-drag pos="351,466,247,34">
Donations are also welcome. 
</v-drag>
</v-click>

<v-click>
<StickyNote v-drag="[583,418,196,66,10]">
EU VAT invoices can be provided
</StickyNote>
</v-click>
