---
layout: section
color: slate
---

# Script Engines

---
layout: top-title
color: slate
---

:: title ::

# Script Engines

:: content ::

- Gost-DOM uses V8
  - Based on v8go project
    - Originally created, but unmaintained by github user [rogchap]
    - Forked and maintained by github user [tommie]
  - Uses a private fork
    - Support essential v8 features not exposed in v8go
    - Some changes merged to tommie's fork.

[rogchap]: https://github.com/rogchap
[tommie]: https://github.com/tommie

---
layout: top-title
color: slate
---

:: title ::

# Script Engine - JavaScript binding code

:: content ::

<div class="ns-c-tight">


Expose native functions to JavaScript code.

- Wrapper Code decoupled to v8
  - Overly complex API
  - Support alternate script engines (e.g., pure Go)
    - An experimental solution using Goja already exists.
- Support non-essential APIs as separate modules. E.g.
  - Passkey
  - Geo location
</div>

---
layout: top-title
color: slate
---

:: title ::

# Script Engine - Registering global constructors

:: content ::

```go
func Bootstrap[T any](reg js.ClassBuilder[T]) {
	js.RegisterClass(reg, "AbortController", "", NewAbortController)
	js.RegisterClass(reg, "AbortSignal", "EventTarget", NewAbortSignal)
	js.RegisterClass(reg, "Attr", "Node", NewAttr)
	js.RegisterClass(reg, "CharacterData", "Node", NewCharacterData)
	js.RegisterClass(reg, "ChildNode", "", NewChildNode)
	js.RegisterClass(reg, "DOMTokenList", "", NewDOMTokenList)
	js.RegisterClass(reg, "Document", "Node", NewDocument)
	js.RegisterClass(reg, "DocumentFragment", "Node", NewDocumentFragment)
	js.RegisterClass(reg, "Element", "Node", NewElement)
	js.RegisterClass(reg, "Event", "", NewEvent)
	js.RegisterClass(reg, "EventTarget", "", NewEventTarget)
	js.RegisterClass(reg, "HTMLCollection", "", NewHTMLCollection)
	js.RegisterClass(reg, "MutationObserver", "", NewMutationObserver)
	js.RegisterClass(reg, "MutationRecord", "", NewMutationRecord)
	js.RegisterClass(reg, "NamedNodeMap", "", NewNamedNodeMap)
	js.RegisterClass(reg, "Node", "EventTarget", NewNode)
	js.RegisterClass(reg, "NodeList", "", NewNodeList)
	js.RegisterClass(reg, "NonDocumentTypeChildNode", "", NewNonDocumentTypeChildNode)
	js.RegisterClass(reg, "ParentNode", "", NewParentNode)
	js.RegisterClass(reg, "Text", "CharacterData", NewText)
}
```

---
layout: top-title
color: slate
---

:: title ::

# Script Engine - Prototype Operations and Attributes

:: content ::


```go {*|10-13|16-24|26-36}{maxHeight: '26rem',lines:true}
func NewHTMLHyperlinkElementUtils[T any](scriptHost js.ScriptEngine[T]) *HTMLHyperlinkElementUtils[T] {
	return &HTMLHyperlinkElementUtils[T]{}
}

func (wrapper HTMLHyperlinkElementUtils[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLHyperlinkElementUtils[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("href", w.href, w.setHref)
	jsClass.CreatePrototypeMethod("toString", w.href)
	jsClass.CreatePrototypeAttribute("origin", w.origin, nil)
    // ...
}

func (w HTMLHyperlinkElementUtils[T]) href(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.href")
	instance, err := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLHyperlinkElementUtils[T]) setHref(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: HTMLHyperlinkElementUtils.setHref")
	instance, err0 := js.As[html.HTMLHyperlinkElementUtils](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetHref(val)
	return nil, nil
}
```


---
layout: top-title
color: slate
---

:: title ::

# Script Engine - Contributing to v8go

:: content ::

<div class="ns-c-tight">

v8go only expose a subset of v8 functionality. Some essential features were
lacking

- Already merged to [tommie/v8go]
    - Directing `console` output to Go.
    - Setting up inheritance hierarchies
      - `FunctionTemplate::Inherit`
      - `FunctionTemplate::InstanceTemplate`
      - `FunctionTemplate::PrototypeTemplate`
      - `ObjectTemplate::SetAccessorProperty`
- Working, but unmerged to v8go
    - ESM support
    - Named and indexed property handlers
    - `Object::GetPrototype`/`Object::SetPrototype`

</div>

[tommie/v8go]: https://github.com/tommie/v8go

---
layout: top-title
color: slate
---

:: title ::

# Script Engine - Garbage Collection

:: content ::

- Go and JavaScript are two GC languages with a C/C++ bridge
- A Go object can keep a JS object alive, and a JS object can keep a Go object
  alive.
- Gost-DOM _will_ leak for a long-running browser context.

Not an issue for the intended use case.
