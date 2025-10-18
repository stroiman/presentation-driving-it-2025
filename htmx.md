## HTMX

```html {*|5}{lines:true}
<!doctype html>
<html><body>
<main>
    <h1>Login</h1>
    <form hx-post="/auth/login" hx-swap="innerHTML">
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
        <p>Don't have an account yet? <a href="register" hx-boost="true">Click here to register. </a>
        </p>
    </form>
</main>
</body>
```

Headers
```
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

# Response

```html {*|18}{lines:true}
HTTP/1.1 200 OK
Cache-Control: no-cache
Content-Type: text/plain; charset=utf-8
...

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
<div id="alert-div" role="alert" aria-live="assertive" class="text-red-700">Email or password did not match</div>
<p class="text-sm font-light text-gray-500 dark:text-gray-400">Don't have an account yet? <a href="register" hx-boost="true">Click here to register. </a>
</p>
```

