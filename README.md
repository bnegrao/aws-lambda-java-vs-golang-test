# aws-lambda-java-vs-golang-test
Two similar lambdas in java and golang to test cold-start performance

this project has two simple and equivalent lambdas implemented in java and golang that are meant to 
test the cold start performance in both runtimes.

there is a cloudformation template.yaml to deploy both.

In my test the cold start was:

golang: 0.82ms
java: 252ms

In the subsequent invokations, the execution times for both languages were about 0.50ms.
