---
title: Building a Headless Browser in Go
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

drawings:
  persist: true

theme: neversink
---

# Building a Headless Browser in Go

Or: How I Learned To Stop Worrying And Love HTML

by **Peter Strøiman**

::note::


---
layout: top-title
color: slate
---

- [github.com/gost-dom](https://github.com/gost-dom)
- [github.com/stroiman](https://github.com/stroiman)
- [linkedin.com/in/stroiman](https://linkedin.com/in/stroiman)
- [stroiman.photography](https://stroiman.photography)

::title::

# About me

::content::

- Coding since the age of 8
    - Professionally since 1997
- Passionate about Test-Driven Development
- Fine Arts Photographer

---
layout: top-title
color: slate
---

:: title ::

# First, a Revelation

:: content ::

Learning about HTMX, I relalized how 90% of all front-end code I had written the
last 10 years could have been written much, much simpler.

An idea for a web application, and I wanted to use Go for the backend.

Eager to apply TDD, I looked at how people would test  at the landscape and found ...

---
layout: top-title
color: slate
---

:: title ::

# Playwright :(

:: content ::

- Better than Selenium/WebDriver, but still browser automation, with all the
  flaws:

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

<v-clicks>

- Written **after** the code
- They did not provide a feedback loop to help implement functionality

</v-clicks>


---
src: ./pages/test-driven-development.md
---

---
src: ./pages/htmx.md
---

---
src: ./pages/htmxtdd.md
---

---
src: ./pages/example.md
---

---
src: ./pages/building.md
---

<!-- --- -->
<!-- src: ./pages/features.md -->
<!-- --- -->

---
src: ./pages/outtro.md
---
