# Running Tests

To run all SSI related tests:

```
go test -v -count=1 "./x/ssi/tests" 
```

If you want to run a single unit test of a particular test file (say `TestVerificationMethodRotation()`):

```
go test -v -count=1 -run "^TestVerificationMethodRotation$" "./x/ssi/tests"
```

To run multiple tests:

```
go test -v -count=1 -run ^(TestDidResolve|TestDidParam)$ "./x/ssi/tests"
```
