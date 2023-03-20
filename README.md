# gocalc
A simple command line expression evaluator written in Go.

## Operators
gocalc supports the following operators in the expression.

| Operator   |      Description   | Example Syntax    |
|:----------:|:-------------------|:------------------|
| +          | Addition           | 1 + 2             |
| -          | Subtraction        | 1 - 2             |
| *          | Multiplication     | 10 * 20           |
| /          | Division           | 10 / 20           |
| ^          | Power              | 2 ^ 4             |
| ( )        | Parentheses        | 2 * (1 + (3 / 4)) |

## Variables
gocalc supports variables in the expression.  Variables are accessed via the $identifier syntax.  $ans is reserved for the result of the last operation.

## Input base
gocalc support the following binary (base 2), octal (base 8), decimal/real (base 10) and hexadecimal (base 16).

| Syntax                |      Description              | Example Syntax |
|:----------------------|:------------------------------|:------------   |
| b$number              | Input binary value            | b$01010101     |
| o$number              | Input octal value             | o$777          |
| number                | Input deciaml value (default) | 1234.5678      |
| h$number              | Input hexadecimal number      | h$1234ABCD     |

## Output Format
To set the output format type the following commands followed by ENTER.

| Command          |      Description                                                        | Example Syntax |
|:-----------------|:------------------------------------------------------------------------|:---------------|
| fix *precision*  | Fix point output mode with optional precision (default precision is 2)  | fix 4          |
| sci *precision*  | Scientific output mode with optional precision (default precision is 2) | sci 3          |
| real *precision* | Real output mode with optional precision (default precision is -1)      | real 5         |
| bin              | Binary output mode                                                      | bin            |
| oct              | Octal output mode                                                       | oct            |
| hex              | Hexadecimal output mode                                                 | hex            |
