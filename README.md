# goa-lang
Ruby inspired syntax that is transpiled to go.
Bringing syntax sugar to improve development speed.

# Supported features
- Basic testable program
  - [X] Main function
  - [X] #Println
  - [ ] Untyped constants
    - [X] Boolean
      - [X] true
      - [X] false
    - [X] String ("")
    - [ ] Integer
      - [ ] Integer unsigned (1)
      - [ ] Integer signed(-1)
    - [ ] Float
      - [ ] Basic Float unsigned (1.1)
      - [ ] Basic Float signed (-1.1)
      - [ ] Integer mantissa + Integer exponent (1e11)
      - [ ] Integer mantissa + Basic Float exponent (1e1.1)
      - [ ] Float mantissa + Integer exponent (1.1e11)
      - [ ] Basic Float mantissa + Basic Float exponent (1.1e1.1)
    - [ ] Rune
      - [ ] 'a'
    - [ ] Imaginary
      - Integer imaginary (1i)
      - Float imaginary (1.1i)
      - Integer imaginary + Integer imaginary (1 + 1i)
      - Integer imaginary + Float imaginary (1 + 1.1i)
      - Float imaginary + Integer imaginary (1.1 + 1i)
      - Float imaginary + Float imaginary (1.1 + 1.1i)
    - [ ] Nil
- Universe block
  - [ ] variable declaration
    - [ ] ...datatypes
  - [ ] builtin functions

# Extended features
- [ ] C like enums