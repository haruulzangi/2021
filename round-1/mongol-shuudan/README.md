## Монгол шуудан

http://zipcode.mn/list хуудсыг доорх код ашиглаж crawl хийж data.json үүсгэв.

```js
console.log(
  JSON.stringify(
    Object.fromEntries(
      Array.from(document.getElementsByClassName('list-link'))
        .map((el) =>
          el.innerText
            .split('\n')
            .map((s) => s.trim())
            .filter(Boolean),
        )
        .map(([a, b]) => [b, a]),
    ),
  ),
);
```
