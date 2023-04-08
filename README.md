# xk6-enhanced

A [k6 extension](https://k6.io/docs/extensions/) adds enhanced JavaScript compatibility (TypeSyript, import JSON files, etc) to [k6](https://k6.io) scripts. 

Use `--compatibility-mode=enhanced` to activate enhanced JavaScript compatibility mode.

```bash
./k6 run --compatibility-mode=enhanced script.js
```

> See [script.ts](script.ts), [script.js](script.js), [user.ts](user.ts) for basic example

## Features

In enhanced compatibility mode the test script will be loaded using [esbuild](https://esbuild.github.io/). Most of the esbuild features will be available in script.

 - TypeScript language support
    ```bash
    ./k6 run --compatibility-mode=enhanced script.ts
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
   - import TypeScript module from JavaScript script/module
   - import JavaScript module from TypeScript script/module

## Limitations

Remote modules loaded by k6, so no enhanced features (TypeScript, etc) not available (yet) for remote modules. Remote module loading support is a planned feature.

## Download

You can download pre-built k6 binaries from [Releases](https://github.com/szkiba/xk6-enhanced/releases/) page. Check [Packages](https://github.com/szkiba/xk6-enhanced/pkgs/container/xk6-enhanced) page for pre-built k6 Docker images.

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download `xk6`:
  ```bash
  $ go install go.k6.io/xk6/cmd/xk6@latest
  ```

2. Build the binary:
  ```bash
  $ xk6 build --with github.com/szkiba/xk6-enhanced@latest
  ```

## Usage

```bash
$ ./k6 run --compatibility-mode=enhanced script.js
```

## Docker

You can also use pre-built k6 image within a Docker container. To do that, you'll need to execute something more-or-less like the following:

**Linux**

```plain
docker run -v $(pwd):/scripts -it --rm ghcr.io/szkiba/xk6-enhanced:latest run --compatibility-mode=enhanced /scripts/script.js
```

**Windows**

```plain
docker run -v %cd%:/scripts -it --rm ghcr.io/szkiba/xk6-enhanced:latest run --compatibility-mode=enhanced /scripts/script.js
```

