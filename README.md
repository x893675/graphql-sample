## graphql-sample

### 依赖库

* [gqlgen](https://github.com/99designs/gqlgen)

### request test

1. create question

```graphql
mutation CreateQuestion {
    CreateQuestion(question: { title: "which language is the fatest?", options: [{
        title: "Go",
        position: 1,
        isCorrect: true
    }, {
        title: "java",
        position: 2,
        isCorrect: false
    }, {
        title: "c#",
        position: 3,
        isCorrect: false
    }]}) {
        message
        status
        data {
            id
            title
        }
    }
}
```

