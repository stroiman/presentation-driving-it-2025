---
layout: top-title
color: slate
---

:: title ::

# Building

:: content ::

- HTML Parsing
- An Object-Oriented API
- Integrating V8
- Learnings
- Mistakes

---
layout: top-title
color: slate
---

<StickyNote title="Performance" v-drag="[138,334,180,180,8]">
    The two-step parsing does add some overhead compared to a dedicated parser.
    I prioritised progress.
</StickyNote>

<StickyNote title="Compatibility" v-drag="[313,302,180,180,-7]">
    Not all browser behaviour would work, e.g. `document.write()`. It's not a
    priority to support all features; Focus is to support building <em>modern</em> web
    applications.
</StickyNote>

:: title ::

# HTML Parsing

:: content ::

- Package `golang.org/x/net/html` provides an HTML-5 has a parser

---
src: ./two-step-parsing.md
---

foo

---
layout: top-title
color: slate
---

:: title ::

# Parsing `<script></script>`

:: content ::

Must be executed at "the right time"
