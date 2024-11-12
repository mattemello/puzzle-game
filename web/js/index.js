let position = {
    x: 0,
    y: 0
};

const dimensionCanvs = {
    width: 50,
    height: 50

}

const dimensionGrid = {
    rows: 10,
    collo: 10
}

let ctx = [];

function init() {
    let position = 0

    let posl = 0;
    let post = 0;

    for (let i = 0; i < dimensionGrid.rows; ++i) {
        for (let j = 0; j < dimensionGrid.collo; ++j) {
            document.getElementById('container').innerHTML += '<canvas class="canvas" id="canvas' + position + '"style="position: center; margin: 0px; padding: 0px;" onmouseleave="backColor(' + position + ')" onmouseenter="changeColor(' + position + ')">Sorry, bet no canvas</canvas>';
            ++position;
            posl += dimensionCanvs.width;
        }
        posl = 0;
        post += dimensionCanvs.height;
    }



    document.getElementById('container').innerHTML += '<canvas class="hero" id="hero" style="position: absolute; left: 0px; top: 0px; width: 20px; height: 20px;">Sorry no</canvas>';
    const canvas = document.querySelectorAll("canvas");


    for (let i = 0; i < dimensionGrid.rows * dimensionGrid.collo; ++i) {
        ctx[i] = canvas[i].getContext("2d");
        square(ctx[i]);
    }

    /*FIX: use this the hero do not appear */


    const hero = document.getElementById('hero').getContext("2d");

    drawnChar(hero, 0, 0);
}

function drawnChar(hero, x, y) {
    console.log("enter");
    hero.rect(x, y, dimensionCanvs.width * 4, dimensionCanvs.height * 4);
    hero.lineWidth = 1;
    hero.stroke();
    hero.fillStyle = "black";
    hero.fill();
}

function square(ctx) {
    ctx.rect(position.x, position.y, dimensionCanvs.width * 6, dimensionCanvs.height * 6);
    ctx.lineWidth = 1;
    ctx.stroke();
    ctx.fillStyle = "teal";
    ctx.fill();
}

function backColor(position) {
    ctx[position].fillStyle = "teal";
    ctx[position].fill();
}

function changeColor(position) {
    ctx[position].fillStyle = "white";
    ctx[position].fill();
}

function mov() {
    position.x += 10;
    position.y += 1;

    for (i = 0; i < 6; i++) {
        square(ctx[i]);
    }
}

