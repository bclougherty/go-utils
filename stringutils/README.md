stringutils
===========

    Package stringutils provides some higher-level string operations that
    aren't provided by the strings package.

FUNCTIONS
=========

func CamelToSpinal(input string) string
    CamelToSpinal will convert a string from camelCase or UpperCamelCase
    into spinal-case.

func LcFirst(input string) string
    LcFirst returns a string that matches the input string, but with the
    first letter converted to lowercase.

func SpinalToCamel(input string) string
    SpinalToCamel will convert a string from spinal-case or Train-Case into
    UpperCamelCase

func UcFirst(input string) string
    UcFirst returns a string that matches the input string, but with the
    first letter capitalized.
