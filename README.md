# goa-lang
Ruby inspired syntax that is transpiled to go.
Bringing syntax sugar to improve development speed.

## Usage

## Supported features
- Basic testable program
  - [X] Main function
  - [X] #Println
  - [ ] Untyped constants
    - [X] Boolean
      - [X] true
      - [X] false
    - [X] String ("hello")
    - [X] Integer
      - [X] Integer unsigned (1)
      - [X] Integer signed(-1)
    - [X] Nil
      - [] Needs pointer variable declaration with assignment to be tested
- Universe block
  - [ ] variable declaration
    - [X] bool
      - [X] bool
      - [X] *bool
    - [X] string
      - [X] string
      - [X] *string
    - [ ] int
      - [ ] int
      - [ ] *int
      - ... other sizes
    - [ ] uint
      - [ ] uint
      - [ ] *uint
      - ... other sizes
    - ... other datatypes
  - [ ] builtin functions
    - [ ] ... to be defined

## Future features
- [ ] Interfaces
- [ ] Untyped constants
  - [ ] Float
    - [ ] Basic Float unsigned (1.1)
    - [ ] Basic Float signed (-1.1)
    - [ ] Integer mantissa + Integer exponent (1e11)
    - [ ] Integer mantissa + Basic Float exponent (1e1.1)
    - [ ] Float mantissa + Integer exponent (1.1e11)
    - [ ] Basic Float mantissa + Basic Float exponent (1.1e1.1)
  - [ ] Imaginary
    - Integer imaginary (1i)
    - Float imaginary (1.1i)
    - Integer imaginary + Integer imaginary (1 + 1i)
    - Integer imaginary + Float imaginary (1 + 1.1i)
    - Float imaginary + Integer imaginary (1.1 + 1i)
    - Float imaginary + Float imaginary (1.1 + 1.1i)


## Extended features
- [ ] C like enums
- [ ] untyped pointer constant