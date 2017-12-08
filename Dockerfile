FROM vitr/casperjs
COPY ./ home/casperjs-tests/
WORKDIR home/casperjs-tests/
RUN ./init
RUN npm install
CMD node ./src
EXPOSE 3030

