---
layout: section
color: slate
---

# Test-Driven Development

The most efficient way to write the _vast majority of code_.

---
layout: top-title
color: slate
align: l
---

<v-click>
<StickyNote title="Working in Gears" v-drag="[691,269,268,180,-5]">
Kent Beck uses the transmission in a car as a metaphor. 1st gear could represent
running the feedback loop after writing each single line of code, 5th gear could
represent writing multiple tests at once, before implementing the behaviour on
one go.
</StickyNote>
</v-click>

:: title ::

# Test-Driven Development

:: content ::

No matter what you work on, strive to get the best possible feedback loop

## Optimize for Shortest Feedback Loop


- Fast feedback loop allows working in small increments.
- You know immediately when code doesn't produce the intended effect.
- Truncate unproductive code paths
- Use smaller feedback loops when working in areas of high uncertainty.


---
layout: top-title
color: slate
align: l
---

:: title ::

# Test-Driven Development - Applicability

:: content ::

Works well when a small piece of code can provide **_fast and relevant_** feedback.

## Less useful scenarios

- UI work
  - Relevant feedback is looking at the page.
  - Fast feed comes from Live/hot reload in browser.
- And maybe (not my area of expertice)
  - AI Model Development
  - Game Development

For most of the behaviour in the system, the best feedback loop is achieved from
running small snippets of code automatically, e.g., when you save.

<!-- --- -->
<!-- src: ./tdd-not-about-unit-tests.md -->
<!-- --- -->

<!-- --- -->
<!-- src: ./tdd-web-applications.md -->
<!-- --- -->

---
layout: top-title
color: slate
align: l
---

:: title ::

# Test-Driven Development - A philosophy

:: content ::

<div class="ns-c-tight">


- Write tests describing **behaviour**; not implementation details
- Behaviour belong in a specific _layer of the app_.
- Behaviour **can be** business rules
- Behaviour **is not** the entire stack
- Behaviour **can be** technical
  - Verifying access token renewal with exponential backoff on failure.
- Behaviour can result from **collaboration of multiple units**
  - Verify them **as a whole**, not individually.
  - (Except when fast feedback is needed on individual units)
- Only mock at layer boundaries
- Don't mock 3rd party libraries

</div>
