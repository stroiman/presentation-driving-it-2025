---
layout: top-title
color: slate
---

<StickyNote title="Performance" v-drag="[344,348,180,180,8]">

The two-step parsing does add some overhead compared to a dedicated parser.
I prioritised progress.

</StickyNote>

<StickyNote title="Compatibility" v-drag="[674,308,210,180,-7]">

Not all browser behaviour would work, e.g. `document.write()`. It's not a
priority to support all features; Focus is to support building <em>modern</em> web
applications.

</StickyNote>

:: title ::

# HTML Parsing

:: content ::

- HTML parsing is not trivial
- Package `golang.org/x/net/html` provides an HTML-5 compliant parser
  - The model is only a structural representation without behaviour.
  - The DOM is not a perfact mirror of the HTML
  - Certain elements trigger certain behaviour when inserted (e.g., scripts)
- Gost-DOM uses `x/net/html`, and then traverse that structure.

---
src: ./two-step-parsing.md
---

---
layout: top-title
color: slate
---

:: title ::

# Parsing - Script Execution

:: content ::

Scripts must run when mounted in the DOM tree.

```html
<!doctype html>
<html><body>
    <h1>Script Page</h1>
    <p>Visible to script</p>
    <script>
        console.log(document.body.outerHTML)
    </script>
    <p>Not visible to script></p>
</body></html>
```

Logs:

```
<body>
    <h1>Script Page</h1>
    <p>Visible to script</p>
    <script>
      console.log(document.body.outerHTML);
    </script></body>
```

Neither the 2nd paragraph, nor whitespace is yet in the DOM.
