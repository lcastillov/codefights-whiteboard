def newStyleFormatting(s):
    s = s.replace('%%', '{%}')
    for t in "bcdeEfFgGnosxX":
        s = s.replace('%' + t, '{}')
    s = s.replace('{%}', '%')
    return s
