---
layout: top-title
---

Priorities have been supporting basic HTMX applications. 

:: title ::

# Features

:: content ::

- Implements most relevant parts of the core browser APIs
  - DOM/HTML DOM
  - XMLHttpRequest
  - url
  - UI Events
- Executes JavaScript using V8
  - JavaScript engine is pluggable.
  - Goja support experimental - pure Go JavaScript engine
- 100% deterministic code execution
  - No erratic tests
- Time Travel and "wait for all"
  - Testing throttling and debounce without test delays
  - Verifying that an action _doesn't_ produce an effect.
