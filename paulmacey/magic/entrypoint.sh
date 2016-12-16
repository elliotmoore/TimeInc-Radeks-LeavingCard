#!/bin/bash
ACTION=${*}

if [[ $ACTION == 'goodbye' ]] ; then
  cowsay -f tux "$(cat /magic/goodbye.txt)" | lolcat
elif [[ $ACTION == 'radek-coding' ]]; then
  cmatrix
elif [[ $ACTION == 'goat-radek' ]]; then
  cacaview /magic/goat.jpeg
elif [[ $ACTION == 'pug-radek' ]]; then
  cacaview /magic/pug.jpg
elif [[ $ACTION == 'breakfast-radek' ]]; then
  cacaview /magic/breakfast.png
elif [[ $ACTION == 'hair-radek' ]]; then
  cacaview /magic/hair.jpg
elif [[ $ACTION == 'radek' ]]; then
  cacaview /magic/radek.png
else
  $ACTION && cowsay -f tux "$(cat /magic/goodbye.txt)" | lolcat
fi
