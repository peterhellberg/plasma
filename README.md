# Plasma

Based on the article <http://lodev.org/cgtutor/plasma.html>

## Progress

First plot a sine on the X-axis.

![Plasma progress 001](http://assets.c7.se/skitch/plasma_progress_001-20150118-143315.png)

Now with `X+Y`

![Plasma progress 002](http://assets.c7.se/skitch/plasma_progress_002-20150118-143808.png)

With the a square root thrown in for good measure

![Plasma progress 003](http://assets.c7.se/skitch/plasma_progress_003-20150118-145540.png)

Combine the sines

![Plasma progress 004](http://assets.c7.se/skitch/plasma_progress_004-20150118-150354.png)

Generate a palette

![Plasma palette](http://assets.c7.se/viz/plasma-palette.png)

Animate plasma using the palette

![Plasma progress 005](http://assets.c7.se/viz/plasma-progress-005.gif)

Tweaked the plasma generation

![Plasma progress 006](http://assets.c7.se/skitch/plasma_progress_006-20150118-175223.png)

A few more tweaks

![Plasma progress 007](http://assets.c7.se/skitch/plasma_progress_007-20150118-183534.png)

Change the palette

![Plasma palette 2](http://assets.c7.se/viz/plasma-palette-2.png)

Generate plasma using the new palette

![Plasma progress 008](http://assets.c7.se/skitch/plasma_progress_008-20150118-204536.png)

Support for gradient palettes

![Plasma palette 3](http://assets.c7.se/viz/plasma-palette-3.png)

Generate plasma using the gradient palette

![Plasma progress 009](http://assets.c7.se/skitch/plasma_progress_009-20150118-212046.png)

Another animation

![Plasma progress 010](http://assets.c7.se/viz/plasma-progress-010.gif)

## Commands

### CLI

Rendering palette and plasma as PNG images.

```
Usage of plasma:
  -h=512: Height of the image
  -n=1: Number of frames to generate
  -o="plasma.png": Output file name
  -p="palette.png": Palette file name
  -s=16: Scale of the plasma
  -show=false: Show the generated image
  -w=512: Width of the image
```

### GUI

Window with a rotating plasma image.

![Plasma GUI](http://assets.c7.se/skitch/Plasma_GUI-20150122-201318.png)

Can also be transpiled into JavaScript by [GopherJS](http://www.gopherjs.org/)

[![GopherJS Plasma](http://assets.c7.se/skitch/GopherJS_Plasma_in_Safari-20150122-201529.png)](http://data.gopher.se/js/plasma/)

### Joystick GUI

Control the plasma using a PS4 controller

![Plasma Joystick GUI](http://assets.c7.se/skitch/Plasma_Joystick_GUI-20150120-005750.png)

![Animation](http://assets.c7.se/viz/plasma-joystick-gui.gif)

### Pixels GUI

Using `screen.ReplacePixels` to render the plasma.
