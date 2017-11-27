#!/usr/bin/env python3
def isShuffleRecursive(w, u, v): # returns a boolean
    if len(u) == 0:
        return v == w
    if len(v) == 0:
        return u == w
    u1, v1, w1 = u[:len(u)-1], v[:len(v)-1], w[:len(w)-1]
    return (w[len(w)-1] == v[len(v)-1] and isShuffleRecursive(w1, u, v1)) or (w[len(w)-1] == u[len(u)-1] and isShuffleRecursive(w1, u1, v))

if not isShuffleRecursive("010101", "000", "111"):
    print("w:010101 u:000 v:111 should be true")
if not isShuffleRecursive("001111", "011", "011"):
    print("w:001111 u:011 v:011 should be true")
if isShuffleRecursive("100", "", "10"):
    print("w:100 u: v:10 should be false")
