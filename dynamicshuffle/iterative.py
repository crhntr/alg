#!/usr/bin/env python
def isShuffleIterative(w, u, v): # returns a boolean
    if len(w) != (len(u) + len(v)):
        return False
    if len(u) == 0:
        return v == w
    if len(v) == 0:
        return u == w
    table = [[False for j in range(len(v)+1)] for i in range((len(u)+1))]
    table[0][0] = True
    for i in range(len(u)):
        table[i+1][0] = table[i][0] and (w[i] == u[i])
    for j in range(len(v)):
        table[0][j+1] = table[0][j] and (w[j] == v[j])
    for i in range(len(u)):
        for j in range(len(v)):
            table[i+1][j+1] = (w[i+j+1] == u[i] and table[i][j+1]) or (w[i+j+1] == v[j] and table[i+1][j])
    return table[len(u)][len(v)]

if not isShuffleIterative("010101", "000", "111"):
    print("w:010101 u:000 v:111 should be true")
if not isShuffleIterative("001111", "011", "011"):
    print("w:001111 u:011 v:011 should be true")
if isShuffleIterative("100", "", "10"):
    print("w:100 u: v:10 should be false")
