---
layout: top-title
color: slate
---

:: title ::

# Two-Step Parsing ##

:: content ::


<v-drag pos="24,82,334,86">

### The HTML

```html
<!doctype html><html><head><template>
    <div></div>
    <div></div>
    <div></div>
</template>
</head><body>
<h1>Heading</h1>
<p>Body</p>
</body></html>
```

</v-drag>

<ArrowDraw color="red" v-drag="[534,364,95,52,150]" />

<v-drag pos="553,67,424,307">

### `x/net/html` Representation

```plantuml
@startuml

skinparam backgroundColor transparent

object "Node" as doc {
    Type: DocumentNode
}
object "Node" as html {
    Type: ElementNode
    Data: "HTML"
}
object "Node" as head {
    Type: ElementNode
    Data: "HEAD"
}
object "Node" as body {
    Type: ElementNode
    Data: "BODY"
}
object "Node" as h1 {
    Type: ElementNode
    Data: "H1"
}
object "Node" as p {
    Type: ElementNode
    Data: "P"
}
object "Node" as template {
    Type: ElementNode
    Data: "TEMPLATE"
}
object "Node" as d1 {
    Type: ElementNode
    Data: "DIV"
}
object "Node" as d2 {
    Type: ElementNode
    Data: "DIV"
}
object "Node" as d3 {
    Type: ElementNode
    Data: "DIV"
}

doc o- html
html o-- head
html o-- body
head o-- template
template o-- d1
template o-- d2
template o-- d3
body -- h1
body -- p

@enduml
```


</v-drag>

<v-drag pos="339,165,72,44">

## DOM

</v-drag>

<v-drag pos="65,212,497,328">


<!-- note "HTMLTemplateElement" doesn't have any children" as n -->
<!-- HTMLTemplateElement .. n -->

```plantuml
@startuml

skinparam backgroundColor transparent

object HTMLDocument
object HTMLHtmlElement
object HTMLBodyElement
object HTMLHeadElement
object "HTMLHeadingElement" as h1 {
    tagName: "H1"
}
object "HTMLParagraphElement" as p
object HTMLTemplateElement
object DocumentFragment
object "HTMLDivElement" as d1
object "HTMLDivElement" as d2
object "HTMLDivElement" as d3

HTMLDocument o- HTMLHtmlElement
HTMLHtmlElement o-- HTMLHeadElement
HTMLHtmlElement o-- HTMLBodyElement
HTMLBodyElement -[hidden]> HTMLHeadElement
HTMLHeadElement o--- HTMLTemplateElement
HTMLBodyElement -- h1
HTMLBodyElement -- p

h1 --[hidden]> DocumentFragment
HTMLTemplateElement ----l--> "content" DocumentFragment
DocumentFragment o-- d1
DocumentFragment o-- d2
DocumentFragment o-- d3

note bottom of HTMLTemplateElement
    HTMLTemplateElement 
    doesn't have children.
end note

@enduml
```

</v-drag>
