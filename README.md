# Roll

Roll is a command line tool that makes it easy to roll dice. It was
built as a fun little toy that I sometimes use if I have forgotten my
dice or I'm just feeling lazy.

## Usage

To roll a dice simply pass an argument to the tool in the format of
`dn` where `n` is the number of sides of the dice. For example to roll
a 20 sided dice you would pass in `d20`. You can pass in multiple
arguments and the results for each dice roll will be printed, each on
their own line.

### Advanced Syntax

- **Summation** - You can add a number in front of the `d` to have the
	dice rolled multiple times. The total sum of each roll will be
	returned.

- **Modifier** - You can add a modifier to each roll appending a `+n`
	to the argument.

### Examples

```bash
# Roll a d20
roll d20

# Roll a d20 and a d4
roll d20 d4

# Roll 3 d4 and return the total rolled
roll 3d4

# Roll 3 d4, adding 1 to each roll, and returning the total
roll 3d4+1

# Roll a d20 adding 7 and 3 d4 adding 1
roll d20+7 3d4+1
```
