# xk6-ts

>[!WARNING]
>This extension is deprecated because its functionality [has been included in k6](https://grafana.com/docs/k6/latest/using-k6/javascript-typescript-compatibility-mode/).

**TypeScript support for k6**

xk6-ts makes TypeScript a first-class citizen in k6.

```sh
k6 run script.ts
```

<details>
<summary>script.ts</summary>

```ts file=examples/script.ts
import { User, newUser } from "./user";

export default () => {
  const user: User = newUser("John");
  console.log(user);
};
```

</details>

<details>
<summary>user.ts</summary>

```ts file=examples/user.ts
interface User {
  name: string;
  id: number;
}

class UserAccount implements User {
  name: string;
  id: number;

  constructor(name: string) {
    this.name = name;
    this.id = Math.floor(Math.random() * Number.MAX_SAFE_INTEGER);
  }
}

function newUser(name: string): User {
  return new UserAccount(name);
}

export { User, newUser };
```

</details>

That's it. A test written in TypeScript can be run directly with k6. No need for Node.js, babel, webpack, bundler, build step, etc.

Do you think modern JavaScript features make TypeScript pointless? xk6-ts can also be used to support modern JavaScript features in k6. 

```sh
k6 run script.js
```

<details>
<summary>script.js</summary>

```ts file=examples/script.js
import { newUser } from "./user";

export default () => {
  const user = newUser("John");
  console.log(user);
};
```

</details>


xk6-ts can be disabled by setting the `XK6_TS` environment variable to `false`.

To ensure that runtime error messages report the correct source code position, sourcemap generation is enabled by default. Otherwise, due to transpilation and bundling, the source code position is meaningless.
Sourcemap generation can be disabled by setting the value of the `XK6_TS_SOURCEMAP` environment variable to `false`.

## Features

 - TypeScript language support
    ```bash
    k6 run script.ts
    ```
 - remote (https) TypeScript/JavaScript module support
    ```js
    import { User } from 'https://example.com/user.ts'
    console.log(new User())
    ```

 - importing JSON files as JavaScript object
    ```js
    import object from './example.json'
    console.log(object)
    ```
 - importing text files as JavaScript string
    ```js
    import string from './example.txt'
    console.log(string)
    ```
 - mix and match JavaScript and TypeScript
   - import TypeScript module from JavaScript module
   - import JavaScript module from TypeScript module

## Download

You can download pre-built k6 binaries from [Releases](https://github.com/szkiba/xk6-ts/releases/) page. Check [Packages](https://github.com/szkiba/xk6-ts/pkgs/container/xk6-ts) page for pre-built k6 Docker images.

## Build

The [xk6](https://github.com/grafana/xk6) build tool can be used to build a k6 that will include xk6-faker extension:

```bash
$ xk6 build --with github.com/szkiba/xk6-ts@latest
```

For more build options and how to use xk6, check out the [xk6 documentation](https://github.com/grafana/xk6).

## How It Works

Under the hood, xk6-ts uses the [esbuild](https://github.com/evanw/esbuild) library for transpiling and bundling. To be precise, xk6-ts uses the [k6pack](https://github.com/grafana/k6pack) library, which is based on esbuild.

Before the test run, transpilation and bundling are done on the fly.

## Compatibility warning

xk6-ts is currently integrated into k6 by modifying the execution of the `k6 run` command. This is a temporary solution, the final integration will be done in a different way. This temporary integration assumes that the last argument of the `k6 run` command line is the name of the script file. That is, contrary to the way the original `k6 run` command line works, xk6-ts does not accept flags after the script file name. By the way, this assumption is not uncommon, many other commands only accept flags before positional arguments. (the original `k6 run` command also accepts flags after the positional argument).

## Benchmarking

Setting the `XK6_TS_BENCHMARK` environment variable to `true` will log the time spent on TypeScript/JavaScript bundling. This time also includes downloading any remote modules.
