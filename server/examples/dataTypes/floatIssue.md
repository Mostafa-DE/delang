### Floating Point: Issues and Limitations

<br />

Floating-point numbers are represented in computer hardware as **base 2** (binary) fractions. Most decimal fractions cannot be represented exactly as binary fractions. A consequence is that, in general, the decimal floating-point numbers you enter are only approximated by the binary floating-point numbers actually stored in the machine.

<br />

The problem is easier to understand at first in **base 10**. Consider the fraction `1 / 3`. You can approximate that as a **base 10** fraction:

`0.3` or `0.33` or `0.333` and so on. No matter how many digits you’re willing to write down, the result will never be exactly `1 / 3` but will be an increasingly better approximation of `1 / 3`.

<br />

In the same way, no matter how many **base 2** digits you’re willing to use, the decimal value `0.1` cannot be represented exactly as a **base 2** fraction. In **base 2**, `1 / 10` is the infinitely repeating fraction
`0.000110011001100110011001100110011001100110011…`

<br />

So what is the solution? Well, there are several ways to handle this issue. One way is to use a **decimal** type. We will discuss it in the next slide.
