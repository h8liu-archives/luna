# Luna

Luna is a coding project that aims for human-comprehensible yet
working code.

In particular, the project plans to implement:

- A simulator (written in golang) that simulates Raspberry Pi.
- A compiler (written in golang) that (cross-)compiles an
  understandable language to ARM.
- A small operating system that is written in this understandable
  language.

## Why?

Although everyone can access the code that implements open source
projects like GCC and Linux. Fewer and fewer people understands how it
works. Even for volunteers that contribute to the project for years,
they probably only understand part of it. Worse, the understanding
fades through time and left when the programmer is gone. New
programmers have to still treat the working code as a black box and
only work with its interface, but without the ability to understand
the code logic thoroughly, making fundamental improvements on the code
becomes practically impossible. As a result, the wheels have to be
reinvented.

As a computer science student, I reinvented a lot wheels. I am often
blamed by my advisors and collegues for wasting time on building stuff
that already exists. However, I still feel it is neccessary because
as a researcher that is required to diagnose and explain how and why
something works or not, I have to understand how it works internally,
and reinventing the wheel is often the fastest among all, and maybe
the only way to achieve this understanding. For the blaming I took, I
blame the ones that just write the code but never thoroughly explain
how the code works.

If it is possible to understand and refactor GCC, Linux or even LLVM
in a more readable and understandable fashion, I am more than happy to
do so. However, these projects have all been evolved into huge and
largely monolithic (not so clearly modularized) monsters, and I alone
simply will not have the energy to refactor them without huge amount
of help from the developer community. Therefore, I decide to, again,
build something from scratch.

To do that, I decide to build something over bare metal, so that it
does not depend on any software layer that I cannot possibly
understand (i.e. Linux). Sadly, to make it (sort of) more real, it
still needs to be running on some existing hardware like Raspberry Pi.
I used to consider to write this on a much simpler instruction set I
invented (called [E8](http://e8vm.net) ), but I think a real working
system might be eventually more attractive than a slow simulated toy
VM.

## Roadmap

I just started the project.

I will first specify a strict subset of ARM which provides MIPS-like
functionality (much less powerful but simpler to undertand), and then
build a simple stupid compiler targeting this intruction subset. When
I have the compiler working, I will start writing a simple OS with
this compiler.

You are always welcome to contribute. Besides programmers, I am also
trying to find volunteers to review my code and tell me if they can
understand the code or not, and what I might help to make them
understand. 

Thanks.
