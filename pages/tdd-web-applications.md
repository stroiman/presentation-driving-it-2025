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

