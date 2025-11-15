
# GO-LOAN-SIM

Simulate loan billing engine with **Golang**.

Currently the storage I used are **in-memory** storage, but the abstraction is already there so if we want to apply it to another type of storage like any **SQL DB** for example, we could just add the implementation **without breaking** any functionality.

I run this in **simulated** environment so we can test the logic so we can actually **jump the time** to the future.

## How to run

To run this project, execute:

```bash
  git clone https://github.com/fajryhamzah/go-loan-sim
```
Install the dependencies
```bash
  go mod tidy
```

for direct run, you could just
```bash
  go run .
```

or if you want to build it first
for direct run, you could just
```bash
  go build -o bin/loan_sim
```
and run the binary afterward
```bash
  ./bin/loan_sim
```


## Demo

Main Demo of the apps
![DEMO](./assets/main.gif)


Simple demo with lower weekly loans so we can check the flow end to end
![DEMO](./assets/e2e.gif)