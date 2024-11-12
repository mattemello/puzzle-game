let position = {
    x: 0,
    y: 0
};

const dimension = {
    rows: 10,
    collo: 10
}

let ctx = [];

function init() {
    let tableGame = document.getElementById("tableGame");
    let position = 0

    for (let i = 0; i < dimension.rows; ++i) {
        const tr = tableGame.insertRow();
        for (let j = 0; j < dimension.collo; ++j) {
            //cont.innerHTML += '<td><canvas class="canvas" id="canvas" onmouseleave="backColor(' + position + ')" onmouseenter="changeColor(' + position + ')">Sorry, bet no canvas</canvas></td>';
            const td = tr.insertCell();
            td.innerHTML = '<canvas class="canvas" id="canvas' + position + '" onmouseleave="backColor(' + position + ')" onmouseenter="changeColor(' + position + ')">Sorry, bet no canvas</canvas>';
            ++position;
        }

    }


    const canvas = document.querySelectorAll("canvas");

    for (let i = 0; i < dimension.rows * dimension.collo; i++) {
        ctx[i] = canvas[i].getContext("2d");
        square(ctx[i]);
    }

    /*FIX: use this the hero do not appear */
    /*TODO: do this ↓ but for all the canvas */

    document.getElementById('canvas0').style.position = "absolute";
    document.getElementById('canvas0').style.left = 0;
    document.getElementById('canvas0').style.top = 0;

    hero = document.getElementById('hero');

    drawnChar(hero, 0, 0);
}

function drawnChar(hero, x, y) {
    console.log("enter");
    hero.rect(x, y, 100, 100);
    hero.lineWidth = 1;
    hero.stroke();
    hero.fillStyle = "black";
    hero.fill();
}

function square(ctx) {
    ctx.rect(position.x, position.y, 300, 300);
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
    ctx[position].fillStyle = "black";
    ctx[position].fill();
}

function mov() {
    position.x += 10;
    position.y += 1;

    for (i = 0; i < 6; i++) {
        square(ctx[i]);
    }
}

