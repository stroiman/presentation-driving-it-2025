---
layout: top-title
color: slate
align: l
---

<v-click>
<StickyNote title="Working in Gears" v-drag="[560,338,367,117,-6]">
Kent Beck uses the transmission in a car as a metaphor. 1st gear is one line of code at a
time, 5th gear could represent writing multiple tests at once, and implementing
them all in one go.
</StickyNote>
</v-click>

:: title ::

# Test-Driven Development

:: content ::

A way to increase developer efficiency, misunderstood even by many
practitioners.

## Focus on Feedback

- Fast feedback loop allows working in small increments.
- You know immediately when code doesn't produce the intended effect.
- The size of the feedback loop is adapted to your level of confidence with
  specific area of code.

---
layout: top-title
color: slate
align: l
---

:: title ::

# Feedback-Driven Development - Applicability

:: content ::

Works well when a small piece of code can provide **_fast and relevant_** feedback.

## Less useful scenarios

- UI work
  - Relevant feedback is visual inspecting, does it _look right_.
  - Live/hot reload in browser provide fast feedback.
- Training AI models
  - Definitely not fast
- Game engines,
  - I would guess

---
layout: top-title
color: slate
align: l
---

:: title ::

# Test-Driven Development - not about Unit Tests

:: content ::

These are subjective opinions rooted in **_years of experience_**. 

- Don't test single units in isolation.
- Test how units collaborate to provide the _behaviours_ of the application.
  - Don't call a _Controller Method_
  - Send an HTTP request, verify the response, and side effects in the system.
- Only mock at layer boundaries.

## The primary measure of the quality of a test suite

- Does it provide fast feedback?
- Does it enable safe refactoring?

---
layout: top-title
color: slate
---

<v-click>
<v-drag pos="395,346,480,123">
<SpeechBubble position="l" color="orange-light" shape="round" class="text-center"
width="450">
<div class="text-center">
As a developer<br />
In order to work effeciently with code<br />
I want a fast feedback loop for all of my code.<br />
</div>
</SpeechBubble>
</v-drag>

<v-drag pos="276,386,93,111">
<Ghost :size="100" mood="excited" color="#cbd5e1"/>
</v-drag>
</v-click>

:: title ::

# TDD and Web Applications

:: content ::

<div grid grid-cols-3 gap-3 h-75>

<v-click>
<div border="2 solid black/5" rounded-lg overflow-hidden bg="black/5" backdrop-blur-sm h-full>
  <div flex items-center bg="black/8" backdrop-blur px-3 py-2 rounded-md>
    <div text-amber-300 text-sm mr-2 />
    <div font-semibold>
      1990s - SSR
    </div>
  </div>
  <div px-4 py-3>
    <div flex flex-col gap-3>
      <div>
        <div text-xs opacity-70 mb-2>
        The server generates HTML pages. Over time, JavaScript plays an
        increasing role.
      </div>
      </div>
      <div>
        <div text-sm font-medium>Testing strategy</div>
        <ul text-xs opacity-70><li>Submit GET/POST requests and verify rendered
        HTML</li>
        <li>Call controller methods, and pray that it works</li></ul>
      </div>
    </div>
  </div>
</div>
</v-click>

<v-click>
<div border="2 solid black/5" rounded-lg overflow-hidden bg="black/5" backdrop-blur-sm h-full>
  <div flex items-center bg="black/8" backdrop-blur px-3 py-2 rounded-md>
    <div text-amber-300 text-sm mr-2 />
    <div font-semibold>
      2010's - SPAs
    </div>
  </div>
  <div px-4 py-3>
    <div flex flex-col gap-3>
    <div text-xs opacity-70>JavaScript handles all UI logic, communicating with
    a backend through an API.</div>
      <div>
        <div text-sm font-medium>Testing Strategy</div>
        <div text-xs opacity-70><ul><li>UI code is tested in isolation.
        </li><li>Back-end tested in isolation</li></ul>
        </div>
      </div>
      <!--
      <div>
        <div text-sm font-medium>Cons</div>
        <div text-xs opacity-70>Significant increase in complexity</div>
      </div>
      -->
    </div>
  </div>
</div>
</v-click>

<v-click>

<div border="2 solid black/5" rounded-lg overflow-hidden bg="black/5" backdrop-blur-sm h-full>
  <div flex items-center bg="black/8" backdrop-blur px-3 py-2 rounded-md>
    <div text-amber-300 text-sm mr-2 />
    <div font-semibold>
      2020s - Hypermedia
    </div>
  </div>
  <div px-4 py-3>
    <div flex flex-col gap-3>
        <div text-xs opacity-70>Backend delivers chunks of HTML</div>
      <div flex items-center>
        <div i-carbon:help mr-1/>
        <div text-sm font-medium>Testing Strategy ...</div>
      </div>
        <div text-xs opacity-70>Playwright ???</div>
    </div>
  </div>
</div>
</v-click>

</div>

---
layout: top-title
color: slate
---

:: title ::

# Problem with Browser Automation

:: content ::

  <div flex v-click>
    <div i-carbon:debug mr-1/>
    <div text-xs opacity-70>Slow</div>
  </div>
  <div flex v-click>
    <div i-carbon:debug mr-1/>
    <div text-xs opacity-70>Erratic</div>
  </div>
  <div flex v-click>
    <div i-carbon:debug mr-1/>
    <div text-xs opacity-70>Fragile</div>
  </div>

---
layout: top-title
color: slate
align: l
---

:: title ::

# Verifying Behaviour?

:: content ::

Application behaviour becomes a choreography between

- HTML Elements with specific attributes defined
- The hypermedia framework parsing those attributes
- HTTP Response headers and body on generated responses
- Generated by handlers written for that framework
