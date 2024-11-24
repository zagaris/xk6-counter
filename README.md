# xk6-counter

A k6 extension that provides a simple, thread-safe global counter. Counter is shared across all Virtual Users (VUs) in a K6 test.

```JavaScript
import counter from 'k6/x/counter';

export default function () {
  counter.increment();

  console.log(`Counter value: ${counter.Get()}`);
}
```

## Build

### 1. Install `xk6`

```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

### 2. Build `K6` binary

```bash
xk6 build --with github.com/zagaris/xk6-counter@latest
```

### 3. Run test
```bash
./k6 run atomic.test.js
```
