```
asm "test" {
    func main {
        var spbase = 0xc000
        var value = 0x37215478

        sp = spbase
        r0 = value
        call printhex

        .loop
        goto .loop
    }

    func fabonaci {
        // r0: the argument
        if r0 == 0 .ret1
        if r0 == 1 .ret1

        [sp] = lr   // save lr

        [sp+4] = r0 // save r0
        r0 = r0 - 1
        sp = sp + 12
        call fabonaci
        sp = sp - 12
        [sp+8] = r0 // save the return value for later summing

        r0 = [sp+4]
        r0 = r0 - 2
        sp = sp + 12
        call fabonaci
        sp = sp - 12
        r1 = [sp+8] // load the saved sum
        r0 = r0 + r1
        lr = [sp]
        ret

        .ret1
        movi r0, 1
        ret
    }
    
    func putc {
        // todo:
    }

    func printhex {
        [sp] = lr
        movi r4, 8

        .loop
        r4 = r4 - 1
        r5 = r4 << 2
        r1 = r0 >> r5
        r1 = r1 & 0xf
        
        if r1 > 10 .big

        // 0-9
        r1 = r1 + '0'
        goto .print

        // a-z
        .big
        r1 = r1 - 10
        r1 = r1 + 'A'

        .print
        [sp+4] = r0
        [sp+8] = r4
        r0 = r1
        call putc
        r4 = [sp+8]
        r0 = [sp+4]

        if r4 != 0 .loop
    }

    func printnum {
        if r0 < 0 .neg
        ret

        .neg
        [sp] = lr
        [sp+4] = r0 // save r0

        // putc('-')
        sp = sp + 8
        r0 = '-'
        call putc
        sp = sp - 8

        r0 = [sp+4]
        r4 = 0
        r0 = r4 - r0
        sp = sp + 8
        call printnum
        sp = sp - 8
        lr = [sp]
        ret
    }
}
```
