---
layout: top-title
color: slate
---

<StickyNote title="Performance" v-drag="[129,260,180,180,8]">

The two-step parsing does add some overhead compared to a dedicated parser.
I prioritised progress.

</StickyNote>

<StickyNote title="Compatibility" v-drag="[316,233,210,180,-7]">

Not all browser behaviour would work, e.g. `document.write()`. It's not a
priority to support all features; Focus is to support building <em>modern</em> web
applications.

</StickyNote>

:: title ::

# HTML Parsing

:: content ::

- Package `golang.org/x/net/html` provides an HTML-5 compliant parser
- 

---
src: ./two-step-parsing.md
---

---
layout: top-title
color: slate
---

:: title ::

# Parsing `<script></script>`

:: content ::

Must be executed at "the right time"
