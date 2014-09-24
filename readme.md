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

Although everyone can access the code of open source projects like GCC
and Linux, fewer and fewer people understand how it works. Even for
volunteers that contribute to the project for years, they probably
only understand part of it. Worse is, the understanding fades through
time and could be eventually lost when the programmer is gone. New
programmers can sure treat the working code as a black box and use it
via its interface, but without the ability to understand the internal
code logic thoroughly, making fundamental improvements on the code
becomes largely practically impossible. Consequently, as time flies,
wheels have to be reinvented.

In fact, as a computer science student, I reinvented several wheels
myself. I am often blamed by my advisors and collegues for wasting
time on building stuff that already exists. However, I often feel it
is neccessary, because as a researcher, I am required to diagnose and
explain how and why something works or not, and hence have to
understand how it works inside, and reinventing the wheel is often the
fastest among all ways, and maybe the only way, to achieve this
level of understanding. For the blames I received, I blame the ones
that just wrote the code that works but never thoroughly explained how
the code works.

If it is possible to understand and refactor GCC, Linux or even LLVM
in a more readable and understandable fashion, I am more than happy to
do so. However, these projects have all been evolved into huge,
undocumented and largely monolithic (not so clearly modularized)
monsters, and I alone simply will not have the energy to refactor them
without a huge amount of help from the developer community, which
often is driven by feature implementations and project deadline, and
has no fundamental interests on code refactoring at this stage.
Therefore, I decide to reinvent the wheel again, to build something
small from scratch, on my own.

In particular, I decide to build something over bare metal, so that it
does not depend on any software layer that I cannot possibly
understand (i.e. Linux). Sadly, to make it (sort of) more real, it
still needs to be running on some existing hardware like Raspberry Pi.
I used to consider writing for a much simpler instruction set that I
invented myself (called [E8](http://e8vm.net) ), but I think a real
working system might be eventually more attractive than a slow,
simulated toy VM.

## Roadmap

I just started the project. Here is my current plan.

I will first specify a strict subset of ARM which provides MIPS-like
functionality (much less powerful but simpler to undertand), and then
build a simple, stupid, straight-forward, and slow compiler targeting
this intruction subset. When I have the compiler working, I will start
writing a simple OS with this compiler.

Eventually, every piece of code in this project will not only be
working, but also be reviewed by a bunch of reviewers that has no
understanding on the code design. I will use all kinds of methods
(comments, documentation, tutorials, examples, live-demos, playground,
whatever) to help the reviewers to understand the code. After all,
the purpose of this project is to produce human-comprehensible code,
but not yet-another compiler/operating system.

You are always welcome to contribute. Besides programmers, I am also
trying to find volunteers to review the code and provide feedback on
if they can understand the code or not, and what might help to make
them understand. 

Thanks.
