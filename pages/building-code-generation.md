---
layout: section
color: slate
---

# Code Generation

Auto-generate trivial code, and ensure conformance to specs.

---
layout: top-title
color: slate
---

:: title ::

# Code Generation - Web APIs

:: content ::

All operations and attributes are defined in web IDL specifications.

<div class="grid w-full h-fit grid-cols-2 grid-rows-1 mt-10 mb-auto gap-4">
<div class="grid-item grid-col-span-1">

```
interface HTMLAnchorElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions, Reflect] attribute DOMString target;
  [CEReactions, Reflect] attribute DOMString download;
  [CEReactions, Reflect] attribute USVString ping;
  [CEReactions, Reflect] attribute DOMString rel;
  // ...

  // also has obsolete members
};
HTMLAnchorElement includes HTMLHyperlinkElementUtils;
```

</div>
<div class="grid-item grid-col-span-1">


```
interface mixin HTMLHyperlinkElementUtils {
  [CEReactions, ReflectSetter] stringifier 
      attribute USVString href;
  readonly attribute USVString origin;
  [CEReactions] attribute USVString protocol;
  [CEReactions] attribute USVString username;
  [CEReactions] attribute USVString password;
  [CEReactions] attribute USVString host;
  // ...
};
```

</div>
</div>



---
layout: top-title-two-cols
color: slate
---

:: title ::

# Code Generation

:: left ::

Generate code based on web IDL and other specs to avoid writing trivial code,
and ensure conformance to specs.

- JavaScript bindings
- UI Events
- Mapping element tag names to IDL interfaces.
- Generating Go interfaces

[github.com/w3c/webref] provides a curated collection of web IDL and other
relevant specifications in JSON format.

<div v-drag="[507,80,447,420]">

``` {*|15-16}{lines:true}
interface HTMLAnchorElement : HTMLElement {
  [HTMLConstructor] constructor();

  [CEReactions, Reflect] attribute DOMString target;
  [CEReactions, Reflect] attribute DOMString download;
  [CEReactions, Reflect] attribute USVString ping;
  [CEReactions, Reflect] attribute DOMString rel;
  // ...

  // also has obsolete members
};
HTMLAnchorElement includes HTMLHyperlinkElementUtils;

interface mixin HTMLHyperlinkElementUtils {
  [CEReactions, ReflectSetter] stringifier 
      attribute USVString href;
  readonly attribute USVString origin;
  [CEReactions] attribute USVString protocol;
  [CEReactions] attribute USVString username;
  [CEReactions] attribute USVString password;
  [CEReactions] attribute USVString host;
  // ...
};
```

</div>

[github.com/w3c/webref]: https://github.com/w3c/webref

---
layout: top-title
color: slate
---

:: title ::

# Code Generation - JavaScript Bindings

:: content ::

<StickyNote v-drag="[606,372,270,145]" v-click="3">

Plenty of trivial, repetitive code
- encoding, decoding and validating JS values.
- a wrapper function for each DOM function

</StickyNote>

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

<!--

By far the most complex of the different code generators.

-->


---
layout: top-title
color: slate
---

:: title ::

# Code Generation - UI Events

:: content ::

Behaviour of browser events is affected by the  `bubbles` and `cancelable` attributes.

The correct values for each UI event is available from `w3c/webref`

```go
func Click(e dom.Element) bool {
	data := PointerEventInit{}
	event := &event.Event{Type: "click", Data: data}
	event.Bubbles = true
	event.Cancelable = true
	return e.DispatchEvent(event)
}
```

 - `focus`/`blur` events bubble, `focusin`/`focusout` do not.
 - `click` events are cancelable, focus events do not.

---
layout: top-title
color: slate
---

:: title ::

# Code Generation - Tag name To IDL Interface

:: content ::

Ensures that in JavaScript, each element is represented by the correct "class".

```javascript
const element = document.createElement("h6")
Object.getPrototypeOf(element).constructor.name
// "HTMLHeadingElement" 
```

Based on mapping from tag name to web IDL interface.

```go {*|2-3|5-7|}{lines:true}
var HtmlElements = map[string]string {
	"html": "HTMLHtmlElement",
	"head": "HTMLHeadElement",
	// ...
	"article": "HTMLElement",
	"section": "HTMLElement",
	"nav": "HTMLElement",
	// ...
	"form": "HTMLFormElement",
	"label": "HTMLLabelElement",
	"input": "HTMLInputElement",
	// ...
}
```

<StickyNote v-drag="[526,225,276,93,11]" v-click="1">

`HTMLHtmlElement`, `HTMLHeadElement`, and many more only exist in JS, have no Go
implementation.

</StickyNote>

<StickyNote v-drag="[350,320,256,72]" v-click="2">

Many tags don't have a specific web IDL interface associated.

</StickyNote>

<StickyNote v-drag="[507,405,243,98,-6]" v-click="3">

Some elements have corresponding Go types implementing browser behaviour, e.g.,
form handling.

</StickyNote>

---
layout: top-title
color: slate
---

:: title ::

# Code Generation - DOM interfaces

:: content ::

Ensure the implementation of behaviour conforms to specifications.

```go
// This file is generated. Do not edit.

package dom

type ParentNode interface {
	Children() HTMLCollection
	FirstElementChild() Element
	LastElementChild() Element
	ChildElementCount() int
	Prepend(...Node) error
	Append(...Node) error
	ReplaceChildren(...Node) error
	QuerySelector(string) (Element, error)
	QuerySelectorAll(string) (NodeList, error)
}
```

Not used a lot, as this type of code gen was added later.

---
layout: top-title
color: slate
---

:: title ::

# Code Generation - Limitations

:: content ::

Web IDL doesn't provide all information

<div class="ns-c-tight">

- If an operation can throw an `Error`
- Default value for optional arguments

</div>

Web IDL provides defaults, which can be adjusted

```go
var htmlRules = SpecRules{
	"DOMStringMap": {OutputType: OutputTypeStruct},
	"Location": {Operations: OperationRules{
		"assign":  {HasError: true},
	}},
	"History": {
		Operations: OperationRules{
			"pushState": {HasError: true, Arguments: ArgumentRules{
				"data":   {Type: idl.Type{Name: "HistoryState"}},
				"unused": {Ignore: true},
				"url":    {ZeroAsDefault: true},
			}},
		},
    },
}
```

<ArrowDraw v-drag="[371,107,140,52,172]" />

<ArrowDraw v-drag="[409,423,140,52,196]" />

<ArrowDraw v-drag="[301,359,140,40,156]" />

<StickyNote v-drag="[511,72,186,59]">
    Requested by the Rust community as well.
</StickyNote>

<StickyNote v-drag="[429,273,178,85]">
    I decided to not have an unused argument in the Go implementation.
</StickyNote>

<StickyNote v-drag="[555,413,266,96]">

Go doesn't have optional values. A strategy must be chosed for when no value is
provided by JS code.. Here, the _zero value_, i.e. `""` will be used.

</StickyNote>
