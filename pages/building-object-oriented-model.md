---
layout: top-title
color: slate
---

:: title ::

# Object-Oriented Model - Interfaces

:: content ::

```go
type Node interface {
	event.EventTarget
	AppendChild(node Node) (Node, error)
	ChildNodes() NodeList
	CloneNode(deep bool) Node
	InsertBefore(newNode Node, referenceNode Node) (Node, error)
	OwnerDocument() Document
	ParentNode() Node
    // ...
}

type Element interface {
    Node
	ClassList() DOMTokenList
	HasAttribute(name string) bool
	GetAttribute(name string) (string, bool)
	SetAttribute(name string, value string)
	RemoveAttribute(name string)
	GetAttributeNode(string) Attr
	SetAttributeNode(Attr) (Attr, error)
    // ...
}
```

Interface `Element` includes all methods defined in interface `Nide`

---
layout: top-title-two-cols
color: slate
---

<ArrowDraw color="red" v-drag="[346,254,228,51,-37]" v-click="3"/>
<div style="width:300px" v-drag="[194,326,300,53]" v-click="3">This `node` knows nothing about being embedded in a `element`.</div>

<div v-click="4">Go structs do not resemble model inheritance</div>

:: title ::

# Object-Oriented Model - Structs

:: left ::

````md magic-move {lines: true, maxHeight:'20px'}
```go
type eventTarget struct {
    listeners map[string][]EventListener
}
```

```go
type eventTarget struct {
    listeners map[string][]EventListener
}

type node struct {
    eventTarget
    childNodes []*node
    parent *node
    ownerDocument *document
}
```
````

:: right ::

<div class="struct" style="margin-bottom:1rem">
    <span>EventTarget</span>
    <div class="field"><span class="name">listeners</span>
        map[string][]EventListener
    </div>
</div>


<v-click at="2">
<div class="struct">
    <span>Element</span>
    <div class="struct">
        <span>Node</span>
        <div class="field">
            <span>EventTarget</span>
            <div class="field">...</div>
        </div>
        <div class="field"><span class="name">childNodes</span> []Node</div>
        <div class="field"><span class="name">parent</span> Node</div>
        <div class="field"><span class="name">ownerDocument</span> Document</div>
    </div>
    <div class="field"><span class="name">tagName</span> string</div>
    <div class="field"><span class="name">attributes</span> []Attr</div>
</div>
</v-click>

<style>
.struct {
width: max-content;
font-size: 0.8rem;
border: 1px solid black;
display: flex;
         flex-direction: column;
padding: 0.125rem 0.25rem;
         margin: 0.5rem 0;
         span {
             font-weight: bold;
         }
}

.field {
    border: 1px solid black;
    padding: 0.125rem 0.25rem;
    margin: 0.125rem 0;
    &:last-child {
        /* margin-bottom: 0; */
    }
}
</style>

