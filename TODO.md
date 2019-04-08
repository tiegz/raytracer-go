# TODOs

### Work around method chaining problem

### Can we replace the NewX methods with just literals? Is there much constructor logic even needed?

### Can the testing methods just be instance methods on testing.T? e.g. t.assertEqual(e, a)

For example:

```go
  something := thing.DoSomething()
  something = something.DoAnotherThing()
  something = something.DoSomethingElse()
```

Would be nice to enable method chaining here. If we try that now, we get:

`"cannot call pointer method on thing.DoSomething()"`

### Can we remove the setters, since the properties are all public anyway? (e.g. shape.SetMaterial(m) -> shape.Material = m)

### Turn most method receivers into pointers?

- https://golang.org/doc/faq#methods_on_values_or_pointers
- Using value receivers is better for concurrency
