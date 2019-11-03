# hihex
The `hihex` utility works like the `hexdump` command with a special addition to view "high-bit ascii".  This allows you to view the character output of ASCII values that have the high bit set wheras hexdump doesn't understand these values.

This is particularly useful for examining binary files from older computers, such as the Apple II, where they would commonly use the high bit versions of characters to display text in other modes such as inverse.  

Also great for diffs!


# Still don't get it?

Here's a quick example.  Look at this file with hexdump... 

```
    $ hexdump -C binfile
    00000000  48 45 4C 4C 4F 20 57 4F  52 4C 44 21 21 21 21 21  |HELLO WORLD!!!!!|
    00000010  D3 C5 C3 D2 C5 D4 A0 CD  C5 D3 D3 C1 C7 C5 BA A9  |................|
```

See how the second line is filled with periods?  LAME!!!  It could be ASCII with the high bit set, as was common on the Apple II.  

Now you can see the whole picture with hihex...

```
    $ hihex binfile
    000000  48 45 4C 4C 4F 20 57 4F 52 4C 44 21 21 21 21 21  HELLO WORLD!!!!!
    000010  D3 C5 C3 D2 C5 D4 A0 CD C5 D3 D3 C1 C7 C5 BA A9  SECRET MESSAGE:)
```


