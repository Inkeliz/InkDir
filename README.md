InkDir
===================


The easiest way to check what folder is eating your disk. Using InkDir you are capable to see the size of all folders and files sorted by the size, using a simple command on cmd or using a context menu.

----------

Features
=======

- It's more detailed than "properties" option, native on Windows.
 > InkDir give the size of all folders individually, you can see what folder is heavy or just empty.

- It's faster than "properties" option.
> The performance can be three times faster, of course with highest CPU usage.

- It's easy to use.
> InkDir can install a simple "Context Menu", allowing to check the folder with a single click.

- It's free.
> InkDir is totally free, anyone can download, review the code or create changes. 

----------

Get started
=======


You can download the lastest release here or you can build the InkDir using `go build`, the InkDir is self-suficient, doesn't need external library.

CLI
---

You need just run `InkDir -path C:\path\to\location\you\want`, this will list something like that:

    Name                                                                   Size
    ----------------------------------------------------------------------------------------------------
    .idea                                                                  17KB       17892
    .git                                                                   9KB        9286
    Display                                                                3KB        4005
    Setup                                                                  3KB        3925
    Directory                                                              1KB        1737
    Walker                                                                 1KB        1664
    Sorter                                                                 621B       621
    Types                                                                  56B        56
    ----------------------------------------------------------------------------------------------------
    Total                                                                  38KB       39186
    Time                                                                   107.2795ms

You want to include the files on the list? Include `-files`. You have more commands, say `InkDir -h` to see all the possibilites.

Context Menu
----------

If you like to use mouse, you can install the Context Menu, this will create a easiest way to check the size of all folders in the directory. To install, just send `InkDir -installMenu`, that all.

![enter image description here](https://raw.githubusercontent.com/Inkeliz/InkDir/v1.0/Media/contextmenu.png)

This is what you can use now.
