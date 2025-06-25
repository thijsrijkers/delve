# Poncho


## What is Poncho?

Poncho is a minimalist programming language interpreter written in Go, crafted to explore and understand the core concepts of systems programming languages such as C, Zig, and Rust. Unlike many beginner-friendly language interpreters that abstract away memory management, Poncho dives into manual memory allocation, pointer manipulation, and raw buffer control. It allows you to explicitly allocate and free memory, work directly with pointers, and manipulate fixed-size buffers — all from a safe and simple Go environment.


## Why Build Poncho?

Most interpreters focus on high-level languages that hide memory details and hardware interactions. However, to truly grasp how low-level languages achieve their performance and control, it is essential to understand how manual memory management and pointers operate under the hood. Poncho serves as a sandbox for learning these systems programming concepts without the complexity of writing a full compiler or dealing directly with hardware. It helps build intuition about how programs manage resources, manipulate memory, and perform pointer arithmetic, giving you a solid foundation for designing your own language with systems-level features.


## What Does Poncho Offer?

- A lexer and parser that transform source code into an executable abstract syntax tree (AST).
- An interpreter that walks the AST and executes your program.
- A simulated manual memory heap using Go slices for allocation and deallocation.
- Pointer semantics allowing referencing variables by address and dereferencing them.
- Buffer support for creating and manipulating contiguous memory blocks.
- Basic language constructs like variable assignment, arithmetic operations, and printing.


## How Does It Work?

When you execute a Poncho program, the lexer first breaks down the source code into tokens. The parser then organizes these tokens into an AST representing the structure and logic of your program. The interpreter evaluates this AST, handling expressions, statements, and memory operations along the way. When your program requests memory allocation (for example, creating a buffer), the interpreter reserves space in a simulated heap represented by a Go slice. Pointers hold addresses within this heap, and dereferencing pointers or accessing buffer elements translates to reading or writing values at those memory addresses. You can also manually free memory blocks to avoid leaks and reuse memory efficiently.


## Example Poncho Code

```Poncho
# Create a new allocator instance
const allocator = new Allocator()

# Allocate 5 bytes of memory using the allocator
const buf = allocator.alloc(5)

# Write value 42 into the buffer at index 0
buf[0] = 42

# Read and print the value from the buffer
print buf[0]    # prints 42

# Free the allocated buffer memory explicitly
allocator.free(buf)

# Example of pointers and dereferencing
let x = 10
let p = &x
let y = *p
print y          # prints 10
```

In this example, the allocator instance explicitly manages memory allocation and freeing. The alloc method reserves memory blocks, returning pointers that can be used for indexed access or pointer operations. The free method releases memory back to the allocator, enabling reuse. This pattern mirrors the manual memory management approach used in systems languages like Zig.


## Why Implement This in Go?

- Go is fast and statically typed, enabling efficient and reliable interpreter execution.
- Its simple syntax and rich standard library simplify the implementation of lexers, parsers, and memory management.
- Go’s slices and maps naturally lend themselves to simulating memory heaps and environments.
- Portability and tooling make distributing and maintaining the interpreter easier.


## Learning and Inspiration

- *Writing an Interpreter in Go* by Thorsten Ball provides a comprehensive guide to language design and interpreter implementation.
- The Monkey programming language is a minimalist language interpreter written in Go and offers valuable insights.
- Systems programming languages like C, Zig, and Rust inspire manual memory management and pointer semantics.
- Go LLVM bindings offer a future path toward compiling your language to native machine code.


## License

This project is licensed under the MIT License
