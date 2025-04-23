# Delve

**Delve** is a personal project to build a custom **assembler** and **C-like programming language**, from the ground up. It's all about diving deep (pun intended) into how code becomes raw machine instructions. Written entirely in Go.

---

## What Is Delve?

Delve is a two-part toolchain for low-level code experimentation:

### 🔧 Delve ASM
A lightweight assembler that converts human-readable assembly into raw machine code, targeting either:
- A real CPU architecture (TBD)
- Or a **custom virtual CPU** with a simplified instruction set

### 🧬 Delve Lang
A minimalist, C-style programming language that compiles to Delve ASM. It includes:
- Static types (`int`, `bool`)
- Arithmetic & logic
- Control flow: `if`, `while`, `return`
- Functions
- A tiny standard library

---

## 🛣️ Roadmap

### ✅ Phase 1: Assembler (`delve-asm`)
- [ ] Design custom ISA
- [ ] Create assembly syntax
- [ ] Implement lexer and parser
- [ ] Translate instructions to binary
- [ ] Add labels and control flow (jumps)
- [ ] Write a bytecode interpreter / VM

### 🛠️ Phase 2: C-like Language (`delve-lang`)
- [ ] Define grammar and syntax
- [ ] Write tokenizer
- [ ] Build AST parser
- [ ] Convert to intermediate representation (IR or Delve ASM)
- [ ] Build interpreter or codegen backend
- [ ] Add basic standard library (`print`, etc.)

### 💡 Future Goals
- [ ] Support `char`, `string`, and arrays
- [ ] Add `for`, `switch`, and other control structures
- [ ] Optimizations and type checking
- [ ] Compile to real-world assembly (x86, RISC-V)
- [ ] Add developer tooling (syntax highlighter, debugger)

---

## 📁 Project Structure

```text
delve/
├── asm/             # Assembler (Delve ASM)
│   ├── lexer.go     # Tokenizer for assembly syntax
│   ├── parser.go    # Parses instructions and labels
│   └── emitter.go   # Converts parsed instructions to binary output
│
├── lang/            # Delve Lang (C-like language)
│   ├── lexer.go     # Tokenizer for source code
│   ├── parser.go    # Builds AST from tokens
│   └── codegen.go   # Converts AST to Delve ASM or IR
│
├── vm/              # Virtual machine (optional)
│   ├── cpu.go       # Core execution engine
│   ├── memory.go    # Memory model
│   └── loader.go    # Loads binary files for execution
│
├── examples/        # Example programs
│   ├── hello.delve     # Delve Lang hello world
│   └── test.asm     # Assembly test program
│
├── internal/        # Shared utilities and types
│   ├── token.go     # Common token definitions
│   └── errors.go    # Error handling utilities
│
├── cmd/             # CLI entry points
│   ├── delve-asm/   # Command to assemble .asm files
│   └── delve-lang/  # Command to compile/run .dl files
│
└── README.md        # You're here 🙂
```

---

## 📚 Inspiration

Delve is inspired by a mix of legendary projects, educational tools, and low-level programming gems:

- **[Crafting Interpreters](https://craftinginterpreters.com/)** by Bob Nystrom  
  A fantastic guide to building a language from the ground up, with real-world advice on parsers, VMs, and more.

- **[Ben Eater's 8-bit computer](https://eater.net/8bit)**  
  A deep dive into the fundamentals of CPUs, memory, and how machine code really works — in hardware.

- **[The Zig Programming Language](https://ziglang.org/)**  
  A modern, low-level language that strips away unnecessary abstractions. Clean design and compile-time power.

- **[LLVM](https://llvm.org/)**  
  The gold standard of modular compiler infrastructure. A masterclass in IR and backend flexibility.

- **[Nand2Tetris](https://www.nand2tetris.org/)**  
  A hands-on course where you literally build a computer and software stack from NAND gates to an OS and compiler.

- **TinyCC / LCC / 8cc / C4**  
  Minimal C compilers that show just how compact a working compiler can be.

- **Forth and Wirth-style simplicity**  
  Languages and VMs that prioritize minimalism, directness, and small footprints.

- The countless indie devs, compiler nerds, and low-level tinkerers sharing their work through blog posts, GitHub repos, and late-night Hacker News threads.

> This project is about learning deeply by building slowly — the best kind of dive.


