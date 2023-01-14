# gocalc
A simple command line expression evaluator written in Go.

## Supported Operators

| Operator   |      Description   | Example Syntax |
|:----------:|:-------------------|:---------------|
| +          | Addition           | 1 + 2          |
| -          | Subtraction        | 1 - 2          |
| *          | Multiplication     | 10 * 20        |
| /          | Division           | 10 / 20        |
| ^          | Power              | 2 ^ 4          |
| ()         | Parentheses        | 2 * (1 + 3)    |

## Variables
gocalc supports variables in the expression.  Variables are accessed via the $<identifier> syntax.

$ans is reserved for the result of the last operation.
