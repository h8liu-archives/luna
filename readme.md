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
without a huge amount of help from the developer community. The
development of these projects, on the other hand, is often driven by
bug fixes, feature requests and project deadlines, and the developers
who understand the code well often has no fundamental interests on
code refactoring at this stage.  Therefore, again, I decide to
reinvent the wheel and build something small from scratch. 

Some might say I build it for fun. Though I often enjoy coding, I
would really prefer not doing it this time if I can read the code like
reading a well-written book. I tried reading the source code of a lot
of open source projects, but I feel that I just can't understand them
thoroughly; often I even don't know how it is structrued at a high
level and where to start reading.

In particular for this project, I decide to build something over bare
metal, so that it does not depend on any software layer that I cannot
possibly understand (i.e. operating systems like Linux). Sadly, to
make it (sort of) more real, it still needs to be running on some
existing hardware like Raspberry Pi, which has a pretty complex
hardware interface layer.  I used to consider writing the
project for a much simpler instruction set that I designed myself
(called [E8](http://e8vm.net)), but I think a real working system
might be eventually more attractive than a slow, simulated toy VM. The
good part is, I will always have reference working systems that I can
look into and handy tools to play around with when I am stuck on a
problem.

## Roadmap

I just started the project. Here is my current plan.

I will first specify a strict subset of ARM which provides MIPS-like
functionality (much less powerful but simpler to undertand), and then
build a simple, stupid, slow and straight-forward compiler for a
customized language targeting this intruction subset. When I have the
compiler working, I will start writing a simple OS with this language
using this compiler.

Eventually, my dream is that we have a working small toy operating
system that can compile itself on Raspberry Pi, and every piece of
code in this project will not only be working, but also be reviewed by
a bunch of reviewers that has no understanding on the project design
at first. I will use all kinds of methods (comments, documentation,
tutorials, examples, live-demos, playground, whatever) to help the
reviewers understand the implementation details.  After all, the
purpose of this project is to produce human-comprehensible code,
rather than yet-another compiler/operating system.

If you find this project intersting, you are always welcome to
contribute. Besides programmers, I am also trying to find volunteers
to review the code and provide feedback on if they can understand the
code or not, and what might help to make them understand. Please feel
free to send me [email](mailto:liulonnie@gmail.com).

Thanks.
