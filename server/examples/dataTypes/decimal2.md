### Numeric Types (decimal) (Continued)

<br />

if we want to **control** the **precision** of the digits, let’s say we want to round the number to be 2 digits.

<br />

You can achieve that by changing the **\_getDecimalData** dictionary, which is a **global** object that contains `{'prec': 8, 'divPrec': 8}`, where **prec** is used for rounding any operation except the division and the mod operation, to round that you need to change the **divPrec**.

<br />

**Note:** the allowed range for **prec** is from `1 to 8` and for **divPrec** is from `1 to 28` and attempting to store a value outside that range can lead to run time error.
