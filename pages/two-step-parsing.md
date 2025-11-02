---
layout: top-title-two-cols
color: slate
---

:: title ::

## Two-Step Parsing ##

:: left ::

## The HTML

<v-drag pos="62,126,334,86">
```html
<!doctype html><html><head>...</head><body>
...
</body></html>
```
</v-drag>

<ArrowDraw color="red" v-drag="[115,207,95,52,47]" />

<v-drag pos="96,162,270,346">
```plantuml
@startuml

skinparam backgroundColor transparent

object html
object head
object body
object h1
object p
object template
object "div" as d1
object "div" as d2
object "div" as d3

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

:: right ::

## The DOM

<v-drag pos="434,142,222,249">
```plantuml
@startuml

skinparam backgroundColor transparent

object html
object head
object body
object h1
object p
object template

html o-- head
html o-- body
head o-- template
body -- h1
body -- p
@enduml
```
</v-drag>

<v-drag pos="698,320,198,153">
```plantuml
@startuml

skinparam backgroundColor transparent

object HTMLTemplateElement
object DocumentFragment

HTMLTemplateElement -- "content" DocumentFragment
```
</v-drag>

---
src: ./building-object-oriented-model.md
---
