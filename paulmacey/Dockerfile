FROM        wernight/funbox

MAINTAINER  Paul Macey on behalf of RADEK THE WIZARD

COPY        ./magic /magic

ENTRYPOINT  ["/bin/bash", "-c", "/magic/entrypoint.sh ${*}", "--"]
