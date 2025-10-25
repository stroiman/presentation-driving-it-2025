# Building a Headless Browser in Go - Example

This is part of the presentation for Driving IT 2025

## The example

This example starts with a simple web application with three pages

- `/` The index page, visible to all
- `/login` A login page
- `/private` Private data, requires the user to be authenticated.

The example implements the following behaviour

- Requesting a private page "redirects" to the login page.
- Passing invalid credentials will display an error message
- Passing valid credentials will authenticate the user and redirect.
  - When the user was redirected from the `/private` path, redirect there.
  - Otherwise, redirect to the index page.

Redirects, and form posts are not real redirects, but uses HTMX to swap in
content. Tests verify that swapping occurs correctly, and that the browser's URL
is updated correctly.

> [!CAUTION]
> This is not a guide to building session-based authentication. There are many
> security considerations not covered here, e.g. CSRF protection, cookie
> encryption, sensitive data in cookies, and much more.

### The git commit history

The git history has been carefully laid out with each commit showing a step in a
TDD cycle:

- `fail:` A commit that adds a failing test
- `pass:` A commit that makes the test pass
- `pass: incomplete:` The test pass, but the solution is not right
- `refactor:` Refactoring while in the green.

The goal of a `pass:` commit is not always to provide the correct
implementation, but merely to get to a point that the system exhibits the
correct behaviour in the scope of a single test. 

This passing test serves as the feedback loop allowing an efficient
implementation of the correct behaviour.

Two notable examples
- The first marker of a used being logged in is a server-side variable. This works
  well enough in the scope of the test, but obviously is not a secure
  implementation. 
  - The following refactor stores the information in a cookie. In a real
    application, this would go into a dedicated session store.
- When "redirecting" after credentials validate, the possible UI templates are
  hardcoded in the `POST /login` handler.
  - The following refactor calls the root HTTP handler again with the desired
    destination URL. This will activate the `ServeMux` which really has the
    responsibility for rendering the page for a specific request.

## Mocking Authentication

The goal it to use TDD to drive the implementation of _behaviour in the UI
layer_. Validating credentials is a responsibiliy elsewhere.

Therefore, that responsibility is abstracted behind the interface
`Authenticator`. The test injects a stub that allows the test to easily control
the setup.

The dependency tree is naively simple in this small example. For real-world
applications, managing dependencies to support a TDD approach can become
complex. [Surgeon] is an experiment into providing a very simple, yet powerful
solution that allows test code to surgically replace dependencies in a larger
dependency graph.

Surgeon: https://github.com/gost-dom/surgeon


