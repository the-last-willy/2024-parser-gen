package json

const mckeemanGrammar = `json
    element

value
    object
    array
    string
    number
    "true"
    "false"
    "null"

object
    '{' members '}'

members
    member ',' members
    member

member
    ws string ws ':' element

array
    '[' ws ']'
    '[' elements ']'

elements
    element ',' elements
    element

element
    ws value ws

string
    '"' characters '"'

characters
    character characters
    ""

character
    '0020' . '10FFFF' - '"' - '\'
    '\' escape

escape
    '"'
    '\'
    '/'
    'b'
    'f'
    'n'
    'r'
    't'
    'u' hex hex hex hex

hex
    digit
    'A' 'F'
    'a' 'f'

number
    integer fraction exponent

integer
    onenine digits
    digit
    '-' onenine digits
    '-' digit

digits
    digit digits
    digit

digit
    '0'
    onenine

onenine
    '1' . '9'

fraction
    '.' digits
    ""

exponent
    'E' sign digits
    'e' sign digits
    ""

sign
    '+'
    '-'
    ""

ws
    '0020' ws
    '000A' ws
    '000D' ws
    '0009' ws
    ""
`
