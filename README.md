Gobots
======

Small go applications which leverage 0mq patterns for c&c and publishing.


Design will eventually include a dynamic registry for registered bots, for c&c,
configuration and 0mq implementation.

Bots
====

**Parrot**
----------
Chirps out github compare diff urls when web post-commit hooks are 
configured on github.com for a repo..


**Burt**
--------
Basic URl Transormation, should of called this burlr for Basic 
URl Rasterizer. Currently this is a work in progress. It loads a GTK Event loop 
and opens Web Browser windows using the super sweet GTK-Webkit libraries from 
Matz.  I did my best to combine both reactor patterns, its tricky stuff..

* not thread safe 
* Don't enable more than 1 worker!


**WebVu**
---------
First foray into bots and using CGo.. Matz source code really 
opened a few doors. Its astounding that the barrier to entry for C integration
is so low. 

* not thread safe
* Don't enable more than1 worker!


Wish List
=========

Expose functionality using gopy to research the possibilities of using native
CSP routines in a python environment. For idempotent code determine if a solution 
for avoiding GIL threading limitations ? 

