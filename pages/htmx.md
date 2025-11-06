---
layout: section
color: slate
---

# Hypermedia

An architecture for building web applications providing almost the same level of
interactivity as SPAs, but with the simplicity of SSR.

---
layout: top-title
color: slate
---

<v-click>
<SpeechBubble position="l" color="orange-light" shape="round" class="text-center" v-drag="[198,261,263,64]">
    This is the focus in this talk
</SpeechBubble>
</v-click>

True PWAs are not possible; Hypermedia applications cannot run offline.

:: title ::

# Hypermedia

:: content ::

Offers a greatly simplified approach to building web apps over the SPA
approach; yet still allowing for highly dynamic applications.

## Useful Libraries

- Datastar
- HTMX

---
layout: top-title
color: slate
---

:: title ::

# HTMX Example - Login Page

:: content ::

```html {*|5}{lines:true}
<!doctype html>
<html><head><script src=".../htmx.js"></script></head><body>
<main>
    <h1>Login</h1>
    <form hx-post="/auth/login" hx-swap="innerHTML">
        <input type="hidden" name="csrf-id" value="KCFlT6EoBYShEb4ksLKSD"> 
        <input type="hidden" name="csrf-token" value="1KXux82OP3iGQxiLpbfBO">
        <input type="hidden" name="redirectUrl" value="">
        <div>
            <label for="email">Email</label>
            <input id="email" type="text" name="email" required="" value="" 
                   placeholder="" autofocus="" aria-invalid="false">
        </div>
        <div>
            <label for="password">Password</label>
            <input id="password" type="password" name="password" required="" 
                   value="" placeholder="••••••••" aria-invalid="false">
        </div>
        <button id="submit-login-form-button" type="submit">Sign in</button> 
        <p>Don't have an account yet? <a href="register" hx-boost="true">Click here to register. </a>
        </p>
    </form>
</main>
</body></html>
```

---
layout: top-title
color: slate
---

:: title ::

# HTMX Example - The request

:: content ::

```http
POST /auth/login HTTP/1.1
Host: localhost:7331
HX-Request: true
HX-Current-URL: http://localhost:7331/auth/login
Content-Type: application/x-www-form-urlencoded
Origin: http://localhost:7331
Sec-GPC: 1
Connection: keep-alive
Referer: http://localhost:7331/auth/login
Cookie: auth=...=; rl_session=...; ...

csrf-id=KCFlT6EoBYShEb4ksLKSD
&csrf-token=1KXux82OP3iGQxiLpbfBO
&redirectUrl=
&email=user%40example.com
&password=foobar
```

---
layout: top-title
color: slate
---

:: title ::

# HTMX Example - The Response - failed login

:: content ::

```http {*|6-22|18-20}{lines:true}
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
... more headers

<input type="hidden" name="csrf-id" value="KCFlT6EoBYShEb4ksLKSD"> 
<input type="hidden" name="csrf-token" value="1KXux82OP3iGQxiLpbfBO">
<input type="hidden" name="redirectUrl" value="">
<div>
    <label for="email">Email</label>
    <input id="email" type="text" name="email" required="" value="" placeholder="" autofocus="" aria-invalid="false">
</div>
<div>
    <label for="password">Password</label>
    <input id="password" type="password" name="password" required="" value="" placeholder="••••••••" aria-invalid="false">
</div>
<button id="submit-login-form-button" type="submit">Sign in</button> 
<div id="alert-div" role="alert" aria-live="assertive" class="text-red-700">
    Email or password did not match
</div>
<p class="text-sm font-light text-gray-500 dark:text-gray-400">Don't have an account yet? <a href="register" hx-boost="true">Click here to register. </a>
</p>
```

Remember the attribute `hx-swap="innerHTML"`?

---
layout: top-title
color: slate
---

:: title ::

# HTMX Example - The Response - successful login

:: content ::

```http {*|4|5}{lines:true}
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
Hx-Replace-Url: /private
Hx-Retarget: body
... more headers

<html><head>...</head><body>
    <h1>Private page</h1>
    <p>Welcome, Logged In User, ...</p>
</body></htmx>
```

<v-clicks at="1">

- `Hx-Replace-Url` change the entry in the History API..
- `Hx-Retarget` a CSS selector to override the swap target, i.e., swapping the `<body>` element

</v-clicks>
