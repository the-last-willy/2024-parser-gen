package mckeeman

const GrammarSource = `grammar
    rules

space
    '0020'

newline
    '000A'

name
    letter name
    letter

letter
    'a' . 'z'
    'A' . 'Z'
    '_'

indentation
    space space space space

rules
    rule newline rules
    rule

rule
    name newline alternatives
    nothing

nothing
    indentation '"' '"' newline
    ""

alternatives
    alternative alternatives
    alternative

alternative
    indentation items newline

items
    item space items
    item

item
    literal
    name

literal
    range exclude
    singleton
    '"' characters '"'

singleton
    ''' codepoint '''

codepoint
    hexcode
    '0020' . '10FFFF'

hexcode
    "10" hex hex hex hex
    hex hex hex hex hex
    hex hex hex hex

hex
    '0' . '9'
    'A' . 'F'

range
    singleton space '.' space singleton

exclude
    space '-' space singleton exclude
    space '-' space range exclude
    ""

characters
    character characters
    character
    ""

character
    '0020' . '10FFFF' - '"'
`
