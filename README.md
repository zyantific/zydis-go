<p align="center">
  <a href="https://zydis.re/">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/zyantific/zydis/master/assets/img/logo-dark.svg" width="400px">
      <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/zyantific/zydis/master/assets/img/logo-light.svg" width="400px">
      <img alt="zydis logo" src="https://raw.githubusercontent.com/zyantific/zydis/master/assets/img/logo-dark.svg" width="400px">
    </picture>
  </a>
  <i style="font-size:3em; color: #5ba8ff; font-weight:bold">Go</i>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
  <a href="https://github.com/can1357/zydis-go/actions"><img src="https://github.com/can1357/zydis-go/actions/workflows/build.yml/badge.svg" alt="GitHub Actions"></a>
  <a href="https://gitter.im/zyantific/zydis?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=body_badge"><img src="https://badges.gitter.im/zyantific/zyan-disassembler-engine.svg" alt="Gitter"></a>
  <a href="https://discord.zyantific.com/"><img src="https://img.shields.io/discord/390136917779415060.svg?logo=discord&label=Discord" alt="Discord"></a>
</p>

<p align="center">Fast and lightweight x86/x86-64 disassembler and code generation library.</p>

## Features

- Optimized for high performance, runs almost as fast as native Zydis
- No dependencies on Cgo
- Thread-safe by design
- Very small file-size overhead compared to other common disassembler libraries
- Supports all x86 and x86-64 (AMD64) instructions

## Examples

### Disassembler

The following example program uses Zydis to disassemble a given memory buffer and prints the output to the console.

https://github.com/can1357/zydis-go/blob/229b8bb7bf0f346f253337eede5962f799111bd3/examples/disasm-simple/main.go#L12-L42

The above example program generates the following output:

```asm
007FFFFFFF400000  push rcx
007FFFFFFF400001  lea eax, [rbp-0x01]
007FFFFFFF400004  push rax
007FFFFFFF400005  push [rbp+0x0C]
007FFFFFFF400008  push [rbp+0x08]
007FFFFFFF40000B  call [0x008000007588A5B1]
007FFFFFFF400011  test eax, eax
007FFFFFFF400013  js 0x007FFFFFFF42DB15
```

### Encoder

https://github.com/can1357/zydis-go/blob/229b8bb7bf0f346f253337eede5962f799111bd3/examples/encode-simple/main.go#L11-L32

The above example program generates the following output:

```
48 C7 C0 37 13 00 00
```

### More Examples

More examples can be found in the [examples](./examples/) directory of this repository.

## Build

Simply get the package using `go get`:

```bash
go get -u github.com/can1357/zydis-go
```

If you are not on `Windows AMD64` or `Linux AMD64`, you need to build the Zydis library for your platform and place the shared library nearby your executable. You can find the instructions for building the Zydis library [here](https://github.com/zyantific/zydis#build).

## License

zydis-go is licensed under the MIT license, Zydis's license also applies under the same terms. See [LICENSE.md](./LICENSE.md) for more information.
