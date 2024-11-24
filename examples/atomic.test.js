import counter from "k6/x/counter";

const NUMBER_OF_VUS = 50;
const NUMBER_OF_ITERATIONS = 30;

export const options = {
  scenarios: {
    contacts: {
      executor: 'per-vu-iterations',
      vus: NUMBER_OF_VUS,
      iterations: NUMBER_OF_ITERATIONS,
      maxDuration: '30s',
    },
  },
};

export default function () {
  counter.increment();
}

export function teardown() {
  const totalCount = counter.get();

  const expectedValue = NUMBER_OF_VUS * NUMBER_OF_ITERATIONS;

  if (totalCount !== expectedValue) {
    throw new Error(
      `Atomic test failed: expected ${expectedValue}, but got ${totalCount}`
    );
  }
  
  console.log(`Atomic test passed: final value = ${totalCount}`);
}
